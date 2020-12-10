package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Errors in Go")

	_, err := os.Create("/Users/vinod/Desktop/tmp")
	if err != nil {
		fmt.Println("There was an error while creating file 'tmp' on Desktop", err)
	} else {
		fmt.Println("File 'tmp' successfully created!")
	}

}
