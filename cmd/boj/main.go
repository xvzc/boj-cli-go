package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	var args = os.Args

	if len(args) == 1 {
	}

	for _, v := range args {
		fmt.Println(v)
	}
}
