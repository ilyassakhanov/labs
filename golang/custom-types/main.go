package main

import "fmt"

type customString string

// Can extend structures using custom types
func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "Ilyas"
	name.log()
}
