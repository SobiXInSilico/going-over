package main

import "fmt"

func main() {
	var amount int = 6
	var point_at_me int = 10
	var amountPointer *int = &point_at_me
	doubleByValue(amount) // will print 12
	fmt.Println(amount)   //will print 6
	//fmt.Println(&amount)
	fmt.Println(point_at_me)
	doubleByPointer(amountPointer)
}
func doubleByValue(amount int) {
	amount *= 2
	fmt.Println(amount)
	//fmt.Println(&amount)
}
func doubleByPointer(amountPointer *int) {
	*amountPointer *= 2
	fmt.Println(*amountPointer)
}
