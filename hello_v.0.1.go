package main

import "fmt"

import "rsc.io/quote" 
import "strconv"

func Main_2(name string, age int) {
    fmt.Println(quote.Go())
	//age := 30

	age_str := strconv.Itoa(age)

	text := "His/Her name is " + name + " and his/her age is: " + age_str

	fmt.Println(text)
}

