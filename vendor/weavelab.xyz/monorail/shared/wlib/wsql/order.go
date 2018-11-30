package wsql

import (
	"fmt"
	"strings"

	"weavelab.xyz/monorail/shared/go-utilities/dev"
)

type OrderBy struct {
	orders []Order
}

type Order struct {
	field     string
	direction string
}

func NewOrderByFromString(values []string, defaultOrder []string, allowed []string) (OrderBy, error) {

	var orders OrderBy

	for _, v := range values {
		if v == "" {
			continue
		}

		direction := "ASC"
		if strings.HasPrefix(v, "-") {
			v = strings.TrimLeft(v, "-")
			direction = "DESC"
		}

		added := false
		for _, a := range allowed {

			if a == v {
				orders.orders = append(orders.orders, Order{field: v, direction: direction})
				added = true
				break
			}
		}

		if !added {
			err := fmt.Errorf("order by not added %s %v requested: %v", v, allowed, values)
			if dev.IsDev() {
				return orders, err
			}
		}
	}

	if len(orders.orders) == 0 && len(defaultOrder) > 0 {
		return NewOrderByFromString(defaultOrder, []string{}, allowed)
	}

	return orders, nil
}

func (o *OrderBy) GetDirection(field string) string {
	for _, v := range o.orders {
		if v.field == field {
			return v.direction
		}
	}
	return ""
}

func (o *OrderBy) ToSQL() string {
	if len(o.orders) == 0 {
		return ""
	}

	var s string
	for i, v := range o.orders {
		var sep string
		if i > 0 {
			sep = ","
		}
		s += sep + " " + v.field + " " + v.direction
	}

	return s
}
