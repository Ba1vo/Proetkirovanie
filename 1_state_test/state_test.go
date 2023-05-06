package state

import (
	"testing"
)

func TestState(t *testing.T) {
	Books := []CardBook{
		{
			ID:       1,
			Name:     "Конец вечности",
			Price:    "555.55",
			Discount: 5,
			Photo:    "default",
			Authors:  []string{"Айзек Азимов"},
			Amount:   5,
		},
		{
			ID:       2,
			Name:     "Город и звёзды",
			Price:    "777.55",
			Discount: 0,
			Photo:    "default",
			Authors:  []string{"Артур Кларк"},
			Amount:   100,
		},
	}
	order := NewOrder(1, "Pushkina", Books)
	order.OrderInfo()
	order.ChangeInfo(1, "Lenina", Books)
	order.OrderInfo()
	order.Forward()
	order.OrderInfo()
	order.ChangeInfo(1, "Moscovskaya", Books)
	order.Forward()
	order.OrderInfo()
	order.Forward()
	order.Cancel()
}
