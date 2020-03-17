package models

import "time"

const OrderTableName = "orders"

type Order struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	ShowCount int16     `db:"show_count"`
	CreatedAt time.Time `db:"created_at"`
}

func (o Order) TableName() string {
	return OrderTableName
}
