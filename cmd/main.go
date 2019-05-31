package main

import (
	"fmt"
	"time"

	"github.com/horizontalsystems/xrates-kit/pkg/service/xrates"
)

func main() {

	layout := "2006-01-02T15:04:05.000Z"
	str := "2017-05-01T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	t = t.UTC()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	fmt.Println("Getting Data:", xrates.Get("BTC", "USD", "", t.Unix()))

}
