package schemas

import (
	"encoding/json"
	"fmt"
	"onlineshopgo/src/tools"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID                 int           `json:"id"`
	TOTAL              float32       `json:"total"`
	CUSTOMER_ID        int           `json:"customer_id"`
	ADDRESSS_ID        int           `json:"addresss_id"`
	ORDER_STATUS_ID    int           `json:"order_status_id"`
	PAYMENT_METHODS_ID int           `json:"payment_methods_id"`
	ORDER_ITEMS        []Order_items `json:"order_items"`
}

type Order_items struct {
	ID          int `json:"id"`
	QUANTITY    int `json:"quantity"`
	PRODUCTS_ID int `json:"products_id"`
	ORDERS_ID   int `json:"orders_id"`
}

func Validate_order(ctx *gin.Context) {
	var schema Order
	data, _ := ctx.GetRawData()

	json.Unmarshal(data, &schema)
	errors := tools.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	fmt.Println("data", schema)
	ctx.Set("data", schema)
	ctx.Next()
}
