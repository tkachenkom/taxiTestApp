package order

import (
	"github.com/labstack/echo/v4"
	"github.com/tkachenkom/taxiTestApp/server/handlers"
	"net/http"
)

func (h *Handler) GetAll(c echo.Context) error {
	orders, err := h.orderDB.GetAll()
	if err != nil {
		h.log.WithError(err).Error("failed to get order from db")
		return handlers.InternalError(c)
	}

	return c.JSON(http.StatusOK, orders)
}
