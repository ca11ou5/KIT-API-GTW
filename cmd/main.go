package main

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := configs.InitConfig()

	// HTTP Server
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/*")
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")

	auth.InitRoutes(r, &cfg)
	err := r.Run(cfg.Port)
	if err != nil {
		log.Fatal("failed to run http server: " + err.Error())
	}
}
