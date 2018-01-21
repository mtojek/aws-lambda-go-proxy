package main

import (
	"fmt"
        "net"
        "os"

	"github.com/kahlys/tcpproxy"
)

func main() {
	localAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%s", os.Getenv("_LAMBDA_SERVER_PORT")))
	if err != nil {
		panic(err)
	}

	remoteAddress, err := net.ResolveTCPAddr("tcp", os.Getenv("LAMBDA_DEBUG_PROXY"))
	if err != nil {
		panic(err)
	}

        tcpproxy.NewProxy(remoteAddress, nil, nil).ListenAndServe(localAddress)
}
