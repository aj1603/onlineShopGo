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

type Favorite struct {
	ID          int `json:"id"`
	CUSTOMER_ID int `json:"customers_id"`
	PRODUCT_ID  int `json:"products_id"`
}

type Search struct {
	SEARCH string `json:"search" validate:"required"`
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

	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)
	schema.ID = int_id

	json.Unmarshal(data, &schema)
	errors := tools.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Set("data", schema)
	ctx.Next()
}

func Validate_search(ctx *gin.Context) {
	var schema Search

	search_word := ctx.Param("search")

	schema.SEARCH = search_word

	errors := tools.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Next()
}

func Validate_favorite(ctx *gin.Context) {
	var schema Favorite
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
