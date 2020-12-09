package main

import (
	"fmt"
	"HardcodeDev/Golang/homework_1/pkg/math"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	a := math.Fib(num)
	fmt.Println(a)
}
