package brands

import (
	"context"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/brands/schemas"
)

func get_() ([]res.Brand, error) {
	var brands []res.Brand

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT * FROM brands
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand res.Brand

		err := rows.Scan(
			&brand.ID,
			&brand.NAME,
			&brand.IMG_URL,
		)

		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}

func create_(brand *res.Create) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO brands ("name", "img_url") VALUES ($1,$2)`,
		brand.NAME, brand.IMG_URL,
	)
}

func update_(brand *res.Update) {
	db.DB.Exec(
		context.Background(),
		`UPDATE brands SET name = $1, img_url = $2 WHERE id = $3`,
		brand.NAME, brand.IMG_URL, brand.ID,
	)
}

func remove_(id int) {
	db.DB.Exec(
		context.Background(),
		`DELETE FROM brands WHERE id = $1`,
		id,
	)
}
