package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("janjira", "be"))
	fmt.Println(strings.HasSuffix("janjira", "Be"))
}
