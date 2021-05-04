//! Variable di Go-Lang hanya bisa menyimpan tipe data yang sama dan bersifat unique serta harus digunakan setelah dideklarasikan

package main

import "fmt"

func main() {
	/* Variable */
	var name string
	
	name = "Rezky"
	fmt.Println(name)
	
	name = "Rezky Rizaldi" // --> Mengisi ulang variable
	fmt.Println(name)
	
	var fullName = "Muhamad Rezky Rizaldi"
	fmt.Println(fullName)
	
	// var age = 19
	var age int
	age = 19
	fmt.Println(age)

	country := "Indonesia" // --> Shorthand untuk keyword var
	fmt.Println(country)
	
	/* Multiple Variable */
	var (
		firstName = "Rezky"
		lastName = "Rizaldi"
	)
	
	fmt.Println(firstName)
	fmt.Println(lastName)
}
