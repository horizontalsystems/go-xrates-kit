package config

//"github.com/BurntSushi/toml"

var conf Config

// Config is a main configuration Object
type Config struct {
	Ipfs        IpfsConfig
	CoinPaprika CoinPaprikaConfig
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

func init() {
	Load()
}

// Load config file
func Load() Config {

	// if _, err := toml.DecodeFile("../../cmd/config.toml", &conf); err != nil {
	// 	//fmt.Println(err)
	// }

	conf = Config{
		IpfsConfig{
			"https://ipfs-ext.horizontalsystems.xyz",
			"https://ipfs.io",
			"QmXTJZBMMRmBbPun6HFt3tmb3tfYF2usLPxFoacL7G5uMX",
		}, CoinPaprikaConfig{
			"https://api.coinpaprika.com/v1"}}

	return conf
}

// Get config object
func Get() Config {
	return conf
}
