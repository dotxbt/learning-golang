package basic

import "fmt"

func Looping() {
	// for iterator
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while
	x := 0
	for x < 10 {
		fmt.Printf("%d ", x)
		x++
	}

	// for each
	names := []string{"Saya", "tidak", "akan", "mengulangi", "kembali"}
	fmt.Println()
	// value
	for _, val := range names {
		fmt.Printf("%s ", val)
	}

	maps := map[int]string{
		1: "Jona",
		2: "Joni",
		3: "Jono",
		4: "Jone",
		5: "Jonu",
	}

	// key and value
	for k, v := range maps {
		fmt.Printf("\n %v : %v", k, v)
	}

	// value only
	for _, v := range maps {
		fmt.Printf("\n%v", v)
	}
	fmt.Printf("\n\n\n")
}
