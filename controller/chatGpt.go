package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type MsgConfig struct {
	Type string `json:"type,omitempty"`
	Uid  string `json:"uid,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

var connMap = make(map[string]*websocket.Conn)

func (ctrl *controller) GPTChat(c echo.Context) error {
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
				fmt.Println("connMap-->>>>", connMap)
				//go sendMessage(connMap[replyMsg.Uid], replyMsg.Msg)
			}
			//CPT-chat
			if replyMsg.Type == "msg" {
				go ctrl.uc.Chat(c.Request().Context(), connMap[replyMsg.Uid], replyMsg.Msg)
			}
			//DALL-E 2 image generation
			if replyMsg.Type == "img" {
				go ctrl.uc.GenerateImages(c.Request().Context(), connMap[replyMsg.Uid], replyMsg.Msg)
			}
			//Audio Captions
			if replyMsg.Type == "audio" {

			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
