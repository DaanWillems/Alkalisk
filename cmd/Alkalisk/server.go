package main

import (
	"Alkalisk/api"
	"net/http"

	"github.com/DaanWillems/NanoRouter"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() {
	router := NanoRouter.NewRouter()

	router.SetStaticPath("../../static/")
	http.Handle("/", router)

	api.StartApi(router)

	http.ListenAndServe(":8380", nil)
}
