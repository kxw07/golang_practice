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

	array := [3]string{"apple", "banana", "melo"}
	updateArray(array)
	fmt.Println(array)

	slice := []string{"apple", "banana", "melo"}
	updateSlice(slice)
	fmt.Println(slice)
}

func updateArray(array [3]string) {
	array[0] = "Hello"
}

func updateSlice(slice []string) {
	slice[0] = "Hello"
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateAge(newAge int) {
	p.info.age = newAge
}
