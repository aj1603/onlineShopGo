package order

import (
	"encoding/json"
	"fmt"
	token "onlineshopgo/helpers"
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

func get_order_user_id(ctx *gin.Context) {
	claims, err := token.CheckToken(ctx)

	id := claims["user_id"].(float64)

	fmt.Println(id)

	results, _ := get_order_user_id_(id)

	if err != nil {
		ctx.JSON(500, "Zakaz tapylamady")
		return
	}

	ctx.JSON(200, results)
}
