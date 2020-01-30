package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vahanito/lunch-app-go/parsers/delpatio"
	"github.com/vahanito/lunch-app-go/parsers/drhunger"
	"github.com/vahanito/lunch-app-go/parsers/hotelboss"
	"github.com/vahanito/lunch-app-go/parsers/narohu"
	"github.com/vahanito/lunch-app-go/parsers/pivarium"
	"net/http"
)

func main() {
	router := gin.Default()
	setupStaticSources(router)

	router.GET("/api/restaurants", restaurantList)
	router.GET("/api/restaurants/del_patio", registerDelPatio)
	router.GET("/api/restaurants/pivarium", registerPivarium)
	router.GET("/api/restaurants/na_rohu", registerNaRohu)
	router.GET("/api/restaurants/dr_hunger", registerDrHunger)
	router.GET("/api/restaurants/hotel_boss", registerHotelBoss)

	router.Run()
}

func registerHotelBoss(context *gin.Context) {
	context.JSON(200, hotelboss.ParseHotelBoss())
}

func registerDrHunger(context *gin.Context) {
	context.JSON(200, drhunger.ParseDrHunger())
}

func registerNaRohu(context *gin.Context) {
	context.JSON(200, narohu.ParseNaRohu())
}

func registerPivarium(context *gin.Context) {
	context.JSON(200, pivarium.ParsePivarium())
}

func registerDelPatio(context *gin.Context) {
	context.JSON(200, delpatio.ParseDelPatio())
}

func setupStaticSources(router *gin.Engine) {
	router.LoadHTMLGlob("../ui/build/*.html")
	router.Static("/static", "../ui/build/static")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func restaurantList(context *gin.Context) {
	res := []string{"del_patio", "pivarium", "na_rohu", "dr_hunger", "hotel_boss"}
	context.JSON(200, res)
}
