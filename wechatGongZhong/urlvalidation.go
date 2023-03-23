package wechatGongZhong

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WechatGongZhong interface {
	EchoProcRequest(e echo.Context) error
}

const (
	token = "wechat4go"
)

type wechatGongZhong struct {
}

func MakeWechatGongZhong() *wechatGongZhong {
	return &wechatGongZhong{}
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

// WXMsgReceive 微信消息接收
func WXMsgReceive(e echo.Context) error {
	var textMsg WXTextMsg
	//err := e.ShouldBindXML(&textMsg)
	err := e.Bind(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		//return
	}
	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
	// 对接收的消息进行被动回复
	return WXMsgReply(e, textMsg.ToUserName, textMsg.FromUserName)
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

// WXMsgReply 微信消息回复
func WXMsgReply(e echo.Context, fromUser, toUser string) error {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		//return
	}
	//_, _ = c.Writer.Write(msg)
	return echo.NewHTTPError(http.StatusOK, msg)
}
