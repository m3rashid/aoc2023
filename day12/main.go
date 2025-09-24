package main

import (
	"fmt"
)

func main() {
	abs := "./input1.txt"

	output, _ := ReadInput(abs)
	fmt.Println(getSumCombinations(output, parseString))
	fmt.Println(getSumCombinations(output, parseString2)) // 50338344809230
}
