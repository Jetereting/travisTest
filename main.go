package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
)

func main() {

	ip, _ := httplib.Get("http://myexternalip.com/raw").String()
	fmt.Println("ip:::", ip)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ip, _ := httplib.Get("http://myexternalip.com/raw").String()

		c.JSON(200, gin.H{
			"message": "pong4",
			"ip":      ip,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
