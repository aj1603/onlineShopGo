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
	fmt.Println(cart)

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

	fmt.Println("HI")

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
	fmt.Println(cart_sku)
	return nil
}

func get_cart_user_id_(id int) ([]req.Cart, error) {
	var carts []req.Cart

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT
			c.id,
			c.carts_sku,
			c.customers_id,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', ci.id,
					'quantity', ci.quantity,
					'products_id', ci.products_id,
					'carts_id', ci.carts_id
				)
				FROM cart_item ci
				WHERE ci.carts_id = c.id
			) AS cart_item
		FROM carts c
		WHERE c.customers_id = $1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cart req.Cart

		err := rows.Scan(
			&cart.ID,
			&cart.CART_SKU,
			&cart.CUSTOMER_ID,
			&cart.CART_ITEMS,
		)

		if err != nil {
			return nil, err
		}

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
