package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChatMsg struct {
	Msg string `json:"msg"`
}

type ChatGpt interface {
	Chat(e echo.Context) error
}

func (ctrl *controller) ChatGptRoute(g *echo.Group) {
	g.POST("/chat", ctrl.Chat)
}

func (ctrl *controller) Chat(e echo.Context) error {
	input := ""
	mmg := ChatMsg{}
	msg := e.FormValue("msg")
	if err := e.Bind(&mmg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if msg != "" {
		input = msg
	}
	if mmg.Msg != "" {
		input = mmg.Msg
	}

	res, err := ctrl.uc.Chat(e.Request().Context(), input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}
