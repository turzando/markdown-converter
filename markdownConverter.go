package main

import "fmt"

func main() {

	var text string

	fmt.Scan(&text)

	if (text == "") {
		fmt.Println("vazia")
	}

	fmt.Println(text)
}

