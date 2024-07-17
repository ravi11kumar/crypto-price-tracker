package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ravi11kumar/crypto-price-tracker/controllers"
	"github.com/ravi11kumar/crypto-price-tracker/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())

	//Initialize controllers
	priceController := controllers.NewPriceController()

	initPriceRoutes(router, priceController)

	return router
}

// Initialize routes
func initPriceRoutes(router *gin.Engine, priceController *controllers.PriceController) {
	router.GET("/price", priceController.GetPrice)
}
