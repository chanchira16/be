package main

import "fmt"

func main() {
	var name string
	var age int
	n, e := fmt.scanin(&name, &age)
	fmt.Println("name ", name)
	fmt.Println("age ", age)