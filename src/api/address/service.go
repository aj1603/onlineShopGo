package address

import (
	"encoding/json"
	"fmt"
	token "onlineshopgo/helpers"
	req "onlineshopgo/src/api/address/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	claims, err := token.CheckToken(ctx)

	id := claims["user_id"].(float64)

	fmt.Println(id)

	results, _ := get_(id)

	if err != nil {
		ctx.JSON(500, "Addres tapylmady")
		return
	}

	ctx.JSON(200, results)
}

func create(ctx *gin.Context) {
	var address req.Create

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &address)

	create_(&address)

	ctx.JSON(201, "Successfully created")
}

func update(ctx *gin.Context) {
	var address req.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &address)

	update_(&address)

	ctx.JSON(200, "Successfully updated")
}

func remove(ctx *gin.Context) {
	var address req.Delete

	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)
	address.ID = int_id

	remove_(&address)

	ctx.JSON(200, "Successfully delete")
}
