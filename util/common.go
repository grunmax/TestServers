package util

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
)

const INI_FILE = "testservers.ini"
const MIN_RUNES = 2

var errorlog *os.File
var Logger *log.Logger

type Config struct {
	Tcp  TcpConfig
	Http HttpConfig
	Data DataConfig
}

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

type DataConfig struct {
	BufferSize int
	Debug      bool
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
		Logger.Panic(s)
	}
}

func InitLog() {
	logg, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	Logger = log.New(logg, "TestServers: ", log.LstdFlags)
}

func ReadConfig() *Config {
	cfg := new(Config)
	ini_, err := ini.Load([]byte(""), INI_FILE)
	Err("no ini file", err)

	cfg.Tcp.Host = ini_.Section("tcp").Key("host").String()
	cfg.Tcp.Port, err = ini_.Section("tcp").Key("port").Int()
	Err("Wrong ini-value for tcp port", err)
	cfg.Tcp.BufferSize, err = ini_.Section("tcp").Key("buffersize").Int()
	Err("Wrong ini-value for buffersize", err)
	cfg.Tcp.MinRunes, err = ini_.Section("tcp").Key("minrunes").Int()
	Err("Wrong ini-value for minrunes", err)
	if cfg.Tcp.MinRunes < MIN_RUNES {
		Err(fmt.Sprintf("Minimal value for word symbols is %d", MIN_RUNES), "")
	}

	cfg.Http.Host = ini_.Section("http").Key("host").String()
	cfg.Http.Port, err = ini_.Section("http").Key("port").Int()
	Err("Wrong ini-value for http port", err)

	cfg.Data.BufferSize, err = ini_.Section("data").Key("buffersize").Int()
	Err("Wrong ini-value for data buffersize", err)
	cfg.Data.Debug, err = ini_.Section("data").Key("debug").Bool()
	Err("Wrong ini-value for data debug", err)

	return cfg
}
