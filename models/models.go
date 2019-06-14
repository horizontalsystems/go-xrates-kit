package models

import (
	"encoding/json"
)

// LatestXRate is a container for latest exchange rates
type LatestXRate struct {
	CurrencyCode string `json:"currency"`
	TimeStr      string `json:"time_str"`
	Time         int64  `json:"time"`

	Rates map[string]string `json:"rates"`
}

// HistoricalXRate is a container for exchange rates
type HistoricalXRate struct {
	CoinCode     string `json:"coin"`
	CurrencyCode string `json:"currency"`
	Rate         string `json:"rate"`
	Time         int64  `json:"time"`
}

func (rate *HistoricalXRate) ToJsonString() string {
	bytes, err := json.Marshal(rate)
	if err != nil {
		//TODO
	}

	return string(bytes)
}

func (rate *HistoricalXRate) FromJsonString(jsonData string) *HistoricalXRate {
	err := json.Unmarshal([]byte(jsonData), &rate)
	if err != nil {
		//TODO
	}
	return rate
}
