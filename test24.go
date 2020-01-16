package main

import "fmt"

func main() {
	n, e := fmt.Println("hello", "world", 160, 743)
	fmt.println("number of byter written :", n)
	fmt.Println("write error encountered :", e)