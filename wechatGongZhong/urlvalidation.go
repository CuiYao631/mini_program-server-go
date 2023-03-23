package wechatGongZhong

import (
	"context"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WechatGongZhong interface {
	EchoProcRequest(e echo.Context) error
	WXMsgReceive(e echo.Context) error
}

const (
	token = "wechat4go"
)

type wechatGongZhong struct {
	opAi *openai.Client
}

func MakeWechatGongZhong(opAi *openai.Client) *wechatGongZhong {
	return &wechatGongZhong{
		opAi: opAi,
	}
}
func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateurl(e echo.Context) bool {
	w := e.Response()
	timestamp := e.FormValue("timestamp")
	nonce := e.FormValue("nonce")
	signatureGen := makeSignature(timestamp, nonce)

	signatureIn := e.FormValue("signature")
	if signatureGen != signatureIn {
		return false
	}
	echostr := e.FormValue("echostr")

	fmt.Fprintf(w, echostr)
	return true
}

func (ctrl *wechatGongZhong) EchoProcRequest(e echo.Context) error {
	e.Request().ParseForm()
	if !validateurl(e) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
	}
	log.Println("Wechat Service: validateUrl Ok!")
	return echo.NewHTTPError(http.StatusOK)
}

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReceive 微信消息接收
func (ctrl *wechatGongZhong) WXMsgReceive(e echo.Context) error {

	// 获取 POST 请求中的消息数据
	data := make([]byte, e.Request().ContentLength)
	e.Request().Body.Read(data)

	// 将消息数据解析为一个 XML 结构体
	var msg WXTextMsg
	xml.Unmarshal(data, &msg)

	// 根据消息类型生成回复内容
	var replyMsg WXRepTextMsg
	switch msg.MsgType {
	case "text":

		res, err := Chat(ctrl.opAi, msg.Content)
		if err != nil {
			// 默认回复
			replyMsg = WXRepTextMsg{
				ToUserName:   msg.FromUserName,
				FromUserName: msg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      "text",
				Content:      "暂不支持此类消息类型的回复。",
			}
			// 将回复消息转换为 XML 格式
			respData, _ := xml.Marshal(replyMsg)
			return e.Blob(http.StatusOK, "application/xml", respData)
		}
		// 生成文本消息回复
		replyMsg = WXRepTextMsg{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      res,
		}
	default:
		// 默认回复
		replyMsg = WXRepTextMsg{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      "暂不支持此类消息类型的回复。",
		}
	}

	// 将回复消息转换为 XML 格式
	respData, _ := xml.Marshal(replyMsg)

	// 返回回复内容
	return e.Blob(http.StatusOK, "application/xml", respData)
}

func Chat(opAi *openai.Client, input string) (string, error) {
	//client := openai.NewClient("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")
	messages := make([]openai.ChatCompletionMessage, 0)
	//reader := bufio.NewReader(os.Stdin)

	//for {
	//fmt.Print("-> ")
	//text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	})

	resp, err := opAi.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		//continue
	}

	content := resp.Choices[0].Message.Content
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	fmt.Println(content)
	return content, nil
	//}
}
