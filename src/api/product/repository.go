package product

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/product/schemas"
)

func get_(limit, offset int) ([]res.Product, error) {
	var products []res.Product

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT
			products.id,
			products.name,
			products.description,
			products.price::FLOAT,
			products.product_sku,
			products.quantity,
			products.categories_id,
			products.brands_id,
			products.discounts_id,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', products_images.id,
					'img_url', products_images.img_url,
					'products_id', products_images.products_id
				)
				FROM products_images
				WHERE products_images.products_id = products.id
			) AS products_images
		FROM products
		LIMIT $1
		OFFSET $2
		`, limit, offset,
	)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product res.Product

		rows.Scan(
			&product.ID,
			&product.NAME,
			&product.DESCRIPTION,
			&product.PRICE,
			&product.PRODUCT_SKU,
			&product.QUANTITY,
			&product.CATEGORY_ID,
			&product.DISCOUNT_ID,
			&product.BRAND_ID,
			&product.PRODUCT_IMG,
		)

		products = append(products, product)
	}
	return products, nil
}

func create_(product *res.Create) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO products ("name", "description", "price", "product_sku", "quantity", "categories_id", "discounts_id", "brands_id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		product.NAME, product.DESCRIPTION, product.PRICE, product.PRODUCT_SKU, product.QUANTITY, product.CATEGORY_ID, product.DISCOUNT_ID, product.BRAND_ID,
	)
}

func update_(product *res.Update) {
	_, err := db.DB.Exec(
		context.Background(),
		`UPDATE products SET name = $1, description = $3, price = $4, product_sku = $5, quantity = $6, categories_id = $7, brands_id = $8, discounts_id = $9 WHERE id = $2`,
		product.NAME, product.ID, product.DESCRIPTION, product.PRICE, product.PRODUCT_SKU, product.QUANTITY, product.CATEGORY_ID, product.BRAND_ID, product.DISCOUNT_ID,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func remove_(id int) {
	db.DB.Exec(
		context.Background(),
		`DELETE FROM products WHERE id = $1`,
		id,
	)
}

func get_by_id_(id int) res.Product {
	var product res.Product

	db.DB.QueryRow(
		context.Background(),
		`SELECT 
			products.id,
			products.name,
			products.description,
			products.price::FLOAT,
			products.product_sku,
			products.quantity,
			products.categories_id,
			products.discounts_id,
			products.brands_id,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', products_images.id,
		 			'img_url', products_images.img_url,
		 			'products_id', products_images.products_id
				)
				FROM products_images
		 		WHERE products_images.products_id = products.id
			) AS products_images
		FROM products WHERE id = $1
		`,
		id,
	).Scan(&product.ID, &product.NAME, &product.DESCRIPTION, &product.PRICE, &product.PRODUCT_SKU,
		&product.QUANTITY, &product.CATEGORY_ID, &product.DISCOUNT_ID, &product.BRAND_ID, &product.PRODUCT_IMG)
	return product
}

func search_from_word_(search_word string) ([]res.Product, error) {
	var products []res.Product

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT 
			products.id,
			products.name,
			products.description,
			products.price,
			products.product_sku,
			products.quantity,
			products.categories_id,
			products.discounts_id,
			products.brands_id,
			ARRAY(
				SELECT JSON_BUILD_OBJECT(
					'id', products_images.id,
		 			'img_url', products_images.img_url,
		 			'products_id', products_images.products_id
				)
				FROM products_images
		 		WHERE products_images.products_id = products.id
			) AS products_images
		FROM products WHERE name ~* ('^' || $1) or description ~* ('^' || $1)`,
		search_word,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product res.Product

		err := rows.Scan(
			&product.ID,
			&product.NAME,
			&product.DESCRIPTION,
			&product.PRICE,
			&product.PRODUCT_SKU,
			&product.QUANTITY,
			&product.CATEGORY_ID,
			&product.DISCOUNT_ID,
			&product.BRAND_ID,
			&product.PRODUCT_IMG,
		)

		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func favorite_(product *res.Favorite) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO favorities ("customers_id", "products_id") VALUES ($1,$2)`,
		product.CUSTOMER_ID, product.PRODUCT_ID,
	)
}

func get_user_favorite_(id float64) ([]res.Favorite, error) {
	var favoriteis []res.Favorite

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT 
			f.id,
			f.customers_id,
			f.products_id,
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
				'products_images', pi.products_images
			)) AS products
		FROM favorities f
		LEFT JOIN products p ON p.id = f.products_id
		LEFT JOIN (
			SELECT
				products_id,
				array_agg(jsonb_build_object(
					'id', id,
					'img_url', img_url,
					'products_id', products_id
				)) AS products_images
			FROM products_images
			GROUP BY products_id
		) pi ON pi.products_id = p.id
		WHERE f.customers_id = $1
		GROUP BY f.id`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var favorite res.Favorite

		err := rows.Scan(
			&favorite.ID,
			&favorite.CUSTOMER_ID,
			&favorite.PRODUCT_ID,
			&favorite.PRODUCTS,
		)

		if err != nil {
			return nil, err
		}
		favoriteis = append(favoriteis, favorite)
	}

	return favoriteis, nil
}
