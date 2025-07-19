package main

import "fmt"

func main() {
	num := 10

	for num > 0 {
		fmt.Println(num % 10)
		num /= 10
	}
}
