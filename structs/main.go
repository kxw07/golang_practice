package main

import "fmt"

type info struct {
	age int
}

type person struct {
	firstName string
	lastName  string
	info
}

func main() {
	alexInfo := info{age: 20}
	alex := person{"Alex", "Anderson", alexInfo}
	alex.updateFirstName("Kai")
	alex.updateAge(25)
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateAge(newAge int) {
	p.info.age = newAge
}
