package basic

import "fmt"

func runSlice() {
	// slice
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Jun",
		"Jul",
		"Agu",
		"Sep",
		"Okt",
		"Nov",
		"Des",
	}

	var slice1 = months[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[2] = "Mar-new"
	fmt.Println(slice1)

	slice1[0] = "Mar-new2"
	fmt.Println(months)

	var append1 = append(slice1, "Mei-new")
	fmt.Println(append1)
	fmt.Println(slice1)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Malik"
	newSlice[1] = "Azizah"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
