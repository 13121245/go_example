package main

import (
	"example/rpc_test/common"
	"fmt"
	"net/rpc"
)

func main() {
	var args = common.Args{A: 32, B: 64}
	var result = common.Result{}

	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("connect server failed. err:%s", err)
	}

	err = client.Call("MathService.Divide", args, &result)
	if err != nil {
		fmt.Printf("call server failed. err:%v", err)
	}
	fmt.Printf("call rpc server success, result:%f", result.Value)
}
