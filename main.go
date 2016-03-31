package main

import (
	"github.com/grunmax/TestServers/httpserver"
	"github.com/grunmax/TestServers/tcpserver"
	"github.com/grunmax/TestServers/util"
)

var cfgTcp *util.TcpConfig
var cfgHttp *util.HttpConfig

func init() {
	util.InitLog()
	cfgTcp, cfgHttp = util.ReadConfig()
}

func main() {
	go tcpserver.Run(cfgTcp)
	httpserver.Run(cfgHttp)
}
