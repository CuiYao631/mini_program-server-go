/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:29
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 14:56:56
 */

package controller

import (
	"html/template"
	"io"
	"net/http"

	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/CuiYao631/mini_program-server-go/usecase"
	"github.com/labstack/echo/v4"
)

type HoneWallpaper struct {
	URL string
}

// Template 实现Renderer 接口
type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

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

	return c.Render(200, "index.html", "World")
}
func (ctrl *controller) HoneWallpaper(c echo.Context) error {
	url, err := ctrl.uc.GetWallpaper(c.Request().Context(), "homeimage", "homeimage.JPG")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	honeWallpaper := HoneWallpaper{url}
	return echo.NewHTTPError(http.StatusOK, honeWallpaper)
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
			ResourcesID: v.Name,
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
