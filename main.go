package main

import (
	"fmt"
	utilsText "poc_go/utils"
)

func main() {
	fmt.Println("Hello Test")

	var name string = "Ice"
	age := 10

	hello := utilsText.GetHelloText(name)

	fmt.Println(name, age, hello)
}
