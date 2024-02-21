package main

import (
	"fmt"
	"math"
	"unsafe"
)

const loginTokenZ = "abc123"

func main() {
	var i int
	fmt.Println("Default value:", i)

	var str string
	fmt.Println("Default value:", str)

	// Implicit variable declaration
	var website = "example.com"
	fmt.Println(website)
	fmt.Printf("Type: %T\n", website)

	// No-var variable declaration
	// Only allowed for local variables
	jwtToken := "300000"
	fmt.Println(jwtToken)
	fmt.Printf("Type: %T\n", jwtToken)

	fmt.Println("loginTokenZ:", loginTokenZ)

	// Constant variable declaration
	const loginToken = "abc123"
	fmt.Println(loginToken)
	fmt.Printf("Type: %T\n", loginToken)

	// Data type sizes
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", int8(0), unsafe.Sizeof(int8(0)), math.MinInt8, math.MaxInt8)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", int16(0), unsafe.Sizeof(int16(0)), math.MinInt16, math.MaxInt16)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", int32(0), unsafe.Sizeof(int32(0)), math.MinInt32, math.MaxInt32)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", int64(0), unsafe.Sizeof(int64(0)), math.MinInt64, math.MaxInt64)

	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", uint8(0), unsafe.Sizeof(uint8(0)), 0, math.MaxUint8)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", uint16(0), unsafe.Sizeof(uint16(0)), 0, math.MaxUint16)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", uint32(0), unsafe.Sizeof(uint32(0)), 0, math.MaxUint32)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %d, Max: %d\n", uint64(0), unsafe.Sizeof(uint64(0)), 0, uint64(math.MaxUint64))

	fmt.Printf("Type: %T, Size: %d bytes, Min: %.0f, Max: %.0f\n", float32(0), unsafe.Sizeof(float32(0)), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("Type: %T, Size: %d bytes, Min: %.0f, Max: %.0f\n", float64(0), unsafe.Sizeof(float64(0)), -math.MaxFloat64, math.MaxFloat64)

	fmt.Printf("Type: %T, Size: %d bytes, Min: %t, Max: %t\n", bool(false), unsafe.Sizeof(bool(false)), false, true)

	fmt.Printf("Type: %T, Size: %d bytes\n", complex64(0), unsafe.Sizeof(complex64(0)))
	fmt.Printf("Type: %T, Size: %d bytes\n", complex128(0), unsafe.Sizeof(complex128(0)))

	fmt.Printf("Type: %T, Size: %d bytes\n", byte(0), unsafe.Sizeof(byte(0))) // alias for uint8
	fmt.Printf("Type: %T, Size: %d bytes\n", rune(0), unsafe.Sizeof(rune(0))) // alias for int32
}
