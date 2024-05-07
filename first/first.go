package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter no. for table : ")
	input, _ := reader.ReadString('\n')

	// fmt.Printf("tell me the type %T: ", input)

	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("adding 1 to rating : ", numRating+1)
	// }

	var i float64 = 1
	for i <= 10 {
		fmt.Println(numRating, " x ", i, " = ", numRating*i)
		i++
	}
}
