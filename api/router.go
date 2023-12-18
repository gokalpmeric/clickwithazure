package api

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.POST("/triggerJob", TriggerJobHandler)

	router.Run(":8080")
}
