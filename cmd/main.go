package main

import (
	"fmt"
	//"time"

	"github.com/horizontalsystems/go-xrates-kit/kit"

	dtutils "github.com/horizontalsystems/go-xrates-kit/util/datetime"
)

func main() {

	t := dtutils.StrToTime("", "2019-06-04T05:45:26.371Z")

	fmt.Println(t.Unix())
	//var conf = config.Get()

	xratesKit := new(kit.XRatesKit)
	xratesKit.Init(".")

	//fmt.Println("Getting Data Cached:", )
	xratesKit.GetLatest("TRY", []string{"BTC", "BCH", "ETH"})

	xratesKit.GetLatestCached("TRY", []string{"BTC", "BCH", "ETH"})
	//tunix := t.UTC().Unix()
	//cpHandler := handler.CoinPaprika{&config.Load().CoinPaprika}
	//resp, _ := cpHandler.GetLatestXRates ("TRY", &[]string{"BTC","BCH","ETH"} )
	//fmt.Println("Getting Data:", resp)

	_, _ = fmt.Scanln()
}
