package main

import (
	"fmt"
)

func main() {
	fmt.Println(H1("hello"))
}

func H1(el string) string {
	return fmt.Sprintf("<h1>%s</h1>", el)
}