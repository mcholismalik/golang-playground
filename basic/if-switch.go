package basic

import "fmt"

func runIfSwitch() {
	// if
	title := "Senior"

	if title == "Senior" {
		fmt.Println("You're senior")
	} else {
		fmt.Println("You're not senior")
	}
	// if - short statement
	pass := "ab3c"
	if length := len(pass); length < 4 {
		fmt.Println("Too small")
	} else {
		fmt.Println("Normal")
	}

	// switch
	title = "Junior"
	switch title {
	case "Senior":
		fmt.Println("You're senior")
	case "Junior":
		fmt.Println("You're junior")
	default:
		fmt.Println("You're undefined")

	}

	// switch - short statement
	switch length := len(pass); length < 4 {
	case true:
		fmt.Println("Too small")
	default:
		fmt.Println("Normal")
	}

}
