//* Data Type: Number
/* Integer */
//* int8 (min: -128, max: 127)
//* int16 (min: -32768, max: 32767)
//* int32 (min: -2147483548, max: 2147483647)
//* int64 (min: -9223372036854775808, max: 9223372036854775807)

/* Unsigned Integer */
//* uint8 (min: 0, max: 255)
//* uint16 (min: 0, max: 65535)
//* uint32 (min: 0, max: 4294967295)
//* uint64 (min: 0, max: 18446744073709551615)

//* Data Type: Floating Point
//* float32 (min: 1.18x10^-38, max: 3.4x10^38)
//* float64 (min: 2.23x10^-308, max: 1.80x10^308)
//* complex64 (complex number with float32 real and imaginary parts)
//* complex128 (complex number with float64 real and imaginary parts)

/* Alias */
//* uint8 = byte
//* int32 = rune
//* Minimal int32 = int
//* Minimal uint32 = uint

package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)
}