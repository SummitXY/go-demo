package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	simple_client "github.com/SummitXY/thrift-tutorial/simple-client"
	simple_server "github.com/SummitXY/thrift-tutorial/simple-server"
	"time"
)

func main() {
	Tutorial()
}

func Simple() {
	server := flag.Bool("server",true,"是否是服务器")
	flag.Parse()

	if *server {
		simple_server.SimpleServer()
	} else {
		simple_client.SimpleClient()
	}
}

func Tutorial() {
	addr := flag.String("addr","localhost:8090","Address to listen to")
	server := flag.Bool("server",true,"Run server")

	flag.Parse()

	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,

		MaxFrameSize: 1024 * 256,

		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)

	transportFactory := thrift.NewTTransportFactory()

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
