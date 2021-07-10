package basic

import "fmt"

func runLoop() {
	// for
	counter := 0
	for counter < 10 {
		fmt.Println("for ->", counter)
		counter++
	}

	// for - statment
	for i := 0; i < 10; i++ {
		fmt.Println("for-statment ->", i)
	}

	days := []string{"Senin", "Selasa", "Rabu"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for - range
	for i, v := range days {
		fmt.Println(i, v)
	}

	box := make(map[string]string)
	box["title"] = "Magic of thinking big"
	box["color"] = "Red"
	box["owner"] = "Sasha"
	for i, v := range box {
		fmt.Println(i, v)
	}

	// break & continue
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range numbers {
		if v == 4 {
			continue
		}

		if v == 6 {
			break
		}
		fmt.Println("number ->", v)
	}
}
