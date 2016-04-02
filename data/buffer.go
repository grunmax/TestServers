package data

import (
	"fmt"

	"github.com/grunmax/TestServers/util"
)

var buff chan []string
var dataMap map[string]int
var debugMode bool

func Init(cfg *util.DataConfig) {
	debugMode = cfg.Debug
	buff = make(chan []string, cfg.BufferSize)
	dataMap = map[string]int{}
	util.Log("buffer size:", cfg.BufferSize)
}

func Put(list []string) {
	select {
	case buff <- list:
		if debugMode {
			fmt.Println("buffer <-", list)
		}
	default:
		util.Log("buffer is full !!!", list)
	}
}

func Run() {
	for {
		list := <-buff
		if debugMode {
			fmt.Println("received:", list)
		}
		save(list)
	}
}

func save(list []string) {
	for _, word := range list {
		count := dataMap[word]
		if count == 0 {
			dataMap[word] = 1
		} else {
			count += 1
			dataMap[word] = count
		}
	}
	if debugMode {
		fmt.Println(dataMap)
		fmt.Println("top 3:", getTopByValue(dataMap, 3))
	}
}

func GetTop(n int) []string {
	return getTopByValue(dataMap, n)
}
