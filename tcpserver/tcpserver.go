package tcpserver

import (
	"github.com/grunmax/TestServers/util"
	"fmt"
	"net"
)

const (
	PROTOCOL = "tcp"
)

func Run(cfg *util.TcpConfig) {
	listener, err := net.Listen(PROTOCOL, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	util.Err("Error listening:", err)
	fmt.Println(fmt.Sprintf("Listening %s on %s:%d", PROTOCOL, cfg.Host, cfg.Port))
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		util.Err("Error accepting:", err)
		go handleReq(connection, cfg.BufferSize, cfg.MinBuffersize)
	}
}

func handleReq(connection net.Conn, bufferSize int, minBuffersize int) {
	buffer := make([]byte, bufferSize)
	defer connection.Close()

	connWrite := func(code int, isOk bool) {
		const ANSWER_OK = "tcp:ok:%d"
		const ANSWER_ERR = "tcp:err:%d"

		var err error
		response := ""
		if isOk {
			response = ANSWER_OK
		} else {
			response = ANSWER_ERR
		}
		_, err = connection.Write([]byte(fmt.Sprintf(response, code)))
		util.Log("Error tcp writing:", err)
	}

	if bufferLength, err := connection.Read(buffer); err != nil {
		util.Log("Error tcp reading:", err)
	} else if bufferLength < minBuffersize {
		util.Log("Wrong tcp data length:", bufferLength)
		connWrite(bufferLength, false)
	} else {
		inputData := string(buffer[:bufferLength-3])
		util.Log("TCP received:", inputData+"#")
		connWrite(bufferLength, true)
	}

}
