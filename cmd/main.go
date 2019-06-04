package main

import (
	"fmt"

	"github.com/horizontalsystems/xrates-kit/pkg/config"
	"github.com/horizontalsystems/xrates-kit/pkg/handler"
	"github.com/horizontalsystems/xrates-kit/pkg/service"
	dtutils "github.com/horizontalsystems/xrates-kit/pkg/util/datetime"
)

func main() {

	t := dtutils.StrToTime("", "2019-06-04T05:45:26.371Z")

	//var conf = config.Get()

	tunix := t.UTC().Unix()

	xratesService := new(service.XRatesService)
	xratesService.Init()

	//fmt.Println("Getting Data:", xratesService.Get("BTC", "USD", "", t.Unix()))

	cpHandler := handler.CoinPaprika{&config.Get().CoinPaprika}
	resp, _ := cpHandler.GetXRatesAsJSON("BTC", "TRY", "", &tunix)
	fmt.Println("Getting Data:", resp)

}
