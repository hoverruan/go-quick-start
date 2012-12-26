package main

import (
	"fmt"
)

func main() {
	// variables
	var v1 int = 1
	v2 := 2
	v3, v4 := 3, 4
	fmt.Println(v1, v2, v3, v4)

	// multi-line strings
	str1 := `Hello
		World!`
	fmt.Println(str1)

	// const and iota
	const (
		No  = iota
		Yes = iota
	)
	fmt.Printf("Yes = %d, No = %d\n", Yes, No)

	// array
	a1 := [3]int{1, 2, 3}
	fmt.Printf("Len of a1 = %d\n", len(a1))

	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Len of a2 = %d\n", len(a2))

	// slice
	s1 := []int{1, 2, 3}
	fmt.Printf("Len of s1 = %d, cap(s1) = %d\n", len(s1), cap(s1))

	s2 := make([]int, 5)
	fmt.Printf("Len of s2 = %d, cap(s2) = %d\n", len(s2), cap(s2))

	s3 := make([]int, 1, 5)
	fmt.Printf("Len of s3 = %d, cap(s3) = %d\n", len(s3), cap(s3))

	s4 := s1[1:]
	fmt.Printf("Len of s4 = %d, cap(s4) = %d\n", len(s4), cap(s4))

	// map
	m1 := map[string]int{"c": 3, "java": 9, "go": 999}
	fmt.Println("m1[go] =", m1["go"])

	// if else
	if Yes == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	if x := 10; x > 5 {
		fmt.Println("x greater than 5")
	} else if x < 5 {
		fmt.Println("x less than 5")
	}

	// goto
	i1 := 0
ComeHere:
	fmt.Println(i1)
	i1++
	if i1 < 10 {
		goto ComeHere
	}

	// for
	for i2 := 0; i2 < 10; i2++ {
		fmt.Println(i2 * 2)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum", sum)

	for k1, v1 := range m1 {
		fmt.Printf("m1[%s] = %d\n", k1, v1)
	}

	for _, v2 := range m1 {
		fmt.Println(v2)
	}

	// switch
	i3 := 10
	switch i3 {
	case 1:
		fmt.Println("i3 == 1")
	case 2, 3, 4:
		fmt.Println("i3 == 2, 3, 4")
	default:
		fmt.Println("other")
	}

	switch {
	case i3 < 3:
		fmt.Println("i3 < 3")
	case i3 < 6:
		fmt.Println("i3 < 6")
	}
}
