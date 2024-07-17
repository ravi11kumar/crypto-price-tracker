package services

import (
	"encoding/json"
	"fmt"
	"github.com/ravi11kumar/crypto-price-tracker/models"
	"github.com/ravi11kumar/crypto-price-tracker/utils"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type PriceService struct {
	cache      *models.Price
	lastUpdate time.Time
	mu         sync.Mutex
}

func NewPriceService() *PriceService {
	return &PriceService{}
}

func (s *PriceService) GetPrice() (*models.Price, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Since(s.lastUpdate) < utils.ExpiryDuration {
		return s.cache, nil
	}

	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data from Coin Desk API %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	price, err := models.NewPriceFromApiResponse(data)
	if err != nil {
		return nil, err
	}
	s.cache = price
	s.lastUpdate = time.Now()

	return s.cache, nil
}
