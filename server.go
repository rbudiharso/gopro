package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "PONG"})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := gin.New()
	server.GET("/", ping)
	server.GET("/pong", pong)
	server.Run()
}
