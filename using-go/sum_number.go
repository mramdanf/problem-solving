package main

import "fmt"

func sum_number(num ...int) {
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println(sum)
}