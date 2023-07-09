package cart

import (
	res "onlineshopgo/src/api/cart/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.POST("/cart-create", res.Validate_cart, create)
	router.GET("/get-by-user-id/:id", get_cart_user_id)
	router.DELETE("/remove/:id", remove)
}
