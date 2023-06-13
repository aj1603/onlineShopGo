package category

import (
	"context"
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
