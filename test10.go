package main

importn (
	"fmt"
	"string"
)

func main() {
	fmt.Println(strings.HasSuffix("Hello World", "world"))
	fmt.Println(strings.HasSuffix("Hello World",  "World"))
}
