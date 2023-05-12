package socket

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/websocket"
	"net"
)

type MsgConfig struct {
	Type string `json:"type,omitempty"`
	Uid  string `json:"uid,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

var connMap = make(map[string]*websocket.Conn)

type socket struct {
	listen net.Listener
	opAi   *openai.Client
}

func MakeSocket(listen net.Listener, opAi *openai.Client) *socket {
	return &socket{listen: listen, opAi: opAi}
}

func (s *socket) Socket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var err error
		for {
			var reply string

			if err = websocket.Message.Receive(ws, &reply); err != nil {
				fmt.Println("Can't receive")
				break
			}
			replyMsg := MsgConfig{}
			json.Unmarshal([]byte(reply), &replyMsg)

			if replyMsg.Type == "login" && replyMsg.Uid != "" {
				connMap[replyMsg.Uid] = ws
				fmt.Println(connMap)
			}
			if replyMsg.Type != "login" {
				for _, v := range connMap {
					go s.Chat(c.Request().Context(), v, replyMsg.Msg)
					//go sendMessage(v, replyMsg.Msg)
				}
			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
