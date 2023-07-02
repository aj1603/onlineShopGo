package main

import (
	"fmt"
	"onlineshopgo/config"
	db "onlineshopgo/database"
	qr "onlineshopgo/qrcode"
	app "onlineshopgo/src"
	"onlineshopgo/src/tools"
)

func main() {
	qr.GenQR()
	config.Init_config()
	db.Init_db()
	tools.Init_queries()
	router := app.Init_app()
	address := fmt.Sprintf("%v:%v", config.ENV.HOST, config.ENV.PORT)
	router.Run(address)
}
