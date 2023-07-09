package schemas

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
