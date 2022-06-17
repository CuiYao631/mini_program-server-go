package wechatGongZhong

import (
	"crypto/sha1"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
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
