package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var age int
	var heigth float32
	var weigth float32
	n, err := fmt.Scan(&name, &age, &ing, &weigth, &heigth)
	fmt. Println(name, age, ing, weigth, heigth)
	fmt. Println(`name of argument `, n)
	fmt. Pintln(`error `, err)
}