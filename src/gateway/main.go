package main

import (
	"fmt"
	"gosports/common/entity"
	"rsc.io/quote"
)

func main() {
	test := entity.TestObject{ Id:1, Name:"test"}
	fmt.Println(quote.Hello())
	fmt.Println(test)
}
