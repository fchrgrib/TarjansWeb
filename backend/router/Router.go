package router

import (
	"backend/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"*"}
	r.Use(cors.New(config))

	r.POST("/post-graph", controller.Controller)
}
