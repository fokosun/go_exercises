package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(optional())
}

// the init and post statement are optional
// you can replace this block of code in the main function like so
func optional() int {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	return sum
}