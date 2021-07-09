package simple_server

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/SummitXY/thrift-tutorial/gen-go/simple"
	"time"
)

func SimpleServer() {
	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,
		MaxFrameSize: 1024 * 256,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory := thrift.NewTTransportFactory()

	transport, _ := thrift.NewTServerSocket(":8090")

	processor := simple.NewSimpleServiceProcessor(&SimpleServiceHandler{})
	server := thrift.NewTSimpleServer4(processor,transport,transportFactory,protocolFactory)
	server.Serve()
}
