package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var arr []int
	var copyArr []int
	for scanner.Scan() {
		num, ok := strconv.Atoi(scanner.Text())
		if ok != nil {
			fmt.Println(err)
			return
		}

		arr = append(arr, num)
	}
	copyArr = append(copyArr, arr...)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	median := 0.0
	// get count of scores
	var totalElements = len(arr)
	// check if total number of scores is even
	if totalElements%2 == 0 {
		var sumOfMiddleElements = arr[totalElements/2] +
			arr[totalElements/2-1]
		// calculate average of middle elements
		median = float64(sumOfMiddleElements / 2)
	} else {
		// get the middle element
		median = float64(arr[len(arr)/2])
	}
	fmt.Println("Median: ", median)
	var medianPosition = int(median)
	position := 0
	for key, item := range copyArr {
		if item == medianPosition {
			position = key
			fmt.Println("Position: ", key)
		}
	}

	f, err = os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	medianText := "Median: "
	positionText := "Position: "
	var medianString string = strconv.FormatFloat(median, 'E', -1, 32)
	positionStr := strconv.Itoa(position)

	result := medianText + (medianString) + " " + positionText + positionStr
	_, err = f.WriteString(result)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}
