package main

import "fmt"

type person struct {
	name    string
	age     int
	address string

	hobbys hobbys
}

type hobbys struct {
	hobby string
}

func main() {

	person1 := person{"name: Davi", 24, "Home 3", hobbys{"Hobby: Padel"}}
	person2 := person{"name: Kayo", 24, "Address: Streets: Antonio Neves Pedro", hobbys{"Hobby: Programming"}}
	person3 := person{"name: Victória", 26, "Address: Henrique Sales", hobbys{"Hobby: Jiu-Jitsu"}}
	person4 := person{"name: Yoda", 900, "Address: Space", hobbys{"Hobby: Telecinese"}}

	fmt.Println("Characteristics of people")

	fmt.Println("------------------------------------------------------------------------")

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
}
