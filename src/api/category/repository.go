package category

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/category/schemas"
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

func create_(category *res.Create) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO categories ("name") VALUES ($1)`,
		category.NAME,
	)
}

func update_(category *res.Update) {
	db.DB.Exec(
		context.Background(),
		`UPDATE categories SET name = $1 WHERE id = $2`,
		category.NAME, category.ID,
	)
}

func remove_(id int) {
	db.DB.Exec(
		context.Background(),
		`DELETE FROM categories WHERE id = $1`,
		id,
	)
}

func get_by_category_id_(id int) res.Category {
	var category res.Category

	db.DB.QueryRow(
		context.Background(),
		`SELECT 
			categories.id,
			categories.name,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', products.id,
					'name', products.name,
					'description', products.description,
					'price', products.price::FLOAT,
					'product_sku', products.product_sku,
					'quantity', products.quantity,
					'categories_id', products.categories_id,
					'discounts_id', products.discounts_id,
					'brands_id', products.brands_id
				)
				FROM products
		 		WHERE products.categories_id = categories.id
			) AS products
		FROM categories WHERE id = $1
		`,
		id,
	).Scan(&category.ID, &category.NAME, &category.PRODUCTS)
	fmt.Println(category.PRODUCTS[1].ID)
	return category
}
