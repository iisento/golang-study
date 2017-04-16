package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fib0, fib1, fib2 int = 0, 1, 0
	var fib int
	return func() int {
		fib = fib2

		// 次のフィボナッチ数を計算する
		fib0 = fib1
		fib1 = fib2
		fib2 = fib0 + fib1

		// 今回のフィボナッチ数は前回の呼び出しで計算されてい
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
