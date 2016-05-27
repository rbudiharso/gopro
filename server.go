package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := gin.New()
	server.GET("/", ping)
	server.GET("/ping", ping)
	server.Run()
}
