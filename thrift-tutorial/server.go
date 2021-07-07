package main

import (
	"fmt"
	"github.com/SummitXY/thrift-tutorial/gen-go/tutorial"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport

	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	fmt.Printf("%T\n", transport)
	handler := NewCalculatorHandler()
	processor := tutorial.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
