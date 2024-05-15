package main

import "fmt"

func main() {
	ip, err := parseIPv4("172.168.5.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\n", ip)
}
