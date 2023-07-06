package brands

import (
	req "onlineshopgo/src/api/brands/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.POST("/create", req.Validate_create, create)
	router.PUT("/update/:id", req.Validate_update, update)
	router.DELETE("/remove/:id", remove)
}
