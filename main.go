package main

import (
	"fmt"

	"github.com/cuongvu3009/learn-go/piscine"
)

func main() {
	s := "Hello"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
