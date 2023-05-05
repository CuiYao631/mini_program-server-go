package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

type ChatGpt interface {
	Chat(ctx context.Context, input string) (string, error)
}

func (u *usecase) Chat(ctx context.Context, input string) (string, error) {
	//client := openai.NewClient("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")
	//messages := make([]openai.ChatCompletionMessage, 0)
	////reader := bufio.NewReader(os.Stdin)
	//
	//for {
	//	fmt.Print("-> ")
	//	//text, _ := reader.ReadString('\n')
	//	// convert CRLF to LF
	//	input = strings.Replace(input, "\n", "", -1)
	//	messages = append(messages, openai.ChatCompletionMessage{
	//		Role:    openai.ChatMessageRoleUser,
	//		Content: input,
	//	})
	//
	//	resp, err := u.opAi.CreateChatCompletion(
	//		context.Background(),
	//		openai.ChatCompletionRequest{
	//			Model:    openai.GPT3Dot5Turbo,
	//			Messages: messages,
	//		},
	//	)
	//
	//	if err != nil {
	//		fmt.Printf("ChatCompletion error: %v\n", err)
	//		continue
	//	}
	//
	//	content := resp.Choices[0].Message.Content
	//	messages = append(messages, openai.ChatCompletionMessage{
	//		Role:    openai.ChatMessageRoleAssistant,
	//		Content: content,
	//	})
	//	fmt.Println(content)
	//}
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
	stream, err := u.opAi.CreateChatCompletionStream(ctx, req)
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
	}
	return "content", nil
}

//resp, err := u.opAi.CreateChatCompletion(
//	context.Background(),
//	openai.ChatCompletionRequest{
//		Model: openai.GPT3Dot5Turbo,
//		Messages: []openai.ChatCompletionMessage{
//			{
//				Role:    openai.ChatMessageRoleUser,
//				Content: input,
//			},
//		},
//	},
//)
//
//if err != nil {
//	fmt.Printf("ChatCompletion error: %v\n", err)
//	return "ChatCompletion error: " + err.Error(), nil
//}

//fmt.Println(resp.Choices[0].Message.Content)
//return resp.Choices[0].Message.Content, nil
//}
