package basic

import "fmt"

type Address struct {
	Country, Province string
}

func (address *Address) GetProvince() string {
	return "Provinisi : " + address.Province
}

func forceIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func runPointer() {
	address1 := Address{"Indonesia", "Jakarta"}
	address2 := &address1
	fmt.Println(address1)

	address2.Country = "Belanda"
	fmt.Println(address1)

	forceIndonesia(address2)
	fmt.Println(address2)

	var address3 *Address = new(Address)
	address3.Country = "Jepang"
	address3.Province = "Tokyo"
	forceIndonesia(address3)
	fmt.Println(address3)

	fmt.Println(address3.GetProvince())
}
