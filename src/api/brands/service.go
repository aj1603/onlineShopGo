package brands

import (
	"encoding/json"
	"fmt"
	res "onlineshopgo/src/api/brands/schemas"
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
	var brand res.Create

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &brand)
	fmt.Println(byte_data)

	create_(&brand)

	ctx.JSON(201, "Successfully created")
}

func update(ctx *gin.Context) {
	var brand res.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &brand)

	update_(&brand)

	ctx.JSON(201, "Successfully updated")
}

func remove(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	remove_(int_id)

	ctx.JSON(201, "Successfully deleted")
}
