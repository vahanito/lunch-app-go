package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*.html")
	router.Static("/static", "./web/static")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/api/restaurants", test)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.Run()
}

func test(context *gin.Context) {
	res := []string{"foo", "bar"}
	context.JSON(200, res)
}
