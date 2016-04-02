package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/grunmax/TestServers/data"
	"github.com/grunmax/TestServers/util"
)

const (
	ERR_PARAM_CODE   = 1
	ERR_PARAM_TEXT   = "wrong parameter value"
	MSG_NOPARAM_CODE = 1
	MSG_NOPARAM_TEXT = "no parameter"
)

func Run(cfg *util.HttpConfig) {
	http.HandleFunc("/", rootHnd)
	http.HandleFunc("/favicon.ico", iconHndstub) // chrome extra request for root
	http.HandleFunc("/top", gettopHnd)
	util.Log(fmt.Sprintf("Listening http on %s:%d", cfg.Host, cfg.Port), "")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nil)
	util.Err("Error http listening:", err)
}

func iconHndstub(w http.ResponseWriter, r *http.Request) {
	//stub
}

func rootHnd(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi")
}

func gettopHnd(w http.ResponseWriter, r *http.Request) {
	util.Log("HTTP request:", r.RequestURI)

	nFun := func() string {
		N := r.URL.Query().Get("N")
		n := r.URL.Query().Get("n")
		if len(N) != 0 {
			return N
		}
		if len(n) != 0 {
			return n
		}
		return ""
	}

	N := nFun()

	if len(N) != 0 {
		if n, err := strconv.Atoi(N); err != nil {
			json, err := util.ErrorJson(ERR_PARAM_CODE, ERR_PARAM_TEXT)
			if err != nil {
				http.Error(w, http.StatusText(503), 503)
			} else {
				io.WriteString(w, json) // wrong parameter
			}
		} else {
			json, err := util.TopWordsJson(data.GetTop(n))
			if err != nil {
				http.Error(w, http.StatusText(503), 503)
			} else {
				io.WriteString(w, json) // ok parameter
			}
		}
	} else {
		json, err := util.MessageJson(MSG_NOPARAM_CODE, MSG_NOPARAM_TEXT)
		if err != nil {
			http.Error(w, http.StatusText(503), 503)
		} else {
			io.WriteString(w, json) // no parameter
		}

	}
}
