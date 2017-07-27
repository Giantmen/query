package config

type Server struct {
	Name      string
	Accountid string
	Accesskey string
	Secretkey string
	Timeout   int
}

type Config struct {
	Listen string

	Debug    bool
	LogPath  string
	LogLevel string

	Chbtc    []Server
	Yunbi    []Server
	HuobiN   []Server
	HuobiO   []Server
	Btctrade []Server
	Bter     []Server
	Poloniex []Server
}
