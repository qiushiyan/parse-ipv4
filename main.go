package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <ip address>")
		return
	}
	ip, err := parseIPv4(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s -> %d\n", os.Args[1], ip)
}
