package main

import (
	"fmt"
	"go-unisys/utils"
)

func main() {

	var name, email, phone string = "Vinod", "vinod@vinod.co", "9731424784"
	fmt.Printf("Name = %s\nEmail =%s\nPhone =%s\n", name, email, phone)

	var addr string = `ISRO Layout
Bangalore,
India.`

	fmt.Println(addr)

	sq := utils.Square(12)
	fmt.Println("square of 12 is ", sq)

	utils.Fibo()
	utils.HelloAgain()
}
