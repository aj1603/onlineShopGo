package customer

import (
	req "onlineshopgo/src/api/customer/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.POST("/create", req.Validate_create, create)
	router.POST("/update", req.Validate_update, update)
}
