package main

import (
	"flag"

	"github.com/Giantmen/query/api"
	"github.com/Giantmen/query/config"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/solomoner/gozilla"
)

var (
	cfgPath = flag.String("config", ".config.toml", "config file path")
)

func main() {
	flag.Parse()
	glog.Infof("query start")

	var cfg config.Config
	_, err := toml.DecodeFile(*cfgPath, &cfg)
	if err != nil {
		glog.Fatal("DecodeConfigFile error: ", err)
	}

	bourse, err := store.NewService(&cfg)
	if err != nil {
		glog.Error("NewService err", err)
	}
	gozilla.RegisterService(bourse, "trader")
	glog.Infoln("register", "bourse")

	gozilla.DefaultLogOpt.Format += " {{.Body}}"
	glog.Fatal(gozilla.ListenAndServe(cfg.Listen))
}
