package main

import (
	res "./handle"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", res.HandleIndex())

	v1 := route.Group("api/v1")
	{
		v1.GET("/carfix", res.GetCarfix)
	}

	route.Run(":4055")
}
