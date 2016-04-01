package tcpserver

import (
	"fmt"
	"net"

	"github.com/grunmax/TestServers/data"
	"github.com/grunmax/TestServers/util"
)

const PROTOCOL = "tcp"

func Run(cfg *util.TcpConfig) {
	listener, err := net.Listen(PROTOCOL, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	util.Err("Error tcp listening:", err)
	util.Log(fmt.Sprintf("Listening %s on %s:%d", PROTOCOL, cfg.Host, cfg.Port), "")
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		util.Err("Error accepting:", err)
		go handler(connection, cfg.BufferSize, cfg.MinRunes)
	}
}

func handler(connection net.Conn, bufferSize int, minRunes int) {
	buffer := make([]byte, bufferSize)
	defer connection.Close()

	connWriteln := func(bytes, words int, isOk bool) {
		const ANSWER_OK = "tcp:ok:%d:%d\n"
		const ANSWER_ERR = "tcp:err:%d:%d\n"

		var err error
		response := ""
		if isOk {
			response = ANSWER_OK
		} else {
			response = ANSWER_ERR
		}
		_, err = connection.Write([]byte(fmt.Sprintf(response, bytes, words)))
		util.Log("Error tcp writing:", err)
	}

	connWrite := func(buflen, len_ int) {
		if len_ > 0 {
			connWriteln(buflen, len_, true)
		} else {
			connWriteln(buflen, len_, false)
		}
	}

	if bufferLength, err := connection.Read(buffer); err != nil {
		util.Log("Error tcp reading:", err)
	} else {
		inputData := string(buffer[:bufferLength-1])
		inputList := util.RegSplit(inputData, "[^\\S]+")
		okList, badList := util.WordsCheckList(inputList, minRunes)
		if len(badList) == 0 {
			util.Log("TCP words:", fmt.Sprintf("%v", inputList))
		} else {
			util.Log("TCP words:", fmt.Sprintf("%v not %v", inputList, badList))
		}
		data.Put(okList)
		connWrite(bufferLength, len(okList))
	}
}
