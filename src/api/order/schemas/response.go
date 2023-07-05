package schemas

type Order struct {
	ID                 int     `json:"id"`
	TOTAL              float32 `json:"total"`
	CUSTOMER_ID        int     `json:"customer_id"`
	ADDRESSS_ID        int     `json:"addresss_id"`
	ORDER_STATUS_ID    int     `json:"order_status_id"`
	PAYMENT_METHODS_ID int     `json:"payment_methods_id"`
}

type Order_items struct {
	ID          int `json:"id"`
	QUANTITY    int `json:"quantity"`
	PRODUCTS_ID int `json:"products_id"`
	ORDERS_ID   int `json:"orders_id"`
}

type Order_status struct {
	ID     int    `json:"id"`
	STATUS string `json:"status"`
}

type Payment_methods struct {
	ID      int    `json:"id"`
	METHODS string `json:"methods"`
}

type Payment_methods_trans struct {
	ID                 int    `json:"id"`
	METHODS            string `json:"method"`
	LANGUAGES_ID       int    `json:"languages_id"`
	PAYMENT_METHODS_ID int    `json:"payment_methods_id"`
}

type Order_status_trans struct {
	ID              int    `json:"id"`
	STATUS          string `json:"status"`
	LANGUAGES_ID    int    `json:"languages_id"`
	ORDER_STATUS_ID int    `json:"order_status_id"`
}
