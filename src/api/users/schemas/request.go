package schemas

import (
	"encoding/json"
	"fmt"
	"onlineshopgo/src/tools"

	"github.com/gin-gonic/gin"
)

type Register struct {
	NAME      string `json:"name" validate:"required"`
	PASSWORD  string `json:"password" validate:"required"`
	EMAIL     string `json:"email" validate:"required"`
	PHONE_NUM int    `json:"phone_num" validate:"required"`
}

type Update struct {
	ID        int    `json:"id" validate:"required,gt=0"`
	NAME      string `json:"name" validate:"required"`
	PASSWORD  string `json:"password" validate:"required"`
	EMAIL     string `json:"email" validate:"required"`
	PHONE_NUM int    `json:"phone_num" validate:"required"`
}

func Validate_create(ctx *gin.Context) {
	var schema Register
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
