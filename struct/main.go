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
	// it can also be declared as contactInfo with out field name,
	// it will be treated as contactInfo contactInfo
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

	//store pointer to thor in var
	thorPointerAddress := &thor
	thorPointerAddress.updateFirstName("Brother")
	thor.print()

	//BUT WAIT even this works
	tony.updateFirstName("Billionaire")
	tony.print()
}

//print person with reciever function
func (p person) print() {
	fmt.Printf("%+v", p)
}

//below function does not update actual first, but updates on firstName
//property on object 'p'. Main reason for this is go is pass by value
// and hence maintains a pointer to each variable passed to function
func (p person) updateFirstName_NotWorking(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
