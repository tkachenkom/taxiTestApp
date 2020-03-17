package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// BadRequest - handler
func BadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, struct {
		Code  int32   `json:"code"`
		Error string `json:"error"`
	}{
		Code:  http.StatusBadRequest,
		Error: err.Error(),
	})
}

// InternalError - handler
func InternalError(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "")
}

// Unauthorized - handler
func Unauthorized(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, struct {
		Code  int32   `json:"code"`
		Error string `json:"error"`
	}{
		Code:  http.StatusUnauthorized,
		Error: err.Error(),
	})
}
