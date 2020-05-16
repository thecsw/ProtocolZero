package main

import (
	"fmt"
)

func main() {
	fmt.Println(greeting)
	option := getOption()
	switch option {
	case Reddit:
		cleanReddit()
	default:
		fmt.Println("Unknown service!")
	}
}
