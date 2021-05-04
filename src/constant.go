//? Constant merupakan sebuah variable yang value-nya tidak bisa diubah setelah dideklarasikan, namun tidak wajib digunakan

package main

import "fmt"

func main() {
	/* Constant */
	const name string = "Rezky"
	const age = 18
	
	fmt.Println(name)
	fmt.Println(age)
	
	/* Multiple Constant */
	const (
		name2 string = "Rizaldi"
		age2 = 19
	)
	
	fmt.Println(name2)
	fmt.Println(age2)
}