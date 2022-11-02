package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Wallpaper interface {
	UploadWallpaper(c echo.Context) error
	ListWallpaper(c echo.Context) error
	GetWallpaper(c echo.Context) error
	DeleteWallpaper(c echo.Context) error
}

func (ctrl *controller) WallpaperRoute(g *echo.Group) {
	g.POST("/upload", ctrl.UploadWallpaper)
	g.POST("/list", ctrl.ListWallpaper)
	g.POST("/delete", ctrl.DeleteWallpaper)
}

// UploadWallpaper @Summary 上传壁纸
// @Description 上传壁纸
// @Tags UploadWallpaper
// @Accept  json
// @Produce  json
// @Param bucketName form string true "仓库名"
// @Param file form string true "文件路径"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /wallpaper/upload [post]
func (ctrl *controller) UploadWallpaper(c echo.Context) error {
	bucketName := c.FormValue("bucketName")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	_, _, err = ctrl.uc.UploadWallpaper(c.Request().Context(), bucketName, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

// ListWallpaper @Summary 壁纸列表
// @Description 壁纸列表
// @Tags ListWallpaper
// @success 200 {string} entity.Wallpaper "返回数据"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /wallpaper/list [post]
func (ctrl *controller) ListWallpaper(c echo.Context) error {
	wallpaper, err := ctrl.uc.ListWallpaper(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK, wallpaper)
}

// GetWallpaper @Summary 获取单个壁纸
// @Description 获取单个壁纸
// @Tags GetWallpaper
// @Accept  json
// @Produce  json
// @Param bucketName form string true "仓库名"
// @Param fileName form string true "文件名"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /get [post]
func (ctrl *controller) GetWallpaper(c echo.Context) error {
	bucketName := c.FormValue("bucketName")
	fileName := c.FormValue("fileName")
	url, err := ctrl.uc.GetWallpaper(c.Request().Context(), bucketName, fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, url)
}

// DeleteWallpaper @Summary 删除壁纸
// @Description 删除壁纸
// @Tags DeleteWallpaper
// @Accept  json
// @Produce  json
// @Param bucketName form string true "仓库名"
// @Param fileName form string true "文件名"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /delete [post]
func (ctrl *controller) DeleteWallpaper(c echo.Context) error {
	bucketName := c.FormValue("bucketName")
	fileName := c.FormValue("fileName")
	err := ctrl.uc.DeleteWallpaper(c.Request().Context(), bucketName, fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}
