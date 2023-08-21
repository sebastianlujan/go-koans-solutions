package go_koans

import (
	"fmt"
	"math"
)

func aboutBasics() {
	assert(true == true)  // what is truth?
	assert(true != false) // in it there is nothing false

	var i int = 1
	var j int = 42
	println(j)

	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := 1 //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	var odd int = 1
	ten := 10

	//constants are not infered types
	const FIVE = 5
	TWENTY_FIVE := math.Pow(FIVE, 2)

	fmt.Println(5^3, TWENTY_FIVE)

	assert(5%2 == odd)
	assert(5*2 == ten)

	// XOR , 101 ^ 100 = 001
	// XOR , 101 ^ 011 = 110 -> 6
	assert(5^4 == odd)

	var x int
	const ZERO = 0

	assert(x == ZERO) // zero values are valued in Go

	const ZERO_DECIMAL = 0.0

	var f float32
	assert(f == ZERO_DECIMAL) // for types of all types

	var s string
	var void string = ""
	assert(s == void) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}

	VOID := ""

	assert(c.x == ZERO)         // and types within composite types
	assert(c.f == ZERO_DECIMAL) // which match the other types
	assert(c.s == VOID)         // in a typical way
}
