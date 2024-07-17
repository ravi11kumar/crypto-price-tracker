package interfaces

import "github.com/ravi11kumar/crypto-price-tracker/models"

type PriceService interface {
	GetPrice() (*models.Price, error)
}
