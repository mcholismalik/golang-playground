package advanced

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func runIo() {
	// read the whole file at once
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// replacement
	str := string(b)
	fmt.Println(str)
	str = strings.ReplaceAll(str, "T_InterfaceName", "Sample")
	str = strings.ReplaceAll(str, "T_StructName", "sample")

	b = []byte(str)

	// write the whole body at once
	err = ioutil.WriteFile("output.go", b, 0644)
	if err != nil {
		panic(err)
	}

}
