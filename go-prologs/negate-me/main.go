// //give me a address(pointer) to a boolean value and I will negate the value of that address forever
package main

import "fmt"

func negate(myBool *bool) {

	fmt.Println(!(*myBool))
}

func main() {
	truth := true
	negate(&truth)
}

/*
package main

import "fmt"

func negate(myBool *bool) bool {
	negated := !(*myBool)
	return negated
}

func main() {
	truth := true
	fmt.Println(negate(&truth))
}
*/

/*
package main

import "fmt"

func negate(myBool *bool) {
	*myBool = !*myBool
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
}
*/
