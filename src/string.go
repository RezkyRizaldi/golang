//! string di Go-Lang menganut Double Quotes ("")

/* String functions */ 
// len("string") (Mennghitung jumlah karakter di dalam String)
// "string"[number] (Mengambil karakter di dalam String pada index tertentu)

package main

import "fmt"

func main() {
	fmt.Println("Rezky Rizaldi")
	fmt.Println(len("Rezky Rizaldi"))
	fmt.Println("Rezky Rizaldi"[0]) // --> 82 (byte for R)
}