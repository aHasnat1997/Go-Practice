package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	if n > 1 {
		for {
			fmt.Print(n, " ")
			if n == 1 {
				break
			}
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
		}
	}
}
