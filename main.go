package main

import (
	"errors"
	"fmt"
)

func main() {

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
