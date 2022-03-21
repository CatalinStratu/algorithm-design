package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func insertionSort(arr []int) []int {
	array := len(arr)
	for i := 1; i < array; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
func main() {
	f, err := os.Open("intrare.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	counter := 0
	var numbers []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		counter++
		numbers = append(numbers, num)
	}

	now := time.Now()

	_ = insertionSort(numbers)
	fmt.Println(time.Now().Sub(now))

}
