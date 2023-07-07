package order

import (
	res "onlineshopgo/src/api/order/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all")
	router.POST("/order-create", res.Validate_order, create)
}
