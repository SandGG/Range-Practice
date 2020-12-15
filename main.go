package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	age    int
}

var people = []person{
	{name: "Rosa", gender: "female", age: 16},
	{name: "Laura", gender: "female", age: 26},
	{name: "Juan", gender: "male", age: 19},
	{name: "Ricardo", gender: "male", age: 21},
	{name: "Ana", gender: "female", age: 25},
}

func main() {
	i := search()
	if i != -1 {
		fmt.Println("Name :", people[i].name)
		fmt.Println("Gender :", people[i].gender)
		fmt.Println("Age :", people[i].age)
	} else {
		fmt.Println("Data not found")
	}
}

func search() int {
	var search string
	fmt.Println("Enter name to search")
	fmt.Scan(&search)
	for i, v := range people {
		if search == v.name {
			return i
		}
	}
	return -1
}
