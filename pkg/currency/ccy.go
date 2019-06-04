package currency

import (
	"strings"
)

var (
	fiatCurrencies    []string
	digitalCurrencies []string
	baseFiatCurrency  Currency
)

type Currency struct {
	ID            string
	Name          string
	SymbolUnicode string
	Type          currencyType
}
type currencyType int8

const (
	OTHER currencyType = iota
	DIGITAL
	FIAT
)

func init() {
	for k, v := range currencies {
		if v.Type == DIGITAL {
			digitalCurrencies = append(digitalCurrencies, k)
		} else if v.Type == FIAT {
			fiatCurrencies = append(fiatCurrencies, k)
		}
	}
	baseFiatCurrency = currencies["USD"]
}

func GetBaseFiatCurrency() Currency {
	return baseFiatCurrency
}

func GetDigitalCurrencies() []string {
	return digitalCurrencies
}

func GetFiatCurrencies() []string {
	return fiatCurrencies
}

func GetByID(id string) Currency {
	return currencies[id]
}

func GetFullName(id string) string {
	return strings.ToLower(id + "-" + currencies[id].Name)
}

func GetCurrencies(id string) map[string]Currency {
	return currencies
}

var currencies = map[string]Currency{

	"OTHER": Currency{"OTHER", "Other", "", OTHER},
	"BTC":   Currency{"BTC", "Bitcoin", "U+20BF", DIGITAL},
	"ETH":   Currency{"ETH", "Ethereum", "U+039E", DIGITAL},
	"BCH":   Currency{"BCH", "Bitcoin-Cash", "U+20BF", DIGITAL},
	"DASH":  Currency{"DASH", "Dash", "U+20BF", DIGITAL},
	"BAT":   Currency{"BAT", "Basic-Attention-Token", "U+039E", DIGITAL},
	"BNB":   Currency{"BNB", "Binance-coin", "U+039E", DIGITAL},
	"BNT":   Currency{"BNT", "Bancor", "U+039E", DIGITAL},
	"CRO":   Currency{"CRO", "Cryptocom-chain", "U+039E", DIGITAL},
	"DAI":   Currency{"DAI", "Dai", "U+039E", DIGITAL},
	"DGD":   Currency{"DGD", "Digixdao", "U+039E", DIGITAL},
	"DGX":   Currency{"DGX", "Digix-gold-token", "U+039E", DIGITAL},
	"ELF":   Currency{"ELF", "Aelf", "U+039E", DIGITAL},
	"ENJ":   Currency{"ENJ", "Enjin-coin", "U+039E", DIGITAL},
	"EURS":  Currency{"EURS", "Stasis-eurs", "U+039E", DIGITAL},
	"GNT":   Currency{"GNT", "Golem", "U+039E", DIGITAL},
	"GUSD":  Currency{"GUSD", "Gemini-dollar", "U+039E", DIGITAL},
	"HOT":   Currency{"HOT", "Holo", "U+039E", DIGITAL},
	"HT":    Currency{"HT", "Huobi-Token", "U+039E", DIGITAL},

	"IDXM": Currency{"IDXM", "Idex-membership", "U+039E", DIGITAL},
	"KCS":  Currency{"KCS", "Kucoin-shares", "U+039E", DIGITAL},
	"KNC":  Currency{"KNC", "Kyber-Network", "U+039E", DIGITAL},
	"LINK": Currency{"LINK", "ChainLink", "U+039E", DIGITAL},
	"LOOM": Currency{"LOOM", "Loom-network", "U+039E", DIGITAL},
	"LRC":  Currency{"LRC", "Loopring", "U+039E", DIGITAL},
	"MANA": Currency{"MANA", "Decentraland", "U+039E", DIGITAL},
	"MCO":  Currency{"MCO", "mco", "U+039E", DIGITAL},
	"MITH": Currency{"MITH", "Mithril", "U+039E", DIGITAL},
	"MKR":  Currency{"MKR", "Maker", "U+039E", DIGITAL},
	"NEXO": Currency{"NEXO", "Nexo", "U+039E", DIGITAL},
	"NPXS": Currency{"NPXS", "Pundi-x", "U+039E", DIGITAL},
	"OMG":  Currency{"OMG", "OmiseGO", "U+039E", DIGITAL},
	"ORBS": Currency{"ORBS", "Orbs", "U+039E", DIGITAL},
	"PAX":  Currency{"PAX", "Paxos-standard-token", "U+039E", DIGITAL},
	"POLY": Currency{"POLY", "Polymath", "U+039E", DIGITAL},
	"PPT":  Currency{"PPT", "Populous", "U+039E", DIGITAL},
	"R":    Currency{"R", "Revain", "U+039E", DIGITAL},
	"REP":  Currency{"REP", "Augur", "U+039E", DIGITAL},
	"SNT":  Currency{"SNT", "Status", "U+039E", DIGITAL},
	"TUSD": Currency{"TUSD", "TrueUSD", "U+039E", DIGITAL},
	"USDC": Currency{"USDC", "USD-Coin", "U+039E", DIGITAL},
	"USDT": Currency{"USDT", "Tether", "U+039E", DIGITAL},
	"WAX":  Currency{"WAX", "WAX", "U+039E", DIGITAL},
	"WTC":  Currency{"WTC", "Waltonchain", "U+039E", DIGITAL},
	"ZIL":  Currency{"ZIL", "Zilliqa", "U+039E", DIGITAL},
	"ZRX":  Currency{"ZRX", "0x", "U+039E", DIGITAL},

	"USD": Currency{"USD", "US-Dollar", "U+0024", FIAT},
	"EUR": Currency{"EUR", "Euro", "U+20AC", FIAT},
	"GBP": Currency{"GBP", "GBP", "U+00A3", FIAT},
	"JPY": Currency{"JPY", "JPY", "U+00A5", FIAT},
	"CAD": Currency{"CAD", "CAD", "U+0024", FIAT},
	"AUD": Currency{"AUD", "AUD", "U+20B3", FIAT},
	"CNY": Currency{"CNY", "CNY", "U+00A5", FIAT},
	"CHF": Currency{"CHF", "CHF", "U+20A3", FIAT},
	"RUB": Currency{"RUB", "RUB", "U+20BD", FIAT},
	"KRW": Currency{"KRW", "KRW", "U+20A9", FIAT},
	"TRY": Currency{"TRY", "TRY", "U+20BA", FIAT}}
