package main

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "world", 789, 101112, "\n")
	fmt.Print("number of byter written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")

}

