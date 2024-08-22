package main

import (
	"fmt"
	"interview-go/pkg"
)

func main() {

	defer fmt.Println("should be runned first!")
	defer fmt.Println("should be runned second!")
	defer fmt.Println("should be runned third!")
	defer fmt.Println("should be runned fourth!")
	defer fmt.Println("should be runned fifth!")
	defer fmt.Println("should be runned sixth!")

	pkg.Map()
	fmt.Println()

	pkg.Array()
	fmt.Println()

	pkg.EmptyMap()
	fmt.Println()

	pkg.InitializedWithNoElementsMap()
	fmt.Println()

	pkg.Slice()
	fmt.Println()

	var number = 1
	pkg.PointerFunction(&number)
	fmt.Println(number)

	var animal pkg.AnimalSound

	animal = pkg.Dog{}
	fmt.Println("Dog", animal.Sound())

	animal = pkg.Cat{}
	fmt.Println("Cat", animal.Sound())
	fmt.Println()

	pkg.InterfaceAsGenerics("hello")
	pkg.InterfaceAsGenerics(1)
	pkg.InterfaceAsGenerics(2.1)
	fmt.Println()

	pkg.InterfaceToPassNumberOfArguments(1, 2.1, "23")
	fmt.Println()

	pkg.TypeAssertions()
	fmt.Println()

	pkg.Errors()
	fmt.Println()

	if pkg.CheckName("Jason") == nil {
		fmt.Println(true)
	} else {
		fmt.Println("youre failed")
	}
	fmt.Println()

	pkg.HandlingErrorsWithErrorf(-12)
	fmt.Println()

	number1 := 15
	number2 := 0

	result, err := pkg.Divide(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
	fmt.Println()

	pkg.ToHandlePanic(4, 2)
	pkg.ToHandlePanic(4, 1)
	pkg.ToHandlePanic(4, 0)

	pkg.Division(4, 2)
	pkg.Division(8, 0)
	pkg.Division(2, 8)

}
