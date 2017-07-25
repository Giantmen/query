package main

import (
	"flag"
	stdlog "log"

	"github.com/Giantmen/query/config"
	"github.com/Giantmen/query/log"
	"github.com/Giantmen/query/store"

	"github.com/BurntSushi/toml"
	"github.com/solomoner/gozilla"
)

var (
	cfgPath = flag.String("config", "config.toml", "config file path")
)

func initLog(cfg *config.Config) {
	log.SetLevelByString(cfg.LogLevel)
	if !cfg.Debug {
		log.SetHighlighting(false)
		err := log.SetOutputByName(cfg.LogPath)
		if err != nil {
			log.Fatal(err)
		}
		log.SetRotateByDay()
	}
}

func main() {
	flag.Parse()
	var cfg config.Config
	_, err := toml.DecodeFile(*cfgPath, &cfg)
	if err != nil {
		stdlog.Fatal("DecodeConfigFile error: ", err)
	}
	initLog(&cfg)

	bourse, err := store.NewService(&cfg)
	if err != nil {
		log.Error("NewService err", err)
	}
	gozilla.RegisterService(bourse, "trader")
	log.Debug("register", "bourse")

	gozilla.DefaultLogOpt.Format += " {{.Body}}"
	stdlog.Fatal(gozilla.ListenAndServe(cfg.Listen))
}
