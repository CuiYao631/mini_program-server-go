package socket

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/websocket"
	"io"
)

func (s *socket) Chat(ctx context.Context, conn *websocket.Conn, input string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: input,
			},
		},
		Stream: true,
	}
	stream, err := s.opAi.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return "", err
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return "", err
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return "", err
		}

		fmt.Printf(response.Choices[0].Delta.Content)
		sendMessage(conn, response.Choices[0].Delta.Content)
	}
	return "content", nil
}

func sendMessage(conn *websocket.Conn, msg string) {
	//msg := replyMsg.Uid + "说:" + replyMsg.Msg
	//if connUid == replyMsg.Uid {
	//	fmt.Println(msg)
	//	if replyMsg.Type == "login" {
	//		msg = "你好！我是你的AI助理，有什么可以帮助你的吗？"
	//	} else {
	//		msg = "你说：" + replyMsg.Msg
	//	}
	//
	//}

	if err := websocket.Message.Send(conn, msg); err != nil {
		fmt.Println("Can't send")
	}

}
