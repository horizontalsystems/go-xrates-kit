package main

import (
	"fmt"

	//"time"

	"github.com/horizontalsystems/go-xrates-kit/config"
	"github.com/horizontalsystems/go-xrates-kit/handler"
	"github.com/horizontalsystems/go-xrates-kit/kit"

	dtutils "github.com/horizontalsystems/go-xrates-kit/util/datetime"
)

func main() {

	t := dtutils.StrToTime("", "2019-06-04T05:45:26.371Z")

	//var conf = config.Get()

	tunix := t.UTC().Unix()

	xratesKit := new(kit.XRatesKit)
	xratesKit.Init(".")

	fmt.Println("Getting Data:", xratesKit.Get("BTC", "TRY", "", t.Unix()))

	cpHandler := handler.CoinPaprika{&config.Load().CoinPaprika}
	resp, _ := cpHandler.GetHistoricalXRates("BTC", "TRY", "", &tunix)
	fmt.Println("Getting Data:", resp)

	_, _ = fmt.Scanln()
}
