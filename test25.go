package main

import "fmt"

func main() {
	fmt.Print("import : ")
	var text string
	fmt.Scan(&text)
	fmt.Println(' read "' , text, '" from standard input )
}