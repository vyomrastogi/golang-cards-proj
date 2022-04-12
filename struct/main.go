package main

import "fmt"

//embedded struct of contactInfo
type contactInfo struct {
	email   string
	zipCode int
}

//define a new custom type `person` whicn `extends` struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//create a new person
	//go relies on order of fields to assign values to struct fields
	tony := person{"Tony", "Stark", contactInfo{}} // NOT RECOMMENDED

	//another syntax to define struct
	thor := person{firstName: "Thor", lastName: "Odinson", contact: contactInfo{
		email: "thor@avengers.com", zipCode: 00000}}

	fmt.Println(tony)
	//OP: {Tony Stark} {Thor Odinson}

	thor.print()
	//OP: {firstName:Thor lastName:Odinson contact:{email:thor@avengers.com zipCode:0}}
}

//print person with reciever function
func (p person) print() {
	fmt.Printf("%+v", p)
}
