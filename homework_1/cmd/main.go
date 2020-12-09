package main

import (
	"HardcodeDev/Golang/homework_1/pkg/math"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	result := math.Fib(num)
	fmt.Println(result)
}
