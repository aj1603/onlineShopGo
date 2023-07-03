package users

import (
	req "onlineshopgo/src/api/users/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.POST("/register", req.Validate_create, register)
	router.POST("/login", req.Validate_update, login)
}
