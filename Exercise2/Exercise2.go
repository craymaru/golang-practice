package main

import (
	"fmt"

	"./temp"
)

func main() {

	var beDivNum int = 20
	var divNum float64 = 0.002

	var printNum float64
	printNum = temp.DivideCall(beDivNum, divNum)
	fmt.Println(printNum)
}
