package main

import (
	"fmt"
	// "os"
	"GO/list"
	"GO/files"
	// "strconv"
)

func main() {

	
    fmt.Println("Hello, World!")
	Main_2("Sillas", 25)
	c := sum(25, 5)

	fmt.Printf("He is gonna turn %.0f years old in 2030", c)
	fmt.Println()
	files.Fmt_write("./exemple.txt", "Conde")

	fmt.Println("-----------------------//------------------")

	var list_ex []float64

	list_ex = append(list_ex, 1,5,7,3,2,77,3)
	list.List_text(list_ex)

	fruits := []string{"Apple", "Banana", "Cherry"}

	list.List_each(fruits)

	fmt.Println("-----------------------//------------------")
	//list2 := []interface{}{"Alice", 30, 3.14, true}

	list2 := []any{"Alice", 30, 3.14, true}

	list2 = append(list2, "Sillas", "Conde")

	list.List_interface(list2)
}


func sum(a float64, b float64) (float64) {

	//c := a + b
	return a + b
}


