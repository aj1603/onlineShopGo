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
			products_images.id,
			products_images.img_url,
			products_images.products_id
		FROM products LEFT JOIN products_images ON products.id = products_images.products_id
		LIMIT $1
		OFFSET $2
		`, limit, offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product res.Product
		var products_images res.Product_images

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
			&products_images.ID,
			&products_images.IMG_URL,
			&products_images.PRODUCT_ID,
		)

		if err != nil {
			return nil, err
		}
		product.PRODUCT_IMG = products_images
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
		`SELECT id, name, description, price, product_sku, quantity, categories_id, discounts_id,
		brands_id
		FROM products WHERE id = $1`,
		id,
	).Scan(&product.ID, &product.NAME, &product.DESCRIPTION, &product.PRICE, &product.PRODUCT_SKU,
		&product.QUANTITY, &product.CATEGORY_ID, &product.DISCOUNT_ID, &product.BRAND_ID)
	return product
}

func get_by_category_id_(categories_id int) ([]res.Product, error) {
	var products []res.Product

	rows, err := db.DB.Query(
		context.Background(),
		`SELECT id, name, description, price, product_sku, quantity, categories_id, 
		discounts_id, brands_id
		FROM products WHERE categories_id = $1`,
		categories_id,
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
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func search_from_word_(search_word string) ([]res.Product, error) {
	var products []res.Product
	fmt.Println(search_word)

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
			products_images.id,
			products_images.img_url,
			products_images.products_id
		FROM products LEFT JOIN products_images ON products.id = products_images.products_id
		WHERE name LIKE ('`+search_word+`%') or description LIKE ('`+search_word+`%')`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product res.Product
		var products_images res.Product_images

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
			&products_images.ID,
			&products_images.IMG_URL,
			&products_images.PRODUCT_ID,
		)

		if err != nil {
			return nil, err
		}
		product.PRODUCT_IMG = products_images
		products = append(products, product)
	}

	return products, nil
}
