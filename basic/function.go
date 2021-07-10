package basic

import "fmt"

func sayHello(person string) {
	fmt.Println("Hello", person)
}

func getPerson(name, title, position string) string {
	return name + " is " + title + " " + position
}

func getFullName() (string, string) {
	return "Dirty", "Malik"
}

func getMaskedName(a, b string, filter func(string) string) (firstName, lastName string) {
	firstName = "Mr. " + filter(a)
	return
}

func sumNumber(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func filterName(name string) string {
	if name == "Dirty" {
		name = "***"
	}
	return name
}

type BlackList func(string) bool

func blackList(name string) bool {
	return name == "user"
}

func checkBlocked(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You're blocked")
	} else {
		fmt.Println("You're not blocked")
	}
}

func runFunction() {
	firstName, lastName := getFullName()
	filter := filterName
	maskedFirstName, _ := getMaskedName(firstName, lastName, filter)
	person := getPerson(maskedFirstName, "Senior", "Backend")
	sayHello(person)

	total := sumNumber(5, 5, 5, 5)
	fmt.Println("Sum ->", total)

	slice := []int{8, 8, 8, 8}
	totalSlice := sumNumber(slice...)
	fmt.Println("Sum slice ->", totalSlice)

	isSenior := func(title string) bool {
		return title == "Senior"
	}
	fmt.Println("is senior ->", isSenior("Senior"))

	checkBlocked("user", blackList)
}
