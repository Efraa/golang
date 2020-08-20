package main

import "fmt"

// import (
// 	"fmt"
// )

// La manera de exportar funciones o variables es declarandolos con la primera letra en Mayus

// func GetListOfMembers() {

// }

// Global scope

// Vars
var numberVar int
var stringVar string
var booleanVar bool

func main() {
	// Asignar tipo en la declaracion de la variable
	countMembers := 231003
	var memberName string = "John Doe"

	fmt.Println("Hello Goo!")
	fmt.Println(countMembers)
	fmt.Println(numberVar)
	fmt.Println(stringVar)
	fmt.Println(booleanVar)
	fmt.Println(memberName)
}
