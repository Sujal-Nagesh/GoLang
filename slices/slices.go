package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("this is slices file")

	aa := []int{1, 2, 3, 4, 5}
	fmt.Printf("this is type of %T\n", aa)
	fmt.Println(aa)

	ab := []int{1, 2, 3, 4, 5}

	a := append(aa, ab...)
	fmt.Println(a)

	a = append(a, 10, 10, 23, 34)
	fmt.Println(a)

	sort.Ints(a)
	fmt.Println(a)

	abc := make([]int, 3)

	abc[0] = 32
	abc[1] = 43
	abc[2] = 54
	fmt.Println(abc)
	fmt.Printf("this is type of %T\n", abc)

	ad := append(abc, a...)
	fmt.Println(ad)

	fmt.Println(sort.IntsAreSorted(ad))

}
