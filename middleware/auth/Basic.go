package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Basic(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.AbortWithStatus(401)
	}
	authHeaders := make([]string, 2)
	authHeaders = strings.Split(header, " ")
	if len(authHeaders) != 2 {
		c.AbortWithStatus(400)
	}
	credential, err := base64.StdEncoding.DecodeString(strings.TrimSpace(authHeaders[1]))
	if err != nil {
		c.AbortWithError(400, err)
	}
	usernamePassword := strings.Split(string(credential), ":")
	if strings.TrimSpace(usernamePassword[0]) != "Basic" {
		c.AbortWithStatus(401)
	}

	fmt.Println("=============================")
	fmt.Println("Scheme:", authHeaders[0])
	fmt.Println("credential:", string(credential[:]))
	c.Next()
}
