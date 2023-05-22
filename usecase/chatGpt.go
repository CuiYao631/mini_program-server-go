package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/websocket"
	"io"
)

type ChatGpt interface {
	Chat(ctx context.Context, conn *websocket.Conn, input string)
}

func (uc *usecase) Chat(ctx context.Context, conn *websocket.Conn, input string) {
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
	stream, err := uc.opAi.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")

		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
		}

		if len(response.Choices) > 0 {
			//一段语句结束后会返回 stop 判断如果返回了stop 就不再往socket中放数据
			if response.Choices[0].FinishReason == "stop" {
				break
			}
			//打印数据
			fmt.Printf(response.Choices[0].Delta.Content)

			//TODO  后面需要保存聊天记录

			sendMessage(conn, response.Choices[0].Delta.Content)
		} else {
			//sendMessage(conn, "哎呀,我开小差了,请重新输入")
			break
		}

	}
}

func sendMessage(conn *websocket.Conn, msg string) {
	if err := websocket.Message.Send(conn, msg); err != nil {
		fmt.Println("Can't send")
	}
}
