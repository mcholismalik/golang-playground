package basic

import (
	"errors"
	"fmt"
	"strconv"
)

func makeScore(value int) (string, error) {
	if value > 100 {
		return "", errors.New("Value must less than or equal 100")
	}

	return "Your score is " + strconv.Itoa(value), nil
}

func runError() {
	score, err := makeScore(101)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(score)
}
