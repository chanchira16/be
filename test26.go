package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var age int
	var heigth float32
	var weigth float32
	n, err := fmt.Scan(&name, &age, &ing, &heigth, &weigth)
	fmt.println(name, age, ing, heigth, weigth)
	fmt.println(`name of argumment `, n)
	fmt.println(`error `, err)
