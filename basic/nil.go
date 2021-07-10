package basic

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func runNil() {
	customer := NewMap("Andi")
	if customer == nil {
		fmt.Println("Data empty")
	} else {
		fmt.Println(customer["name"])
	}
}
