package basic

import "fmt"

type Car struct {
	carName string
	carEngine string
	carWheel string
}

func ACar() {
	car := Car {
		carName: "Lamborgini",
		carEngine: "4500cc",
		carWheel: "4WD",
	}

	fmt.Println(car)
}