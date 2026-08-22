package httpServer

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "Service is running smoothly",
		"time" : time.Now().UTC(),
	});
}