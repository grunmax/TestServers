package httpserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/grunmax/TestServers/util"
)

func Run(cfg *util.HttpConfig) {
	http.HandleFunc("/", rootHnd)
	http.HandleFunc("/favicon.ico", iconHndstub) // chrome extra request for root
	util.Log(fmt.Sprintf("Listening http on %s:%d", cfg.Host, cfg.Port), "")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nil)
	util.Err("Error http listening:", err)
}

func iconHndstub(w http.ResponseWriter, r *http.Request) {
	//
}

func rootHnd(w http.ResponseWriter, r *http.Request) {
	util.Log("HTTP request:", r.RequestURI)
	io.WriteString(w, "Hello")
}
