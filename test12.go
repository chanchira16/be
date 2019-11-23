package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(search.Index("Hello World", "world"))
	fmt.Println(search.Index("Hello World", "World"))
}
