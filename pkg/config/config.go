package config

var conf MainConfig

// MainConfig is a main configuration Object
type MainConfig struct {
	Ipfs        IpfsConfig
	CoinPaprika CoinPaprikaConfig
	FiatXRates  FiatXRatesConfig
}

//IpfsConfig conf object
type IpfsConfig struct {
	URL       string
	PublicURL string
	IpnsID    string
}

//CoinPaprikaConfig conf object
type CoinPaprikaConfig struct {
	APIURL string
}

//FiatXRatesConfig conf object
type FiatXRatesConfig struct {
	APIURL string
}

func init() {
	Load()
}

// Load config file
func Load() MainConfig {

	conf = MainConfig{
		IpfsConfig{
			"https://ipfs-ext.horizontalsystems.xyz",
			"https://ipfs.io",
			"QmXTJZBMMRmBbPun6HFt3tmb3tfYF2usLPxFoacL7G5uMX"},
		CoinPaprikaConfig{
			"https://api.coinpaprika.com/v1"},
		FiatXRatesConfig{
			"https://api.exchangeratesapi.io"}}

	return conf
}

// Get config object
func Get() *MainConfig {
	return &conf
}
