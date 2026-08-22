package httpServer

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health",health)
	// r := gin.Default()
	return r
}