// вызываем ./main 10 20 30

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		num, _ := strconv.Atoi(v)
		result := fib(num)
		fmt.Println(result)
	}
}

func fib(num int) int {
	if num <= 1 { return num }
	return fib(num - 1) + fib(num - 2)
}
