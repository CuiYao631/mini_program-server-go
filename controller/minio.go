/*
 * @Author: CuiYao
 * @Date: 2022-01-28 10:15:24
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:52:42
 */
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Minio interface {
	//路由
	Route(g *echo.Group)

	//*******Bucket操作
	//创建Bucket
	MakeBucket(c echo.Context) error
	//列出Bucket
	ListBucket(c echo.Context) error
	//Bucket是否存在
	ExistsBucket(c echo.Context) error
	//删除Bucket
	Delete(c echo.Context) error
	//列出Bucket中的对象
	ListObject(c echo.Context) error

	//******Object 对象操作
	//文件下载
	FGetObject(c echo.Context) error
	//文件上传
	FPutObject(c echo.Context) error
	//流上传
	PutObject(c echo.Context) error
	//流下载
	GetObject(c echo.Context) error
	//获取URL
	GetObjectUrl(c echo.Context) error
	//删除对象
	DeleteObject(c echo.Context) error
}

func (ctrl *controller) MinioRoute(g *echo.Group) {
	g.POST("/geturl", ctrl.GetObjectUrl)
}

//*******Bucket操作
//创建Bucket
func (ctrl *controller) MakeBucket(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//列出Bucket
func (ctrl *controller) ListBucket(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//Bucket是否存在
func (ctrl *controller) ExistsBucket(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//删除Bucket
func (ctrl *controller) Delete(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//列出Bucket中的对象
func (ctrl *controller) ListObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//******Object 对象操作
//文件下载
func (ctrl *controller) FGetObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//文件上传
func (ctrl *controller) FPutObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//流上传
func (ctrl *controller) PutObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//流下载
func (ctrl *controller) GetObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

//获取URL
func (ctrl *controller) GetObjectUrl(c echo.Context) error {
	bucketName := c.FormValue("bucketName")
	objectName := c.FormValue("objectName")
	if bucketName == "" || objectName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "data is nil")
	}
	url, err := ctrl.uc.GetObjectUrl(c.Request().Context(), bucketName, objectName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, url)
}

//删除对象
func (ctrl *controller) DeleteObject(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
