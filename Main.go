package main

import (
	"fmt"
	"learning-golang/basic"
	"log"
)

func main() {

	fmt.Println("------------------ SHOW ALL -----------------")
	fmt.Println(":::: BASIC ::::")
	// print message
	basic.Hello()

	// error return
	msg, err := basic.HelloWithError("Sabituddin Bigbang")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	// random message
	msg2, err := basic.HelloRandomMessage("Sabituddin Bigbang")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg2)

	// multiple return
	names := []string{"Sabit", "Udin", "Bigbang"}
	msg3, err := basic.HelloPeoples(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg3)

	// data structure
	basic.BasicType()
	basic.ACar()

}
