package main

import (
	"github.com/grunmax/TestServers/data"
	"github.com/grunmax/TestServers/httpserver"
	"github.com/grunmax/TestServers/tcpserver"
	"github.com/grunmax/TestServers/util"
)

var cfg *util.Config

func init() {
	util.InitLog()
	cfg = util.ReadConfig()
	data.Init(&cfg.Data)
}

func main() {
	go data.Run()
	go tcpserver.Run(&cfg.Tcp)
	httpserver.Run(&cfg.Http)
}
