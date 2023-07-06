package address

import (
	"encoding/json"
	req "onlineshopgo/src/api/address/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	results, err := get_()

	if err != nil {
		ctx.JSON(500, err.Error())
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

func get_by_category_id(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	res := get_by_category_id_(int_id)

	ctx.JSON(200, res)
}
