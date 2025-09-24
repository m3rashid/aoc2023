package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	file := os.Args[1]
	fmt.Println(file)
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	Solution1(lines)
	Solution2(lines)
}
