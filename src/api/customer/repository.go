package customer

import (
	"context"
	"fmt"
	db "onlineshopgo/database"
	res "onlineshopgo/src/api/customer/schemas"
)

func get_() ([]res.Customer, error) {
	var customers []res.Customer

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			password,
			email,
			phone_num
		FROM customers
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var customer res.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.NAME,
			&customer.PASSWORD,
			&customer.EMAIL,
			&customer.PHONE_NUM,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func create_(customer *res.Create) {
	db.DB.Exec(
		context.Background(),
		`INSERT INTO customers ("name", "password", "email", "phone_num") VALUES ($1,$2,$3,$4)`,
		customer.NAME, customer.PASSWORD, customer.EMAIL, customer.PHONE_NUM,
	)
}

func update_(customer *res.Update) {
	_, err := db.DB.Exec(
		context.Background(),
		`UPDATE products SET name = $1, password = $3, email = $4, phone_num = $5 WHERE id = $2`,
		customer.NAME, customer.ID, customer.PASSWORD, customer.EMAIL, customer.PHONE_NUM,
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}
