/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:29
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 14:56:56
 */

package controller

import (
	"net/http"

	"github.com/CuiYao631/mini_program-server-go/usecase"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Resources
}
type controller struct {
	uc usecase.Usecase
}

func MakeController(uc usecase.Usecase) *controller {
	return &controller{uc: uc}
}
func (ctrl *controller) Root(c echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
