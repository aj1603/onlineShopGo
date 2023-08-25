package cart

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	req "onlineshopgo/src/api/cart/schemas"
)

var cart_sku int = 3

func create_(cart *req.Cart) error {
	var cartID int

	_, err := db.DB.Exec(
		context.Background(),
		`INSERT INTO carts ("carts_sku", "customers_id")
		 VALUES ($1, $2)`,
		cart_sku, cart.CUSTOMER_ID,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	db.DB.QueryRow(
		context.Background(),
		`Select id from carts order by id DESC limit 1`,
	).Scan(&cart.ID)

	cartID = cart.ID

	for _, item := range cart.CART_ITEMS {
		_, err = db.DB.Exec(
			context.Background(),
			`INSERT INTO cart_item ("quantity", "products_id", "carts_id")
			 VALUES ($1, $2, $3)`,
			item.QUANTITY, item.PRODUCT_ID, cartID,
		)
		if err != nil {
			return err
		}
	}

	cart_sku += 1
	return nil
}

func get_cart_user_id_(id float64) ([]req.Cart_res, error) {
	var carts []req.Cart_res

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT
			c.id,
			c.carts_sku,
			c.customers_id,
			(
				SELECT
					jsonb_agg(jsonb_build_object(
						'id', ci.id,
						'quantity', ci.quantity,
						'products_id', ci.products_id,
						'carts_id', ci.carts_id,
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
							WHERE p.id = ci.products_id
						)
					))
				FROM cart_item ci
				WHERE ci.carts_id = c.id
			) AS cart_item
		FROM carts c
		WHERE c.customers_id = $1
		GROUP BY c.id`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cart req.Cart_res

		rows.Scan(
			&cart.ID,
			&cart.CART_SKU,
			&cart.CUSTOMER_ID,
			&cart.CART_ITEMS,
		)

		carts = append(carts, cart)
	}

	return carts, nil
}

func remove_(id int) {
	db.DB.Exec(
		context.Background(),
		`DELETE FROM cart_item WHERE carts_id = $1`,
		id,
	)
	db.DB.Exec(
		context.Background(),
		`DELETE FROM carts WHERE id = $1`,
		id,
	)

}
