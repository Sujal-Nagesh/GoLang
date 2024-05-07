package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is a ptr file")

	var pt = 12
	var a = &pt
	fmt.Println("this is new", a)
	*a = *a + 2
	fmt.Println(pt)

}
