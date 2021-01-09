package main

import (
	"fmt"
	"learning-golang/basic"
	"learning-golang/network"
	"learning-golang/oop"
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

	// ============================== OOP ==============================
	// constructor
	mahasiswa, err := oop.NewStudent("62363824", "Anak Nakal", "Jl. Jalan Gak Pernah Pulang", "Teknik Informatika")
	if err != nil {
		log.Fatal(err.Error())
	}

	// call method
	mahasiswa.PrintStudent()

	// call setter
	mahasiswa.SetName("Teman Amjay")

	// call getter
	fmt.Printf("\ngetter : %v\n", mahasiswa.GetName())
	mahasiswa.PrintStudent()

	// ========================== NETWORK ==============================
	fmt.Println("\n\n:::: NETWORK ::::")
	network.Server()

}
