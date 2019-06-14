package models

type Currency struct {
	ID            string
	Name          string
	SymbolUnicode string
}


func GetBaseCurreny() Currency {
	return CURRENCIES["USD"]
}

var CURRENCIES = map[string]Currency{

	"USD": Currency{"USD", "US-Dollar", "U+0024"},
	"EUR": Currency{"EUR", "Euro", "U+20AC"},
	"GBP": Currency{"GBP", "GBP", "U+00A3"},
	"JPY": Currency{"JPY", "JPY", "U+00A5"},
	"CAD": Currency{"CAD", "CAD", "U+0024"},
	"AUD": Currency{"AUD", "AUD", "U+20B3"},
	"CNY": Currency{"CNY", "CNY", "U+00A5"},
	"CHF": Currency{"CHF", "CHF", "U+20A3"},
	"RUB": Currency{"RUB", "RUB", "U+20BD"},
	"KRW": Currency{"KRW", "KRW", "U+20A9"},
	"TRY": Currency{"TRY", "TRY", "U+20BA"}}
