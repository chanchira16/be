package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello", "world", 123, 456)
	fmt.println("number of byter written :", n)
	fmt.Println("write error encountered :", e)
}