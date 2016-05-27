package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/rbudiharso/gopro/middleware/auth"
	"os"
	"runtime"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func pong(c *gin.Context) {
	var lines []string
	inFile, _ := os.Open("/home/rbudiharso/Downloads/POINT2420160401122701.csv")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		md5sum := md5.Sum(scanner.Bytes())
		lines = append(lines, hex.EncodeToString(md5sum[:]))
	}
	c.JSON(200, lines)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := gin.New()
	server.Use(auth.Basic)
	server.GET("/", ping)
	server.GET("/ping", ping)
	server.GET("/pong", pong)
	server.Run()
}
