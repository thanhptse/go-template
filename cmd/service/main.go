package main

import (
	"flag"

	"github.com/thanhptse/go-template/config"
	"github.com/thanhptse/go-template/server"
	"go.uber.org/zap"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.Parse()

	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config failed")
		panic(err)
	}

	s, err := server.NewServer(cfg)
	if err != nil {
		zap.S().Errorf("Create server failed with err %v", err)
		panic(err)
	}

	s.Init()

	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("Start server failed with err %v", err)
		panic(err)
	}
}
