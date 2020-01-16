package main

import (
	"fmt"
	"log"
)

func testFunc(arg int) (int, error) {
	if arg > 80000 {
		return -1, argError{arg: arg, prob: "Test Func Error"}
	}

	return arg + 1, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main() {
	fmt.Println("-- Errors Examples --")

	_, error := testFunc(98977)

	if error != nil {
		log.Print(error)
	}

}
