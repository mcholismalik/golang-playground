package basic

import "fmt"

type HasName interface {
	GetName() string
}

type Customer struct {
	Name, Country string
	Age           int
}

func (customer Customer) Show() {
	fmt.Println("User", customer.Name, "from", customer.Country, "is", customer.Age, "yo")
}

func (customer Customer) GetName() string {
	return customer.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func Any() interface{} {
	// return "string"
	return 1
}

func runStructInterface() {
	fmt.Println("Start program")

	// way 1
	var customer Customer
	customer.Name = "Malik"
	customer.Country = "Indonesia"
	customer.Age = 20
	customer.Show()

	// way 2
	customer2 := Customer{
		Name:    "Aisyah",
		Country: "Belanda",
		Age:     22,
	}
	customer2.Show()

	// way 3
	customer3 := Customer{"Yogi", "Palestina", 19}
	customer3.Show()

	// get name
	SayHello(customer3)

	// interface any
	fmt.Println(Any())
}
