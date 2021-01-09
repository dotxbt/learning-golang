package basic

import "fmt"

func BasicType() {
	var1 := "this is string"
	var2 := 'C'
	var3 := 6969692789473957395
	var4 := true
	var5 := 69699379842497924872487.6289479279
	var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"} 
	var7 := []string{"foo", "bar", "baz"}
	var8 := complex(9, 15) 

	fmt.Printf("var1 = %T\n", var1) 
    fmt.Printf("var2 = %T\n", var2) 
    fmt.Printf("var3 = %T\n", var3) 
    fmt.Printf("var4 = %T\n", var4) 
    fmt.Printf("var5 = %T\n", var5) 
    fmt.Printf("var6 = %T\n", var6) 
    fmt.Printf("var7 = %T\n", var7) 
    fmt.Printf("var8 = %T\n", var8) 
}