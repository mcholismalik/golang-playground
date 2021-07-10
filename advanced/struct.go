package advanced

import (
	"fmt"
	"reflect"
)

func main() {
	runStruct()
}

func runStruct() {
	type S struct {
		Id   int    `species:"gopher" color:"blue"`
		Name string `species:"gopher" color:"blue"`
	}

	type SC struct {
		S
		Id   int    `species:"gopher" color:"red"`
		Name string `species:"gopher" color:"red"`
	}

	s := S{
		Id:   1,
		Name: "yey",
	}

	sc := SC{
		S: S{
			Id:   1,
			Name: "Parent",
		},
		Id:   1,
		Name: "Child",
	}

	st := reflect.TypeOf(s)
	sct := reflect.TypeOf(sc)
	fmt.Println("st", st)

	sv := reflect.ValueOf(&s)
	fmt.Printf("%+v\n", sv)

	scv := reflect.ValueOf(&sc)
	fmt.Printf("%+v\n", scv)
	fmt.Printf("%+v\n", sc.Id)

	num := st.NumField()
	fmt.Println(num)

	field := st.Field(0)
	fmt.Println(field.Type)
	fmt.Println(field.Tag.Get("color"))

	field2 := sct.Field(1)
	fmt.Println("here", field2.Type)
	fmt.Println(field2.Tag.Get("color"))

	value := sv.Elem().Field(0)
	fmt.Println(value)

	// reverse engineering of struct
}
