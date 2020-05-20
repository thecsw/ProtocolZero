package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh/terminal"
)

func show(arr []string, pre string, num int) {
	min := len(arr) - num
	if min < 0 {
		min = 0
	}
	i := 0
	for _, a := range arr[min:] {
		fmt.Printf("[%d] %s %s\n", min+i, pre, a)
		i++
	}
}

func clear() {
	cl := exec.Command("clear")
	cl.Stdout = os.Stdout
	cl.Run()
}

func getOption() (option Website) {
	fmt.Print(options)
	for {
		_, err := fmt.Scanf("%d", &option)
		if err == nil {
			return
		}
		fmt.Println("Error:", err)
		fmt.Print("OPTION:")
	}
	return
}

func getPass(prompt string) string {
	for {
		fmt.Printf(prompt)
		pass, err := terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err == nil {
			return string(pass)
		}
		fmt.Println("Input failed:", err, ", Try again!")
	}
}
