package product

import (
	req "onlineshopgo/src/api/product/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.POST("/create", req.Validate_create, create)
	router.PUT("/update", req.Validate_update, update)
	router.DELETE("/remove/:id", req.Validate_delete, remove)
	router.GET("/get-by-id/:id", get_by_id)
	router.GET("/get-by-category-id/:id", get_by_category_id)
}
