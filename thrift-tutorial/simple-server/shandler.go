package simple_server

import (
	"context"
	"strconv"
)

type SimpleServiceHandler struct {

}

func (ssh *SimpleServiceHandler) Add(ctx context.Context, num1 int32, num2 string) (_r int32, _err error) {
	num2int32, _ := strconv.Atoi(num2)
	return num1 + int32(num2int32), nil
}
