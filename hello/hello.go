package main

import (
	"fmt"
	"sync/atomic"
)
import "rsc.io/quote"

func main() {
	a := int32(0)
	atomic.AddInt32(&a, 1)
	fmt.Println(a)
	atomic.AddInt32(&a, 1)
	fmt.Println(a)
	fmt.Println("hello world!")
	fmt.Println(quote.Go())
}
