package main

import (
	"fmt"
	"myapp/a"
	nested "myapp/nested"
	"myapp/z"

	goinittutoriallibrary "github.com/david-yappeter/go-init-tutorial-library"
)

var globalVar = 10 // 1st

func init() { // 2nd
	_ = goinittutoriallibrary.Var
	_ = nested.A
	_ = a.A
	_ = z.A
	fmt.Println("")
	fmt.Println("==========================")
	fmt.Println("1st init()")
	fmt.Println("==========================")
	fmt.Println("global Var: ", globalVar)
	globalVar = 5
}

func init() { // 3rd
	fmt.Println("")
	fmt.Println("==========================")
	fmt.Println("3nd init()")
	fmt.Println("==========================")
	fmt.Println("global Var: ", globalVar)
	globalVar = 100
}

func main() { // 5th
	fmt.Println("")
	fmt.Println("==========================")
	fmt.Println("main()")
	fmt.Println("==========================")
	fmt.Println("globalVar: ", globalVar)
}

func init() { //4th
	fmt.Println("")
	fmt.Println("==========================")
	fmt.Println("4nd init()")
	fmt.Println("==========================")
	fmt.Println("global Var: ", globalVar)
	globalVar = 50
}
