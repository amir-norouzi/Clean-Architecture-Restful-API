package server

import (
	"RestApi/config"
	"RestApi/http/controllers"
	"github.com/gin-gonic/gin"
)

var (
	//read app url from config file
	url = config.LoadConfig("APP_URL")
	//read app port from config file
	port = config.LoadConfig("APP_PORT")
)

func Run() {

	gin.SetMode(config.LoadConfig("APP_MODE"))
	r := gin.Default()
	//Program routes
	r.GET("/getuser", controllers.GetUsers())
	err := r.Run("" + url + ":" + port + "")
	if err != nil {
		return
	}
}
