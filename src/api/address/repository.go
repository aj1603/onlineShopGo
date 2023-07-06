package address

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/address/schemas"
)

func get_() ([]res.Category, error) {
	var categories []res.Category

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT * FROM categories
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category res.Category

		err := rows.Scan(
			&category.ID,
			&category.NAME,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func create_(address *res.Create) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO addresss ("region", "city", "address_line", "customers_id") VALUES ($1, $2, $3, $4)`,
		address.REGION, address.CITY, address.ADDRESS_LINE, address.CUSTOMERS_ID,
	)
}

func update_(address *res.Update) {
	db.DB.Exec(
		context.Background(),
		`UPDATE addresss SET region = $1, city = $3, address_line = $4, customers_id = $5 WHERE id = $2`,
		address.REGION, address.ID, address.CITY, address.ADDRESS_LINE, address.CUSTOMERS_ID,
	)
}

func remove_(address *res.Delete) {
	fmt.Println(address.ID)
	db.DB.Exec(
		context.Background(),
		`DELETE FROM addresss WHERE id = $1`,
		address.ID,
	)
}

func get_by_category_id_(id int) res.Category {
	var category res.Category

	db.DB.QueryRow(
		context.Background(),
		`SELECT
			c.id,
			c.name,
			array_agg(jsonb_build_object(
				'id', p.id,
				'name', p.name,
				'description', p.description,
				'price', p.price::float,
				'product_sku', p.product_sku,
				'quantity', p.quantity,
				'categories_id', p.categories_id,
				'discounts_id', p.discounts_id,
				'brands_id', p.brands_id,
				'images', pi.images
			)) AS products
		FROM categories c
		LEFT JOIN products p ON p.categories_id = c.id
		LEFT JOIN (
			SELECT
				products_id,
				array_agg(jsonb_build_object(
					'id', id,
					'img_url', img_url,
					'products_id', products_id
				)) AS images
			FROM products_images
			GROUP BY products_id
		) pi ON pi.products_id = p.id
		WHERE c.id = $1
		GROUP BY c.id, c.name`,
		id,
	).Scan(&category.ID, &category.NAME, &category.PRODUCTS)
	return category
}
