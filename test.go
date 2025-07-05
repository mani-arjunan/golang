package main

import "fmt"

// multiple return example
func yearsUntil(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age

	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age

	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

// struct example
type Person struct {
	name string
	age  int
}

func printPerson(p Person) {
	fmt.Printf("Person Name is '%v' and age is %v", p.name, p.age)
}

// multiple structs example

type Address struct {
	pincode int
	city    string
	street  string
	country string
}

type Person2 struct {
	name    string
	age     int
	address Address
}

func printPerson2(p Person2) {
	fmt.Printf("Person Name is %v and age is %v and address is %v", p.name, p.age, p.address)
}

// structs with method

type Counter struct {
	name  string
	count int
}

func (count Counter) getCounter() int {
	return count.count
}

func main() {
	// const age = "28"
	// const name = "Manikandan"
	// total := name + " " + age

	// fmt.Printf("%v", total)

	// var1, _, _ := yearsUntil(28)
	// yearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking := yearsUntil(28)
	//
	// fmt.Printf("%v ", var1)
	//
	// printPerson(Person{
	//   name: "Manikandan",
	//   age: 29,
	// })

	printPerson2(Person2{
		name: "Hema",
		age:  27,
		address: Address{
			street:  "Hema Street",
			city:    "Chennai",
			country: "India",
			pincode: 600001,
		},
	})

	c1 := Counter{
		name:  "Manikandan",
		count: 0,
	}

	c1.count++
	c1.count++

  fmt.Printf("count for c1 is %v", c1.getCounter())
	c1.count++
	c1.count++
  fmt.Printf("count for c1 is %v", c1.getCounter())

  fmt.Printf("Mani")
  fmt.Printf("Continue")
}
