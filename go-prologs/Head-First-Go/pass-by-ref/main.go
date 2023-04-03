package main

import "fmt"

func main() {
	var amount int = 6
	var point_at_me int = 10
	var amountPointer *int = &point_at_me
	fmt.Printf("Our local integer is %d.\n", amount)
	fmt.Print("It passed by value(copied) to a double function and printed from there: ")
	doubleByValue(amount)
	fmt.Printf("But it's still the same amount outer the function scope: %d.\n", amount)
	fmt.Printf("\nOur pointer points to integer: %d.\n", point_at_me)
	fmt.Print("It passed by reference(original value) to a double function and printed from there: ")
	doubleByPointer(amountPointer)
	fmt.Printf("And it has changed outer the function scope either, like forever: %d.\n", *amountPointer)
}
func doubleByValue(amount int) {
	amount *= 2
	fmt.Print(amount, ".\n")
}
func doubleByPointer(amountPointer *int) {
	*amountPointer *= 2
	fmt.Print(*amountPointer, ".\n")

}
