package main

import (
	"fmt"
	"net/http"
	"net/rpc"

	"example/rpc_test/common"
)

func main() {
	var service = new(common.MathService)

	err := rpc.Register(service)
	if err != nil {
		fmt.Printf("rpc server register service failed. error:%s", err)
	}

	rpc.HandleHTTP()
	fmt.Printf("server start...")
	err = http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("Listen and server is failed. error:%v\n", err)
	}
	fmt.Printf("server stop....")
}
