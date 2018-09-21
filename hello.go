package main

import (
	"fmt"
	"os"
	"time"
)

var funcMap = map[string]interface{}{
	"hello":     hello,
	"values":    values,
	"variables": variables,
	"for":       forexample,
	"if":        ifexample,
	"switch":    switchexample,
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("no argument")
		return
	}

	arg := os.Args[1]
	targetFunc, isExist := funcMap[arg]
	if !isExist {
		fmt.Println("wrong argument")
		return
	}
	targetFunc.(func())()
}

func switchexample() {
	var t = time.Now()
	var sec = t.Second()
	fmt.Println("now :", sec)

	switch sec {
	case 2:
		fmt.Println("2")
	case 5, 6:
		fmt.Println("5 or 6")
	default:
		fmt.Println("default")
	}

	switch {
	case sec%5 == 0:
		fmt.Println("sec divisible by 5")
	case sec%2 == 0:
		fmt.Println("sec divisible by 2")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func ifexample() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func forexample() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func variables() {
	var a = "string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func hello() {
	var message = "Hello, world.\n"
	fmt.Printf(message)
}
