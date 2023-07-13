package schemas

import (
	"encoding/json"
	"onlineshopgo/src/tools"

	"github.com/gin-gonic/gin"
)

type Order_res struct {
	ID                 int               `json:"id"`
	TOTAL              float32           `json:"total"`
	CUSTOMER_ID        int               `json:"customer_id"`
	ADDRESSS_ID        int               `json:"addresss_id"`
	ORDER_STATUS_ID    int               `json:"order_status_id"`
	PAYMENT_METHODS_ID int               `json:"payment_methods_id"`
	ORDER_ITEMS        []Order_items_res `json:"order_items"`
}

type Order_items_res struct {
	ID          int       `json:"id"`
	QUANTITY    int       `json:"quantity"`
	PRODUCTS_ID int       `json:"products_id"`
	ORDERS_ID   int       `json:"orders_id"`
	PRODUCTS    []Product `json:"products"`
}

type Product struct {
	ID          int              `json:"id"`
	NAME        string           `json:"name"`
	DESCRIPTION string           `json:"description"`
	PRICE       float32          `json:"price"`
	PRODUCT_SKU string           `json:"product_sku"`
	QUANTITY    int              `json:"quantity"`
	CATEGORY_ID int              `json:"categories_id"`
	DISCOUNT_ID int              `json:"discounts_id"`
	BRAND_ID    int              `json:"brands_id"`
	PRODUCT_IMG []Product_images `json:"products_images"`
}

type Product_images struct {
	ID         int    `json:"id"`
	IMG_URL    string `json:"img_url"`
	PRODUCT_ID int    `json:"products_id"`
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

	ctx.Set("data", schema)
	ctx.Next()
}
