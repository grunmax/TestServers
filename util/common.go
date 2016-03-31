package util

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
)

var errorlog *os.File
var Logger *log.Logger

type TcpConfig struct {
	Host       string
	Port       int
	BufferSize int
	MinRunes   int
}

type HttpConfig struct {
	Host string
	Port int
}

func Log(userMessage string, v interface{}) {
	if v != nil {
		s := fmt.Sprintf(":%s  %v\n", userMessage, v)
		fmt.Printf(s)
		Logger.Printf(s)
	}
}

func Err(userMessage string, e interface{}) {
	if e != nil {
		s := fmt.Sprintf("ERROR:%s  %v\n", userMessage, e)
		fmt.Printf(s)
		Logger.Printf(s)
		os.Exit(1)
	}
}

func ErrPanic(userMessage string, e interface{}) {
	if e != nil {
		s := fmt.Sprintf("ERROR:%s  %v\n", userMessage, e)
		fmt.Printf(s)
		Logger.Panicf(s)
	}
}

func InitLog() {
	logg, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	Logger = log.New(logg, "TestServers: ", log.LstdFlags)
}

func ReadConfig() (*TcpConfig, *HttpConfig) {
	const MIN_RUNES = 2

	iniFile := "testservers.ini"
	tcpCfg := new(TcpConfig)
	httpCfg := new(HttpConfig)
	cfg, err := ini.Load([]byte(""), iniFile)
	Err("no ini file", err)

	tcpCfg.Host = cfg.Section("tcp").Key("host").String()
	tcpCfg.Port, err = cfg.Section("tcp").Key("port").Int()
	Err("Wrong ini-value for tcp port", err)
	tcpCfg.BufferSize, err = cfg.Section("tcp").Key("buffersize").Int()
	Err("Wrong ini-value for buffersize", err)
	tcpCfg.MinRunes, err = cfg.Section("tcp").Key("minrunes").Int()
	Err("Wrong ini-value for minrunes", err)
	if tcpCfg.MinRunes < MIN_RUNES {
		Err(fmt.Sprintf("Minimal value for word symbols is %d", MIN_RUNES), "")
	}

	httpCfg.Host = cfg.Section("http").Key("host").String()
	httpCfg.Port, err = cfg.Section("http").Key("port").Int()
	Err("Wrong ini-value for http port", err)

	return tcpCfg, httpCfg
}
