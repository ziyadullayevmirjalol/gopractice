package pkg

import "fmt"

func SendData(ch chan string) {
	ch <- "Successfully send to the Channel"
	fmt.Println("Sending blocked cause of no reciever...")
}
func RecieveData(ch chan string) {
	fmt.Println(<-ch)
}
