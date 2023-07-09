package src

import (
	"io"
	"onlineshopgo/config"
	"onlineshopgo/src/api/address"
	"onlineshopgo/src/api/brands"
	"onlineshopgo/src/api/cart"
	"onlineshopgo/src/api/category"
	"onlineshopgo/src/api/order"
	"onlineshopgo/src/api/product"
	"onlineshopgo/src/api/users"
	"onlineshopgo/src/tools"
	"os"

	"github.com/gin-gonic/gin"
)

func Init_app() *gin.Engine {
	set_mode(config.ENV.GIN_MODE)
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)
	router.Static("/public", config.ENV.DIR)
	router.Use(tools.Cors)
	set_routers(router)
	return router
}

func set_mode(mode string) {
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		file, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}

}

func set_routers(router *gin.Engine) {
	create_router(router, "product", product.Controller)
	create_router(router, "category", category.Controller)
	create_router(router, "order", order.Controller)
	create_router(router, "brand", brands.Controller)
	create_router(router, "user", users.Controller)
	create_router(router, "cart", cart.Controller)
	create_router(router, "address", address.Controller)
}

func create_router(
	router *gin.Engine,
	name string,
	controller func(router *gin.RouterGroup),
) {
	group_name := router.Group(name)
	controller(group_name)
}
