package main

import "fmt"


type contactInfo struct {
    email string
    zipCode int
}

type person struct {
  firstName string
  lastName string
  // assign contactInfo Type
  contactInfo
}

// Go automatically assumes struct in the following
// function

func main() {

// in GO every include line comma
  jim := person{
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo{
      email: "jim@jim.com",
      zipCode: 20781,
    },
  }

  // fmt.Printf("%+v", jim)

  // make a struct variable
  // alex := person{"Alex", "Anderson"}
  // // or
  // john := person{firstName: "John", lastName: "Smith"}
  // // this can assign struct to a variable with empty value or zero, or false
  // var simon person

  // // assignment to empty variable
  // simon.firstName = "Simone"
  // simon.lastName = "Washington"

  // fmt.Println(alex)
  // fmt.Println(john)
  // fmt.Println(simon)
  jim.updateName("jimmy")
  jim.print()
  // fmt.Printf("%+v", alex)
}

func (p person) updateName(newFirstName string) {
  p.firstName = newFirstName
}


func (p person) print() {
    fmt.Printf("%+v", p)
}

