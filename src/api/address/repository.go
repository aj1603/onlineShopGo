package address

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/address/schemas"
)

func get_(id float64) ([]res.Address, error) {
	var addresss []res.Address

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT * FROM addresss
		WHERE customers_id = $1
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var address res.Address

		err := rows.Scan(
			&address.ID,
			&address.REGION,
			&address.CITY,
			&address.ADDRESS_LINE,
			&address.CUSTOMERS_ID,
		)

		if err != nil {
			return nil, err
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
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
