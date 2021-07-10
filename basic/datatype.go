package basic

import "fmt"

func datatype() {
	// variable
	var (
		firstName = "Muhammad"
		lastName  = "Cholis"
	)
	fmt.Println("Hello ", firstName, lastName)

	// number
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(100000)
	fmt.Println("Number ", nilai32, nilai64)

	// string
	var oString = string(lastName[2])
	fmt.Println("String ", oString)

	// bool
	type Marriage bool
	var isMarriage Marriage = false
	fmt.Println("Marriage ", isMarriage)
}
