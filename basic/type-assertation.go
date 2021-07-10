package basic

import "fmt"

func any() interface{} {
	return "sasha"
}

func runAssert() {
	any := any()
	name := any.(string)
	fmt.Println(name)
}
