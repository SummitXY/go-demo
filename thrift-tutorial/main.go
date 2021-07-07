package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	addr := flag.String("addr","localhost:8090","Address to listen to")
	server := flag.Bool("server",true,"Run server")

	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if *server {
		fmt.Println("开始运行Thrift Server")
		if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr); err != nil {
			fmt.Println("error running server:", err)
		}
	}
}
