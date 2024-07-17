package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ravi11kumar/crypto-price-tracker/interfaces"
	"github.com/ravi11kumar/crypto-price-tracker/services"
	"net/http"
)

type PriceController struct {
	service interfaces.PriceService
}

func NewPriceController() *PriceController {
	return &PriceController{
		service: services.NewPriceService(),
	}
}

func (pc *PriceController) GetPrice(ctx *gin.Context) {
	price, err := pc.service.GetPrice()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": price})
}
