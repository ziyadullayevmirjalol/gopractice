package pkg

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func PrintLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(string(letter))
		time.Sleep(500 * time.Millisecond)
	}
}
