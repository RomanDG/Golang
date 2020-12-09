package math

func Fib(num int) int {
	if num <= 1 { return num }
	return Fib(num - 1) + Fib(num - 2)
}
