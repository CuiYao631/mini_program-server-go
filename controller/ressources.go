/*
 * @Author: CuiYao
 * @Date: 2021-12-22 14:23:18
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:16:55
 */
package controller

import (
	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Resources interface {
	//上传
	UploadResourcesIcon(c echo.Context) error
	//创建新的资源
	CreateResources(c echo.Context) error
	//更新资源
	UpdateResources(c echo.Context) error
	//资源列表
	ListResources(c echo.Context) error
	//获取资源
	GetResources(c echo.Context) error
	//删除资源
	DeleteResources(c echo.Context) error
}

func (ctrl *controller) ResourcesRoute(g *echo.Group) {
	g.POST("/create", ctrl.CreateResources)
	g.POST("/update", ctrl.UpdateResources)
	g.POST("/list", ctrl.ListResources)
	g.POST("/get", ctrl.GetResources)
	g.POST("/del", ctrl.DeleteResources)
}

// CreateResources @Summary 创建资源
// @Description 创建资源 先上传壁纸再保存资源
// @Tags 创建资源
// @Accept  json
// @Produce  json
// @Param file form string true "文件"
// @Param name form string true "资源名称"
// @Param tag form string true "标签"
// @Param desc form string true "描述"
// @Param explain form string true "链接"
// @Param topping form string true "是否置顶"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /create [post]
func (ctrl *controller) CreateResources(c echo.Context) error {

	bucketName := os.Getenv("RESOURCES_PHOTO_PATH")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	url, err := ctrl.uc.UploadResourcesIcon(c.Request().Context(), bucketName, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	isTop := false
	name := c.FormValue("name")
	tag := c.FormValue("tag")
	desc := c.FormValue("desc")
	explain := c.FormValue("explain")
	Url := c.FormValue("url")
	topping := c.FormValue("topping")
	if topping == "true" {
		isTop = true
	}
	recs := &entity.Resources{
		Icon:    url,
		Name:    name,
		Tag:     tag,
		Desc:    desc,
		Explain: explain,
		Url:     Url,
		Topping: isTop,
	}
	err = ctrl.uc.CreateResources(c.Request().Context(), *recs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

// UpdateResources @Summary 创建资源
// @Description 创建资源 先上传壁纸再保存资源
// @Tags UpdateResources
// @Accept  json
// @Produce  json
// @Param file form string true "文件"
// @Param name form string true "资源名称"
// @Param tag form string true "标签"
// @Param desc form string true "描述"
// @Param explain form string true "链接"
// @Param topping form string true "是否置顶"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /create [post]
func (ctrl *controller) UpdateResources(c echo.Context) error {
	recs := &entity.Resources{}
	if err := c.Bind(recs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := ctrl.uc.CreateResources(c.Request().Context(), *recs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func (ctrl *controller) ListResources(c echo.Context) error {
	res, err := ctrl.uc.ListResources(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func (ctrl *controller) GetResources(c echo.Context) error {
	id := c.FormValue("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is nil")
	}
	rec, err := ctrl.uc.GetResources(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, rec)
}

// DeleteResources @Summary 删除资源
// @Description 删除资源
// @Tags DeleteResources
// @Accept  json
// @Produce  json
// @Param id form string true "资源id"
// @success 200 {string} string "OK"
// @Failure 400 {string} string "解析请求body错误"
// @Failure 404 {string} string "请求路径错误"
// @Failure 500 {string} string "服务器内部错误"
// @Router /create [post]
func (ctrl *controller) DeleteResources(c echo.Context) error {
	id := c.FormValue("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is nil")
	}
	err := ctrl.uc.DeleteResources(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}
