package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func greet(ctx *gin.Context) {
	time := time.Now()
	message := "Hi " + ctx.Param("name")
	ctx.JSON(200, gin.H{"time": time, "greeting": message})
}

func main() {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	gin.SetMode(gin.ReleaseMode)

	router.GET("/:name", greet)
	router.Run(":5004")
}
