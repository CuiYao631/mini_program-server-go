/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:29
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 14:56:56
 */

package controller

import (
	"net/http"

	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/CuiYao631/mini_program-server-go/usecase"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	HoneWallpaper(c echo.Context) error
	// Resources 资源
	Resources
	// Wallpaper 壁纸
	Wallpaper
}
type controller struct {
	uc usecase.Usecase
}

func MakeController(uc usecase.Usecase) *controller {
	return &controller{uc: uc}
}
func (ctrl *controller) Root(c echo.Context) error {

	return echo.NewHTTPError(http.StatusOK, "Welcome to the xcui_Toolbox")
}
func (ctrl *controller) HoneWallpaper(c echo.Context) error {
	url, err := ctrl.uc.GetWallpaper(c.Request().Context(), "homeimage", "homeimage.JPG")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return echo.NewHTTPError(http.StatusOK, url)
}
func (ctrl *controller) Home(c echo.Context) error {
	resources, err := ctrl.uc.ListResources(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusOK)
	}
	var images []entity.RotationImage
	for _, v := range resources {
		image := entity.RotationImage{
			Url:         v.Icon,
			ResourcesID: v.ID,
		}
		images = append(images, image)

	}

	//TODO 这段代码啥作用，忘了
	// for _, v := range resources {
	// 	if v.Topping {
	// 		image := entity.RotationImage{
	// 			Url:         v.Icon,
	// 			ResourcesID: v.ID,
	// 		}
	// 		images = append(images, image)
	// 	}
	// }
	return echo.NewHTTPError(http.StatusOK, images)
}
