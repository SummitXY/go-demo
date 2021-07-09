package simple_client

import (
	"context"
	"fmt"
	"github.com/SummitXY/thrift-tutorial/gen-go/simple"
)

var defaultCtx = context.Background()

func handleClient(client *simple.SimpleServiceClient) {
	res, _ := client.Add(defaultCtx,13,"25")
	fmt.Println("result is ", res)
}

func SimpleClient() {
}
