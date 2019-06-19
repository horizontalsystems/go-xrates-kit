package main

import (
	"fmt"
	"log"

	//"time"

	"github.com/horizontalsystems/go-xrates-kit/kit"

	dtutils "github.com/horizontalsystems/go-xrates-kit/util/datetime"
)

func main() {

	t := dtutils.StrToTime("", "2019-06-04T05:45:26.371Z")

	fmt.Println(t)
	fmt.Println(t.Unix())
	//var conf = config.Get()

	xratesKit := new(kit.XRatesKit)
	xratesKit.Init(".")

	//fmt.Println("Getting Data Cached:", )

	coinCodes := kit.CoinCodes{}
	coinCodes.FromStringArray([]string{"BTC", "BCH", "ETH", "DASH"})


	listener := RatesListener{}
	xratesKit.SubscribeToLatestRates(10, "USD", &coinCodes, &listener)

	//xratesKit.GetLatest("TRY", &coinCodes)

	//xratesKit.GetLatestCached("TRY", &coinCodes)
	//tunix := t.UTC().Unix()
	//cpHandler := handler.CoinPaprika{&config.Load().CoinPaprika}
	//resp, _ := cpHandler.GetLatestXRates ("TRY", &[]string{"BTC","BCH","ETH"} )
	//fmt.Println("Getting Data:", resp)

	_, _ = fmt.Scanln()

	listener2 := RatesListener2{}
	xratesKit.SubscribeToLatestRates(10, "EUR", &coinCodes, &listener2)

	_, _ = fmt.Scanln()

}

type RatesListener struct{}

func (l *RatesListener) OnUpdate(rates *kit.XRates) {

	log.Print("*********** listener: got new rates ", rates)

}


type RatesListener2 struct{}

func (l *RatesListener2) OnUpdate(rates *kit.XRates) {

	log.Print("++++++++++++ listener2: got new rates ", rates)

}