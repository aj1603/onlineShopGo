package address

import (
	req "onlineshopgo/src/api/address/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/user-address", get)
	router.POST("/create", req.Validate_create, create)
	router.PUT("/update/:id", req.Validate_update, update)
	router.DELETE("/remove/:id", remove)
}
