package main

import "fmt"

func main() {
	// Just like with arrays, if you know in advance what values a slice will start with, you can initialize the slice with those values using a slice literal:
	notes := []string{"do", "re", "me", "fa", "sol", "la", "si"}
	//or like this:
	primes := []int{
		2, // a multi-line slice literal
		3,
		5,
	}
	fmt.Println(notes)
	fmt.Println(primes)
	fmt.Println()

	//As with arrays, if you access a slice element that no value has been assigned to, you’ll get the zero value for that type back.
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])
	fmt.Println()

	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println()

	//Because a slice is just a view into the contents of an array, if you change the underlying array, those changes will also be visible within the slice:
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1) //[a X b c d e]
	fmt.Println(slice1) //[a X c]
	fmt.Println()

	//Assigning a new value to a slice element will change the corresponding element in the underlying array:
	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	slice2[1] = "X"
	fmt.Println(array2) //[f g h X j]
	fmt.Println(slice2) //[h X j]
	fmt.Println()

	//If multiple slices point to the same underlying array, a change to the array’s elements will be visible in all the slices:
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3] //[a b c d]
	slice4 := array3[2:5] //[c d e]
	array3[2] = "X"
	fmt.Println(array3) //[a b X d e]
	fmt.Println(slice3) //[a b X]
	fmt.Println(slice4) //[X d e]
}
