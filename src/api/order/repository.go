package order

import (
	"context"
	db "onlineshopgo/database"
	req "onlineshopgo/src/api/order/schemas"
)

func get_() ([]req.Order, error) {
	var orders []req.Order

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT
			o.id,
			o.total,
			o.customers_id,
			o.addresss_id,
			o.order_status_id,
			o.payment_methods_id,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', oi.id,
					'quantity', oi.quantity,
					'products_id', oi.products_id,
					'orders_id', oi.orders_id
				)
				FROM order_items oi
				WHERE oi.orders_id = o.id
			) AS order_items
		FROM orders o
		ORDER BY id`,
	)

	if err != nil {
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order req.Order
		err := rows.Scan(
			&order.ID,
			&order.TOTAL,
			&order.CUSTOMER_ID,
			&order.ADDRESSS_ID,
			&order.ORDER_STATUS_ID,
			&order.PAYMENT_METHODS_ID,
			&order.ORDER_ITEMS,
		)

		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

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

func get_order_user_id_(id float64) ([]req.Order_res, error) {
	var orders []req.Order_res

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT
			o.id,
			o.total,
			o.customers_id,
			o.addresss_id,
			o.order_status_id,
			o.payment_methods_id,
			(
				SELECT
					jsonb_agg(jsonb_build_object(
						'id', oi.id,
						'quantity', oi.quantity,
						'products_id', oi.products_id,
						'orders_id', oi.orders_id,
						'products', (
							SELECT
								jsonb_agg(jsonb_build_object(
									'id', p.id,
									'name', p.name,
									'description', p.description,
									'price', p.price::float,
									'product_sku', p.product_sku,
									'quantity', p.quantity,
									'categories_id', p.categories_id,
									'discounts_id', p.discounts_id,
									'brands_id', p.brands_id,
									'products_images', (
										SELECT
											jsonb_agg(jsonb_build_object(
												'id', pi.id,
												'img_url', pi.img_url,
												'products_id', pi.products_id
											))
										FROM products_images pi
										WHERE pi.products_id = p.id
									)
								))
							FROM products p
							WHERE p.id = oi.products_id
						)
					))
				FROM order_items oi
				WHERE oi.orders_id = o.id
			) AS order_items
		FROM orders o
		WHERE o.customers_id = $1
		GROUP BY o.id`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var order req.Order_res

		err := rows.Scan(
			&order.ID,
			&order.TOTAL,
			&order.CUSTOMER_ID,
			&order.ADDRESSS_ID,
			&order.ORDER_STATUS_ID,
			&order.PAYMENT_METHODS_ID,
			&order.ORDER_ITEMS,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
