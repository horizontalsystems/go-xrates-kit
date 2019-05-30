package currency

type Currency struct {
	ID            string
	Name          string
	SymbolUnicode string
	CurrencyType  currencyType
}
type currencyType int8

const (
	OTHER currencyType = iota
	DIGITAL
	FIAT
)

func GetByID(id string) Currency {
	return currencies[id]
}

func GetCurrencies(id string) map[string]Currency {
	return currencies
}

var currencies = map[string]Currency{

	"OTHER": Currency{"", "Other", "", OTHER},
	"BTC":   Currency{"", "Bitcoin", "U+20BF", DIGITAL},
	"ETH":   Currency{"", "Ethereum", "U+039E", DIGITAL},
	"BCH":   Currency{"", "Bitcoin-Cash", "U+20BF", DIGITAL},
	"DASH":  Currency{"", "Dash", "U+20BF", DIGITAL},
	"BAT":   Currency{"", "Basic-Attention-Token", "U+039E", DIGITAL},
	"BNB":   Currency{"", "Binance-coin", "U+039E", DIGITAL},
	"BNT":   Currency{"", "Bancor", "U+039E", DIGITAL},
	"CRO":   Currency{"", "Cryptocom-chain", "U+039E", DIGITAL},
	"DAI":   Currency{"", "Dai", "U+039E", DIGITAL},
	"DGD":   Currency{"", "Digixdao", "U+039E", DIGITAL},
	"DGX":   Currency{"", "Digix-gold-token", "U+039E", DIGITAL},
	"ELF":   Currency{"", "Aelf", "U+039E", DIGITAL},
	"ENJ":   Currency{"", "Enjin-coin", "U+039E", DIGITAL},
	"EURS":  Currency{"", "Stasis-eurs", "U+039E", DIGITAL},
	"GNT":   Currency{"", "Golem", "U+039E", DIGITAL},
	"GUSD":  Currency{"", "Gemini-dollar", "U+039E", DIGITAL},
	"HOT":   Currency{"", "Holo", "U+039E", DIGITAL},
	"HT":    Currency{"", "Huobi-Token", "U+039E", DIGITAL},

	"IDXM": Currency{"", "Idex-membership", "U+039E", DIGITAL},
	"KCS":  Currency{"", "Kucoin-shares", "U+039E", DIGITAL},
	"KNC":  Currency{"", "Kyber-Network", "U+039E", DIGITAL},
	"LINK": Currency{"", "ChainLink", "U+039E", DIGITAL},
	"LOOM": Currency{"", "Loom-network", "U+039E", DIGITAL},
	"LRC":  Currency{"", "Loopring", "U+039E", DIGITAL},
	"MANA": Currency{"", "Decentraland", "U+039E", DIGITAL},
	"MCO":  Currency{"", "mco", "U+039E", DIGITAL},
	"MITH": Currency{"", "Mithril", "U+039E", DIGITAL},
	"MKR":  Currency{"", "Maker", "U+039E", DIGITAL},
	"NEXO": Currency{"", "Nexo", "U+039E", DIGITAL},
	"NPXS": Currency{"", "Pundi-x", "U+039E", DIGITAL},
	"OMG":  Currency{"", "OmiseGO", "U+039E", DIGITAL},
	"ORBS": Currency{"", "Orbs", "U+039E", DIGITAL},
	"PAX":  Currency{"", "Paxos-standard-token", "U+039E", DIGITAL},
	"POLY": Currency{"", "Polymath", "U+039E", DIGITAL},
	"PPT":  Currency{"", "Populous", "U+039E", DIGITAL},
	"R":    Currency{"", "Revain", "U+039E", DIGITAL},
	"REP":  Currency{"", "Augur", "U+039E", DIGITAL},
	"SNT":  Currency{"", "Status", "U+039E", DIGITAL},
	"TUSD": Currency{"", "TrueUSD", "U+039E", DIGITAL},
	"USDC": Currency{"", "USD-Coin", "U+039E", DIGITAL},
	"USDT": Currency{"", "Tether", "U+039E", DIGITAL},
	"WAX":  Currency{"", "WAX", "U+039E", DIGITAL},
	"WTC":  Currency{"", "Waltonchain", "U+039E", DIGITAL},
	"ZIL":  Currency{"", "Zilliqa", "U+039E", DIGITAL},
	"ZRX":  Currency{"", "0x", "U+039E", DIGITAL},

	"USD": Currency{"", "US-Dollar", "U+0024", FIAT},
	"EUR": Currency{"", "Euro", "U+20AC", FIAT},
	"GBP": Currency{"", "GBP", "U+00A3", FIAT},
	"JPY": Currency{"", "JPY", "U+00A5", FIAT},
	"CAD": Currency{"", "CAD", "U+0024", FIAT},
	"AUD": Currency{"", "AUD", "U+20B3", FIAT},
	"CNY": Currency{"", "CNY", "U+00A5", FIAT},
	"CHF": Currency{"", "CHF", "U+20A3", FIAT},
	"RUB": Currency{"", "RUB", "U+20BD", FIAT},
	"KRW": Currency{"", "KRW", "U+20A9", FIAT},
	"TRY": Currency{"", "TRY", "U+20BA", FIAT}}
