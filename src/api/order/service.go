package order

import (
	"encoding/json"
	req "onlineshopgo/src/api/order/schemas"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var order req.Order

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &order)

	create_(&order)

	ctx.JSON(201, "Successfully created")
}

func get(ctx *gin.Context) {
	results, err := get_()

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, results)
}
