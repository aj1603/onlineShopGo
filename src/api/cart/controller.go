package cart

import (
	res "onlineshopgo/src/api/order/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.POST("/order-create", res.Validate_order, create)
	router.GET("/get-by-user-id/:id", get_order_user_id)
}
