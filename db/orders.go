package db

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/tkachenkom/taxiTestApp/db/models"
)

type OrdersQ interface {
	Insert(order models.Order) error
	Update(order models.Order) error
	Delete(id string) error
	GetByID(id string) (models.Order, error)
	GetAll() ([]models.Order, error)
	DeleteOne() error
	GetOne() (models.Order, error)
}

type OrdersWrapper struct {
	parent *DB
}

func (d *DB) OrdersQ() OrdersQ {
	return &OrdersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (w *OrdersWrapper) Insert(order models.Order) error {
	return w.parent.db.Model(&order).Insert()
}

func (w *OrdersWrapper) Update(order models.Order) error {
	return w.parent.db.Model(&order).Update()
}

func (w *OrdersWrapper) Delete(id string) error {
	_, err := w.parent.db.Delete(models.OrderTableName, dbx.HashExp{"id": id}).Execute()
	return err
}

func (w *OrdersWrapper) GetByID(id string) (models.Order, error) {
	var res models.Order
	err := w.parent.db.Select().From(models.OrderTableName).Where(dbx.HashExp{"id": id}).One(&res)
	return res, err
}

func (w *OrdersWrapper) GetAll() ([]models.Order, error) {
	var res []models.Order
	err := w.parent.db.Select().From(models.OrderTableName).All(&res)
	return res, err
}

func (w *OrdersWrapper) GetOne() (models.Order, error) {
	var res models.Order
	err := w.parent.db.Select().From(models.OrderTableName).One(&res)
	return res, err
}

func (w *OrdersWrapper) DeleteOne() error {
	err := w.parent.db.Delete(models.OrderTableName, nil).One(&struct{}{})
	return err
}
