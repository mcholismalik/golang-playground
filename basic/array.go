package basic

import "fmt"

func array() {
	// array
	var names [3]string
	names[0] = "Aisyah"
	names[1] = "Indra"
	names[2] = "Alizza"
	fmt.Println("Array", names[0])

	var values = [2]int{
		21,
		13,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
