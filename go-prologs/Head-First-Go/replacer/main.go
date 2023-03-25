package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "g# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(strings.Title(fixed))
}
