package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/SummitXY/thrift-tutorial/gen-go/tutorial"
)

var defaultCtx = context.Background()

func handleClient(client *tutorial.CalculatorClient) error {
	client.Ping(defaultCtx)
	fmt.Println("ping()")

	sum, _ := client.Add(defaultCtx, 1,10 )
	fmt.Print("1+10=", sum, "\n")

	work := tutorial.NewWork()
	work.Op = tutorial.Operation_SUBTRACT
	work.Num1 = 9
	work.Num2 = 3
	result, _ := client.Calculate(defaultCtx, 1, work)
	fmt.Println("Calculate value:", result)

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	defer transport.Close()
	if err = transport.Open(); err != nil {
		return err
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	return handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot)))
}