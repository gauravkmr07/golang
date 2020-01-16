package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	fmt.Println("--- Strings Functions ---")
	var str string = "Hello Dada!!"

	p("Contains :: ", s.Contains(str, "ll"))
	p("Explode :: ", s.Split(str, " "))
	p("To Lower :: ", s.ToLower(str))
	p("Replace :: ", s.Replace(str, "o", "O", 1))
}
