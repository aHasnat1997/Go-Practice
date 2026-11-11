package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var missing int64
	for i := 1; i <= n; i++ {
		missing ^= int64(i)
	}

	for i := 0; i < n-1; i++ {
		var number int
		fmt.Scan(&number)
		missing ^= int64(number)
	}

	fmt.Println(missing)
}
