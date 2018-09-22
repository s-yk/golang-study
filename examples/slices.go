package examples

import (
	"fmt"
)

func Slice() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println("length:", len(s))

	s = append(s, "four")
	s = append(s, "five", "six")
	fmt.Println("emp2:", s)

	cop := make([]string, len(s))
	copy(cop, s)
	fmt.Println("cpy:", cop)

	fmt.Println("2-5", s[2:5])
	fmt.Println("-5", s[:5])
	fmt.Println("2-", s[2:])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
