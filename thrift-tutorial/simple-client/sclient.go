package simple_client

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/SummitXY/thrift-tutorial/gen-go/simple"
	"time"
)

var defaultCtx = context.Background()

func handleClient(client *simple.SimpleServiceClient) {
	res, _ := client.Add(defaultCtx,13,"25")
	fmt.Println("result is ", res)
}

func SimpleClient() {
	var transport thrift.TTransport
	transport, _ = thrift.NewTSocketConf("localhost:8090", &thrift.TConfiguration{
		ConnectTimeout: time.Second, // Use 0 for no timeout
		SocketTimeout:  time.Second, // Use 0 for no timeout
	})

	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,

		MaxFrameSize: 1024 * 256,

		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)

	transportFactory := thrift.NewTTransportFactory()
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	transport.Open()

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	handleClient(simple.NewSimpleServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
