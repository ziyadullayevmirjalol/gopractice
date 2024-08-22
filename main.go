package main

import (
	"errors"
	"fmt"
)

func main() {

	defer fmt.Println("should be runned first!")
	defer fmt.Println("should be runned second!")
	defer fmt.Println("should be runned third!")
	defer fmt.Println("should be runned fourth!")
	defer fmt.Println("should be runned fifth!")
	defer fmt.Println("should be runned sixth!")

	Map()
	fmt.Println()

	Array()
	fmt.Println()

	EmptyMap()
	fmt.Println()

	InitializedWithNoElementsMap()
	fmt.Println()

	Slice()
	fmt.Println()

	var number = 1
	PointerFunction(&number)
	fmt.Println(number)

	var animal AnimalSound

	animal = Dog{}
	fmt.Println("Dog", animal.Sound())

	animal = Cat{}
	fmt.Println("Cat", animal.Sound())
	fmt.Println()

	InterfaceAsGenerics("hello")
	InterfaceAsGenerics(1)
	InterfaceAsGenerics(2.1)
	fmt.Println()

	InterfaceToPassNumberOfArguments(1, 2.1, "23")
	fmt.Println()

	TypeAssertions()
	fmt.Println()

	Errors()
	fmt.Println()

	if CheckName("Jason") == nil {
		fmt.Println(true)
	} else {
		fmt.Println("youre failed")
	}
	fmt.Println()

	HandlingErrorsWithErrorf(-12)
	fmt.Println()

	number1 := 15
	number2 := 0

	result, err := divide(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
	fmt.Println()

	ToHandlePanic(4, 2)
	ToHandlePanic(4, 1)
	ToHandlePanic(4, 0)

	division(4, 2)
	division(8, 0)
	division(2, 8)

}

// ////////////////////////////////////////////////////////////////////////////
func Map() {
	var mymap = make(map[string]interface{})

	mymap["a"] = 1
	mymap["b"] = 2
	mymap["c"] = 3

	for i, v := range mymap {
		fmt.Printf("%s is %v\n", i, v)
	}
}

func Array() {
	var myarr = [5]int{}

	myarr[0] = 5
	myarr[1] = 1
	myarr[2] = 2
	myarr[3] = 3
	myarr[4] = 2

	for i, v := range myarr {
		fmt.Printf("%d is %v\n", i, v)
	}
}

func Slice() {
	var slc = []int{1, 2, 2, 34, 5, 6, 6, 7, 8, 9, 0}

	var copyslc = make([]int, len(slc))

	for i := 0; i < len(slc); i++ {
		if slc[i] > 9 {
			slc = append(slc[:i], slc[i+1:]...)
		}
	}

	n := copy(copyslc, slc)
	fmt.Println(slc, copyslc, n)
}

func EmptyMap() {
	var emptmap map[int]int

	if len(emptmap) == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func InitializedWithNoElementsMap() {
	m := make(map[int]int)
	fmt.Println(len(m) == 0)
}

func PointerFunction(num *int) {
	*num = 100
}

type AnimalSound interface {
	Sound() string
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Meow"
}

type Dog struct{}

func (c Dog) Sound() string {
	return "Bark"
}

func InterfaceAsGenerics(i interface{}) {
	fmt.Println(i)
}

func InterfaceToPassNumberOfArguments(i ...interface{}) {
	fmt.Println(i)
}

func TypeAssertions() {
	var i interface{}

	i = "1"
	intAssign, boolval := i.(int)
	// using boolval to avoid panic
	fmt.Println(intAssign, boolval)
}

func Errors() {

	msg := "jaso"

	myerr := errors.New("You hafta be Jason!")
	if msg != "jason" {
		fmt.Println(myerr)
	}
}

func CheckName(name string) error {
	if name != "Jason" {
		newErr := errors.New("You hafta be Jason!")
		return newErr
	}
	return nil
}

func HandlingErrorsWithErrorf(age int) {

	error := fmt.Errorf("%v is a negative value\n", age)
	if age < 0 {
		fmt.Print(error)
	} else {
		fmt.Println("Age is %v", age)
	}
}

type error interface {
	Error() string
}

type DevisionByZeroMessage struct {
	message string
}

func (z DevisionByZeroMessage) Error() string {
	return "Number cannot be divided by zerooo, Call of Doooo, Advanced Poo Poo"
}

func divide(n1 int, n2 int) (int, error) {
	if n2 == 0 {
		return 0, &DevisionByZeroMessage{}
	} else {
		return n1 / n2, nil
	}
}

func division(num1, num2 int) {

	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func HandlePanic() {
	h := recover()
	if h != nil {
		fmt.Println("Recovered message:", h)
	}
}

func ToHandlePanic(n1 int, n2 int) {
	defer HandlePanic()

	if n2 == 0 {
		panic("num2 is 0")
	} else {
		res := n1 / n2
		fmt.Println("Res:", res)
	}

}
