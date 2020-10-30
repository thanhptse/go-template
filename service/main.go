package main

import (
	"github.com/thanhptse/go-template/server"
	"go.uber.org/zap"
)

func main() {
	s, err := server.NewServer(nil)
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
