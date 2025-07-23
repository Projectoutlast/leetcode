package main

import "fmt"

func main() {
	var first string
	second := ""
	var third = new(string)

	fmt.Println(first, second, *third)
}
