package product

import (
	req "onlineshopgo/src/api/product/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.POST("/create", req.Validate_create, create)
	router.PUT("/update/:id", req.Validate_update, update)
	router.DELETE("/remove/:id", remove)
	router.GET("/get-by-id/:id", get_by_id)
	router.GET("/search/:search", req.Validate_search, search)
	router.POST("/favorite", req.Validate_favorite, favorite)
	router.GET("/user/favorite", get_user_favorite)
}
