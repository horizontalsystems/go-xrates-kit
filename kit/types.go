package kit

import (
	"github.com/horizontalsystems/go-xrates-kit/models"
	"log"
)

// ****************************************************

type XRate models.XRate

// ****************************************************

type XRates struct{ xrates []XRate }

func (a *XRates) Get(index int) (xrate *XRate) {
	if index < 0 || index >= len(a.xrates) {
		log.Fatalf("index out of bounds")
	}
	return &a.xrates[index]
}

func (a *XRates) Size() int {
	return len(a.xrates)
}

// ****************************************************

type CoinCodes struct{ coinCodes []string }

func (c *CoinCodes) Init(size int) {
	c.coinCodes = make([]string, size)
}

func (c *CoinCodes) FromStringArray(arr []string) {
	c.coinCodes = make([]string, len(arr))

	for i, s := range arr {
		c.coinCodes[i] = s
	}
}

func (s *CoinCodes) Get(index int) (coinCode string) {
	if index < 0 || index >= len(s.coinCodes) {
		log.Fatalf("index out of bounds")
		return
	}
	return s.coinCodes[index]
}

func (s *CoinCodes) Size() int {
	return len(s.coinCodes)
}

func (s *CoinCodes) Set(index int, coinCode string) {
	if index < 0 || index >= len(s.coinCodes) {
		log.Fatalf("index out of bounds")
		return
	}
	s.coinCodes[index] = coinCode
}

func (s *CoinCodes) toStringArray() []string {
	result := make([]string, s.Size())

	for index, coinCode := range s.coinCodes {
		result[index] = coinCode
	}

	return result
}
