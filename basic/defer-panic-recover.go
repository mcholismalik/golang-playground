package basic

import "fmt"

func end() {
	fmt.Println("Handle logging")
	msg := recover()
	fmt.Println("Handle panic", msg)
}

func calculate(value int, err bool) {
	defer end()

	if err {
		panic("Sample panic")
	}

	fmt.Println("Start calculate ...")
	result := 10 / value
	fmt.Println("Result is ", result)
}

func main() {
	calculate(9, true)
}
