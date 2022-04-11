package main

import "fmt"

//define a new custom type `person` whicn `extends` struct
type person struct {
	firstName string
	lastName  string
}

func main() {

	//create a new person
	//go relies on order of fields to assign values to struct fields
	tony := person{"Tony", "Stark"} // NOT RECOMMENDED

	//another syntax to define struct
	thor := person{firstName: "Thor", lastName: "Odinson"}

	fmt.Println(tony, thor)
	//OP: {Tony Stark} {Thor Odinson}

	fmt.Printf("%+v", thor)
	//OP: {firstName:Thor lastName:Odinson}
}
