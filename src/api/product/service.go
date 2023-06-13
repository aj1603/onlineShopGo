package product

import (
	"encoding/json"
	req "onlineshopgo/src/api/product/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "5")
	intLimit, _ := strconv.Atoi(limit)
	offset := ctx.DefaultQuery("offset", "0")
	intOffset, _ := strconv.Atoi(offset)
	results, err := get_(intLimit, intOffset)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, results)
}

func create(ctx *gin.Context) {
	var product req.Create

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &product)

	create_(&product)

	ctx.JSON(201, "Successfully created")
}

func update(ctx *gin.Context) {
	var product req.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &product)

	update_(&product)

	ctx.JSON(200, "Successfully updated")
}

func remove(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	remove_(int_id)

	ctx.JSON(200, "Successfully delete")
}

func get_by_id(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	res := get_by_id_(int_id)

	ctx.JSON(200, res)
}

func get_by_category_id(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	res, _ := get_by_category_id_(int_id)

	ctx.JSON(200, res)
}
