package models

import (
	"strings"
)

type Coin struct {
	ID            string
	Name          string
	SymbolUnicode string
}


func (coin *Coin) GetFullName() string {
	return strings.ToLower(coin.ID + "-" + coin.Name)
}



var COINS = map[string]Coin{

	"OTHER": Coin{"OTHER", "Other", ""},
	"BTC":   Coin{"BTC", "BitCoin", "U+20BF"},
	"ETH":   Coin{"ETH", "Ethereum", "U+039E"},
	"BCH":   Coin{"BCH", "BitCoin-Cash", "U+20BF"},
	"DASH":  Coin{"DASH", "Dash", "U+20BF"},
	"BAT":   Coin{"BAT", "Basic-Attention-Token", "U+039E"},
	"BNB":   Coin{"BNB", "Binance-Coin", "U+039E"},
	"BNT":   Coin{"BNT", "Bancor", "U+039E"},
	"CRO":   Coin{"CRO", "Cryptocom-chain", "U+039E"},
	"DAI":   Coin{"DAI", "Dai", "U+039E"},
	"DGD":   Coin{"DGD", "Digixdao", "U+039E"},
	"DGX":   Coin{"DGX", "Digix-gold-token", "U+039E"},
	"ELF":   Coin{"ELF", "Aelf", "U+039E"},
	"ENJ":   Coin{"ENJ", "Enjin-Coin", "U+039E"},
	"EURS":  Coin{"EURS", "Stasis-eurs", "U+039E"},
	"GNT":   Coin{"GNT", "Golem", "U+039E"},
	"GUSD":  Coin{"GUSD", "Gemini-dollar", "U+039E"},
	"HOT":   Coin{"HOT", "Holo", "U+039E"},
	"HT":    Coin{"HT", "Huobi-Token", "U+039E"},

	"IDXM": Coin{"IDXM", "Idex-membership", "U+039E"},
	"KCS":  Coin{"KCS", "KuCoin-shares", "U+039E"},
	"KNC":  Coin{"KNC", "Kyber-Network", "U+039E"},
	"LINK": Coin{"LINK", "ChainLink", "U+039E"},
	"LOOM": Coin{"LOOM", "Loom-network", "U+039E"},
	"LRC":  Coin{"LRC", "Loopring", "U+039E"},
	"MANA": Coin{"MANA", "Decentraland", "U+039E"},
	"MCO":  Coin{"MCO", "mco", "U+039E"},
	"MITH": Coin{"MITH", "Mithril", "U+039E"},
	"MKR":  Coin{"MKR", "Maker", "U+039E"},
	"NEXO": Coin{"NEXO", "Nexo", "U+039E"},
	"NPXS": Coin{"NPXS", "Pundi-x", "U+039E"},
	"OMG":  Coin{"OMG", "OmiseGO", "U+039E"},
	"ORBS": Coin{"ORBS", "Orbs", "U+039E"},
	"PAX":  Coin{"PAX", "Paxos-standard-token", "U+039E"},
	"POLY": Coin{"POLY", "Polymath", "U+039E"},
	"PPT":  Coin{"PPT", "Populous", "U+039E"},
	"R":    Coin{"R", "Revain", "U+039E"},
	"REP":  Coin{"REP", "Augur", "U+039E"},
	"SNT":  Coin{"SNT", "Status", "U+039E"},
	"TUSD": Coin{"TUSD", "TrueUSD", "U+039E"},
	"USDC": Coin{"USDC", "USD-Coin", "U+039E"},
	"USDT": Coin{"USDT", "Tether", "U+039E"},
	"WAX":  Coin{"WAX", "WAX", "U+039E"},
	"WTC":  Coin{"WTC", "Waltonchain", "U+039E"},
	"ZIL":  Coin{"ZIL", "Zilliqa", "U+039E"},
	"ZRX":  Coin{"ZRX", "0x", "U+039E"}}

