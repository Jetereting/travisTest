package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	ip, _ := httplib.Get("http://myexternalip.com/raw").String()
	fmt.Println("ip:::", ip)

	i := 0
	for {
		fmt.Println("和树金   ", i)
		i++
		time.Sleep(2 * time.Second)
	}

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
