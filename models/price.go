package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Price struct {
	Bitcoin map[string]string `json:"bitcoin"`
}

func NewPriceFromApiResponse(data map[string]interface{}) (*Price, error) {

	bpi, ok := data["bpi"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid response format")
	}
	bitcoin := make(map[string]string)
	for currency, value := range bpi {
		rateStr, ok := value.(map[string]interface{})["rate"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid API response format")
		}

		// Replace commas with empty strings to handle thousand separators
		rateStr = strings.ReplaceAll(rateStr, ",", "")

		rateFloat, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			return nil, err
		}
		bitcoin[currency] = fmt.Sprintf("%.4f", rateFloat)
	}
	return &Price{Bitcoin: bitcoin}, nil
}
