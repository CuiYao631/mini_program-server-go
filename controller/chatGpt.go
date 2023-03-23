package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChatGpt interface {
	Chat(e echo.Context) error
}

func (ctrl *controller) ChatGptRoute(g *echo.Group) {
	g.POST("/chat", ctrl.Chat)
}

func (ctrl *controller) Chat(e echo.Context) error {
	msg := e.FormValue("msg")
	res, err := ctrl.uc.Chat(e.Request().Context(), msg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}
