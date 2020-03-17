package order

import (
	"github.com/labstack/echo/v4"
	"github.com/tkachenkom/taxiTestApp/server/handlers"
	"net/http"
)

func (h *Handler) GetOrder(c echo.Context) error {
	order, err := h.orderDB.GetOne()
	if err != nil {
		h.log.WithError(err).Error("failed to get order from db")
		return handlers.InternalError(c)
	}

	order.ShowCount++

	err = h.orderDB.Update(order)
	if err != nil {
		h.log.WithError(err).Error("failed to update order from db")
		return handlers.InternalError(c)
	}

	return c.String(http.StatusOK, order.Name)
}
