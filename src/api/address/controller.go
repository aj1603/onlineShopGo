package address

import (
	req "onlineshopgo/src/api/address/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.GET("/get-by-category-id/:id", get_by_category_id)
	router.POST("/create", req.Validate_create, create)
	router.PUT("/update/:id", req.Validate_update, update)
	router.DELETE("/remove/:id", req.Validate_delete, remove)
}