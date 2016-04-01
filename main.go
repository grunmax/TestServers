package main

import (
	"github.com/grunmax/TestServers/data"
	"github.com/grunmax/TestServers/httpserver"
	"github.com/grunmax/TestServers/tcpserver"
	"github.com/grunmax/TestServers/util"
)

var cfgTcp *util.TcpConfig
var cfgHttp *util.HttpConfig
var cfgData *util.DataConfig

func init() {
	util.InitLog()
	cfgTcp, cfgHttp, cfgData = util.ReadConfig()
	data.Init(cfgData)
}

func main() {
	go data.Run()
	go tcpserver.Run(cfgTcp)
	httpserver.Run(cfgHttp)
}
