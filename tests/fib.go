package main
import "fmt"

func fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return fib(n-2) + fib(n-1)
	}
}

func main() {
	fmt.Println(fib(10))
}
