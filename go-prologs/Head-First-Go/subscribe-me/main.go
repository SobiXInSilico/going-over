package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
func printInfo(s subscriber) {
	fmt.Println("\nName:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
}
func applyDiscount(s *subscriber) {
	s.rate = 0.99
}
func main() {
	subscriber1 := defaultSubscriber("Sobi Baba")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Simin Negahban")
	subscriber2.active = false
	printInfo(subscriber2)
	applyDiscount(&subscriber2)
	fmt.Printf("Discount updated: %.2f\n", subscriber2.rate)
}
