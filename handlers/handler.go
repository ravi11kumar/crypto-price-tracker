package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ravi11kumar/crypto-price-tracker/routers"
)

func SetupRouter() *gin.Engine {
	return routers.InitRouter()
}
