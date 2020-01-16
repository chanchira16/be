package main

imporn "fmt"

func main() {
	n, e := fmt.Print("Hello", "world", 123, 456, "\n")
	fmt.Print("number of byter written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")

}

