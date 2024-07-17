package main

import (
	"github.com/ravi11kumar/crypto-price-tracker/handlers"
	"github.com/ravi11kumar/crypto-price-tracker/utils"
)

func main() {
	utils.LoadConfig()
	r := handlers.SetupRouter()
	err := r.Run(":" + utils.Config.Server.Port)
	if err != nil {
		return
	}
}
