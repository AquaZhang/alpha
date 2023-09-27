package routes

import (
	"main/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // 允许的前端地址
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))
	r.GET("/api/v1/resource/:type", handlers.GetResource)
	r.POST("/api/v1/resource", handlers.CreateResource)
	r.PUT("/api/v1/resource/:id", handlers.UpdateResource)
	r.DELETE("/api/v1/resource/:id", handlers.DeleteResource)

	return r
}
