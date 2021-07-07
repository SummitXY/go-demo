package main

import (
	"context"
	"fmt"
	"github.com/SummitXY/thrift-tutorial/gen-go/shared"
	"github.com/SummitXY/thrift-tutorial/gen-go/tutorial"
)

type CalculatorHandler struct {
	log map[int]*shared.SharedStruct
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{log: make(map[int]*shared.SharedStruct)}
}

func (p *CalculatorHandler) Ping(ctx context.Context) error {
	fmt.Println("pong")
	return nil
}

func (p *CalculatorHandler) Add(ctx context.Context, num1 int32, num2 int32) (int32, error) {
	fmt.Print("add(", num1, ",", num2, ")\n")
	return num1 + num2, nil
}

func (p *CalculatorHandler) Calculate(ctx context.Context, logid int32, w *tutorial.Work) (val int32, err error) {
	fmt.Print("calculate(", logid, ", {", w.Op, ",", w.Num1, ",", w.Num2, "})\n")

	switch w.Op {
	case tutorial.Operation_ADD:
		val = w.Num1 + w.Num2
		break
	case tutorial.Operation_SUBTRACT:
		val = w.Num1 - w.Num2
		break
	case tutorial.Operation_MULTIPLY:
		val = w.Num1 * w.Num2
	case tutorial.Operation_DIVIDE:
		val = w.Num1 / w.Num2
	}

	return val, nil
}

func (p *CalculatorHandler) GetStruct(ctx context.Context, key int32) (*shared.SharedStruct, error) {
	fmt.Print("getStruct(", key, ")\n")
	v , _ := p.log[int(key)]
	return v, nil
}
