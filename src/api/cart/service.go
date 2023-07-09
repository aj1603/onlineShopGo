package cart

import (
	"encoding/json"
	req "onlineshopgo/src/api/cart/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var cart req.Cart

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &cart)

	create_(&cart)

	ctx.JSON(201, "Successfully created")
}

func get_cart_user_id(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	results, err := get_cart_user_id_(int_id)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, results)
}

func remove(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	remove_(int_id)

	ctx.JSON(200, "Successfully delete")
}
