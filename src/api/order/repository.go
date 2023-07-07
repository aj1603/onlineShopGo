package order

import (
	"context"
	db "onlineshopgo/database"
	req "onlineshopgo/src/api/order/schemas"
)

func create_(order *req.Order) error {
	var orderID int

	_, err := db.DB.Exec(
		context.Background(),
		`INSERT INTO orders ("total", "customers_id", "addresss_id", "order_status_id", "payment_methods_id")
		 VALUES ($1, $2, $3, $4, $5)`,
		order.TOTAL, order.CUSTOMER_ID, order.ADDRESSS_ID, 2, order.PAYMENT_METHODS_ID,
	)

	if err != nil {
		return err
	}

	db.DB.QueryRow(
		context.Background(),
		`Select id from orders order by id DESC limit 1`,
	).Scan(&order.ID)

	orderID = order.ID

	for _, item := range order.ORDER_ITEMS {
		_, err = db.DB.Exec(
			context.Background(),
			`INSERT INTO order_items ("quantity", "products_id", "orders_id")
			 VALUES ($1, $2, $3)`,
			item.QUANTITY, item.PRODUCTS_ID, orderID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
