package main

import "fmt"

func coffeeOrder(customerName string, coffeeType string, price int) (string, int) {
	massage := fmt.Sprintf("Order for %s: %s coffee costs %d taka ☕", customerName, coffeeType, price)
	return massage, price
}

func main(){
	myCoffee, price := coffeeOrder("Jon", "Black", 160)
	fmt.Println(myCoffee, price)
}
