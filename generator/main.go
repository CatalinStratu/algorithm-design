package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	customFileName := "listOFNumbers.txt"
	nrOfNumbers := 1000
	f, err := os.Create(customFileName)
	//defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < nrOfNumbers; i++ {
		fmt.Fprintf(f, "%d\n", rand.Intn(nrOfNumbers))
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
