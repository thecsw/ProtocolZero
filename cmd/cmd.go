package cmd

import "fmt"

func Run() {
	fmt.Println(greeting)
	option := getOption()
	switch option {
	case Reddit:
		cleanReddit()
	default:
		fmt.Println("Unknown service!")
	}
}
