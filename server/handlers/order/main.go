package order

import (
	"github.com/sirupsen/logrus"
	"github.com/tkachenkom/taxiTestApp/db"
)

// Handler - order handler
type Handler struct {
	orderDB db.OrdersQ
	log     *logrus.Entry
}

// New - created new admin handler
func New(log *logrus.Entry, orderDB db.OrdersQ) *Handler {
	return &Handler{
		log:     log,
		orderDB: orderDB,
	}
}
