package main

import (
	"errors"
	"fmt"
	"math"
	blah "spotify/utility"
	"spotify/utility/yoyo"
)

//  Struct
type person struct {
	name string
	age  int
}

// where the program starts
func main() {
	var x int = 5
	var y int = 10
	var sum int = x + y

	z := sum
	fmt.Println(sum)
	fmt.Println(z)

	if z > 10 {
		fmt.Println(z)
	} else if x == 5 {
		fmt.Println(z)
	} else {
		fmt.Println(sum)
	}

	var a [5]int
	a[2] = 7

	b := [5]int{3, 4, 2, 4}

	fmt.Println(a)
	fmt.Println(b)

	// Slice - so we don't have set number of arrays in declaration ^
	aa := []int{23, 3, 23, 4}

	aa = append(aa, 66)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dogag"] = 12

	fmt.Println(aa)
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")

	fmt.Println(vertices)

	// Loops - only for available

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// while loo

	pp := 0

	for pp < 5 {
		fmt.Println(pp)
		pp++
	}

	// using range

	zz := []string{"a", "g", "g"}

	for index, value := range zz {
		fmt.Println("Blah", index, value)
	}

	// doing the same thing with a map
	mm := make(map[string]string)
	mm["a"] = "alpha"
	mm["b"] = "beta"

	for key, value := range mm {
		fmt.Println("key: ", key, "value:", value)
	}

	rest := sumo(2, 4)
	fmt.Println(rest)

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// a struc

	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)

	// pointers

	xx := y
	inc(&xx)
	fmt.Println(xx)

	otherShit()
}

func otherShit() {
	fmt.Println(blah.GetName())
	fmt.Println(yoyo.Wing())
}

func inc(x *int) {
	*x++
}

func sumo(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined")
	}

	return math.Sqrt(x), nil
}
