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

func (ctrl *controller) UploadWallpaper(c echo.Context) error {
	bucketName := c.FormValue("bucketName")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	err = ctrl.uc.UploadWallpaper(c.Request().Context(), bucketName, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func (ctrl *controller) ListWallpaper(c echo.Context) error {
	wallpaper, err := ctrl.uc.ListWallpaper(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return echo.NewHTTPError(http.StatusOK, wallpaper)
}

func (ctrl *controller) GetWallpaper(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (ctrl *controller) DeleteWallpaper(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
