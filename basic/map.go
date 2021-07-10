package basic

import "fmt"

func runMap() {
	person := map[string]string{
		"name":  "Cholis",
		"title": "Senior",
	}

	fmt.Println(person)
	fmt.Println(person["title"])

	var newMap = make(map[string]string)
	newMap["position"] = "Backend"
	fmt.Println(newMap)

	delete(person, "title")
	fmt.Println(person)
}
