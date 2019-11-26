package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("Hello World", "H"))
	fmt.Println(strings.Count("Hello World", ""))
}