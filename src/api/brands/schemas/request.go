package schemas

import (
	"encoding/json"
	"onlineshopgo/src/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Create struct {
	NAME    string `json:"name" validate:"required"`
	IMG_URL string `json:"img_url" validate:"required"`
}

type Update struct {
	ID      int    `json:"id" validate:"required,gt=0"`
	NAME    string `json:"name" validate:"required"`
	IMG_URL string `json:"img_url" validate:"required"`
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
