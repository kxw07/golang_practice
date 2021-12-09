package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	edward := person{firstName: "Edward", lastName: "Derrick"}
	fmt.Println(alex, edward)

	var kai person
	kai.lastName = "Wu"
	kai.firstName = "Kai"
	fmt.Println(kai)
	fmt.Printf("%+v", kai)
}
