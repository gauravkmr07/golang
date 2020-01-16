package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("******Go routines example*******")
	myFunc("Gaurav Kumar Sharma")

	go myFunc("---GKS---")

	go func(str string) {
		fmt.Println(":: go anonymous " + str)
	}("Go Gaurav!!")

	defer fmt.Println("done!!")

	time.Sleep(time.Second)

}

func myFunc(str string) {
	for _, val := range [4]int{1, 2, 3, 4} {
		fmt.Println(str + " : " + strconv.Itoa(val))
	}
}
