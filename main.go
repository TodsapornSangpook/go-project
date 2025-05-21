package main

import (
	"fmt"
	utils "poc_go/utils"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello Test")

	var name string = "Ice"
	age := 10
	weight := 40

	hello := utils.GetHelloText(name)
	ageAddWeight := utils.AddNumber(age, weight)

	fmt.Println(hello, ageAddWeight)

	fmt.Println("uuid :", uuid.New().String())

}
