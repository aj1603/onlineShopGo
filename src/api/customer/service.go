package customer

import (
	"encoding/json"
	req "onlineshopgo/src/api/customer/schemas"

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
	var customer req.Create

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &customer)

	create_(&customer)

	ctx.JSON(201, "Successfully created")
}

func update(ctx *gin.Context) {
	var customer req.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &customer)

	update_(&customer)

	ctx.JSON(200, "Successfully updated")
}
