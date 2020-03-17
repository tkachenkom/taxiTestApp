package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tkachenkom/taxiTestApp/db"
	"github.com/tkachenkom/taxiTestApp/server/handlers/order"
	"net/http"
)

// Router for the api public/private endpoints
func Router(
	log *logrus.Entry,
	db db.QInterface,
	e *echo.Echo,
) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello taxi API!")
	})

	e.GET("/order", order.New(log, db.OrdersQ()).GetOrder)
	e.GET("/orders", order.New(log, db.OrdersQ()).GetAll)

	return e
}
