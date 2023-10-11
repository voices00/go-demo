package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.Writer.Write([]byte("hello"))
	})
	router.Run(":9999")
}