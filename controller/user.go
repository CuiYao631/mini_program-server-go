/*
 * @Author: CuiYao
 * @Date: 2021-12-10 17:06:58
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 18:01:25
 */

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	CreateUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	ListUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

func (ctrl *controller) UserRoute(g *echo.Group) {
	g.POST("/login", ctrl.Login)
	g.POST("/add", ctrl.CreateUser)
	g.POST("/update", ctrl.UpdateUser)
	g.POST("/list", ctrl.ListUser)
	g.POST("/delete", ctrl.DeleteUser)
}
func (ctrl *controller) Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	if name == "admin" && password == "admin" {
		return echo.NewHTTPError(http.StatusOK)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "not user")
	}
	return echo.NewHTTPError(http.StatusOK, "username:"+name+"、password:"+password+"")
}
func (ctrl *controller) CreateUser(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusOK, "CreateUser")
}

func (ctrl *controller) UpdateUser(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusOK, "UpdateUser")
}

func (ctrl *controller) ListUser(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusOK, "ListUser")
}

func (ctrl *controller) DeleteUser(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusOK, "DeleteUser")
}
