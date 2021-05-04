package main

import "fmt"

func main() {
  /* Convert Integer */
	var val32 int32 = 128
	var val64 int64 = int64(val32)
	var val8 int8 = int8(val32)
	
	fmt.Println(val32)
	fmt.Println(val64)
	fmt.Println(val8)

  /* Convert String */
	var name = "Rezky"
	// var n uint8 = name[0]
	// var n byte = name[0]
	var n = name[0] // --> 82 (byte for R)
	var nString string = string(n) // --> R (string for 82 byte)
	fmt.Println(n)
	fmt.Println(nString)
}