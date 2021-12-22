/*
 * @Author: CuiYao
 * @Date: 2021-12-22 14:23:18
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 15:14:20
 */
package controller

import (
	"net/http"

	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/labstack/echo/v4"
)

type Resources interface {
	CreateResources(c echo.Context) error
	UpdateResources(c echo.Context) error
	ListResources(c echo.Context) error
	GetResources(c echo.Context) error
	DeleteResources(c echo.Context) error
}

func (ctrl *controller) CreateResources(c echo.Context) error {
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
