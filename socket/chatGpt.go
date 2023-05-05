package socket

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net"
)

func (s *socket) Chat(ctx context.Context, conn net.Conn, input string) (string, error) {
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
		SendMessage(conn, response.Choices[0].Delta.Content)
	}
	return "content", nil
}
func SendMessage(conn net.Conn, message string) {
	data, err := Encode(message)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = conn.Write(data) // 发送数据
	if err != nil {
		fmt.Println(err.Error())
	}

}
