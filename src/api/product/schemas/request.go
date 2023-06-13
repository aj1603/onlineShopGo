package schemas

import (
	"encoding/json"
	"onlineshopgo/src/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Create struct {
	NAME        string  `json:"name" validate:"required"`
	DESCRIPTION string  `json:"description" validate:"required"`
	PRICE       float32 `json:"price" validate:"required"`
	PRODUCT_SKU string  `json:"product_sku" validate:"required"`
	QUANTITY    int     `json:"quantity" validate:"required"`
	CATEGORY_ID int     `json:"categories_id" validate:"required"`
	DISCOUNT_ID int     `json:"discounts_id" validate:"required"`
	BRAND_ID    int     `json:"brands_id" validate:"required"`
}

type Update struct {
	ID          int     `json:"id" validate:"required,gt=0"`
	NAME        string  `json:"name" validate:"required"`
	DESCRIPTION string  `json:"description"`
	PRICE       float32 `json:"price"`
	PRODUCT_SKU string  `json:"product_sku"`
	QUANTITY    int     `json:"quantity"`
	CATEGORY_ID int     `json:"categories_id"`
	DISCOUNT_ID int     `json:"discounts_id"`
	BRAND_ID    int     `json:"brands_id"`
}

type Delete struct {
	ID int `json:"id" validate:"required,gt=0"`
}

func Validate_create(ctx *gin.Context) {
	var schema Create
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

func Validate_update(ctx *gin.Context) {
	var schema Update
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

func Validate_delete(ctx *gin.Context) {
	var schema Delete

	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)
	schema.ID = int_id

	errors := tools.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Next()
}
