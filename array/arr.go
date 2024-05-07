package main

import "fmt"

func main() {
	fmt.Println("this is array file")

	var ar [5]string

	ar[0] = "kite"
	ar[1] = "knit"
	ar[2] = "we"
	ar[3] = "are"
	ar[4] = "the"

	fmt.Println(ar)

	var arrr = [3]string{
		"hi", "hello", "hey",
	}
	fmt.Println(arrr)

	var primes = [6]float32{2, 3, 5, 7.5, 11, 13}
	fmt.Println(primes)

}
