package schemas

import (
	"encoding/json"
	"onlineshopgo/src/tools"

	"github.com/gin-gonic/gin"
)

type Cart struct {
	ID          int          `json:"id"`
	CART_SKU    string       `json:"carts_sku"`
	CUSTOMER_ID int          `json:"customers_id"`
	CART_ITEMS  []Cart_items `json:"cart_items"`
}

type Cart_items struct {
	ID         int `json:"id"`
	QUANTITY   int `json:"quantity"`
	PRODUCT_ID int `json:"products_id"`
	CARD_ID    int `json:"carts_id"`
}

func Validate_cart(ctx *gin.Context) {
	var schema Cart
	data, _ := ctx.GetRawData()

	json.Unmarshal(data, &schema)
	errors := tools.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Set("data", schema)
	ctx.Next()
}
