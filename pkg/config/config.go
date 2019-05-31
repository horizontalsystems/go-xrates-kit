package config

var conf Config

// Config is a main configuration Object
type Config struct {
	Ipfs          IpfsConfig
	CoinPaprika   CoinPaprikaConfig
	FiatXRatesAPI FiatXRatesAPIConfig
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

//FiatXRatesAPIConfig conf object
type FiatXRatesAPIConfig struct {
	APIURL string
}

func init() {
	Load()
}

// Load config file
func Load() Config {

	conf = Config{
		IpfsConfig{
			"https://ipfs-ext.horizontalsystems.xyz",
			"https://ipfs.io",
			"QmXTJZBMMRmBbPun6HFt3tmb3tfYF2usLPxFoacL7G5uMX"},
		CoinPaprikaConfig{
			"https://api.coinpaprika.com/v1"},
		FiatXRatesAPIConfig{
			"https://api.exchangeratesapi.io"}}

	return conf
}

// Get config object
func Get() Config {
	return conf
}
