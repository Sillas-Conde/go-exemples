package list

import (
	"fmt"
	"strconv"
	// "os"
)


func List_text(list []float64) {
	//var numbers []int

    // Append values to the slice
    //numbers = append(numbers, 10, 20, 30, 40)
	numbers := list

	fmt.Println("Os números da lista são:")

	// for i := 0; i < len(numbers); i++ {
    //     fmt.Println(numbers[i])
    // }

	for i := range numbers {
        fmt.Println(numbers[i])
    }

	// forEach value in numbers {

	// }
}

func List_interface(list []interface{}) {

	for _,item := range list {
		fmt.Println(item)
	}

}

func List_each(fruits []string) {

    // Iterating over the slice using range
    for index, fruit := range fruits {

		fmt.Println("Posição: " + strconv.Itoa(index)  + ", Fruta: " + fruit)
        //fmt.Print(fruit + "")
		//fmt.Println(sla)
    }
}