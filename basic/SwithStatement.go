package basic

import "fmt"

func SwitchSt(id int) {
	switch id {
	case 1: 
		fmt.Println("Select One") 
		return
	case 2: 
		fmt.Println("Select Two") 
		return
	case 3: 
		fmt.Println("Select Three") 
		return
	default :
		fmt.Println("Select Others") 
		return

		
	}
}