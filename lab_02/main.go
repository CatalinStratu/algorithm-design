package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
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
	var matrix [][]int
	for scanner.Scan() {
		nums := scanner.Text()
		numsStr := strings.Split(nums, " ")
		counter++
		numsInt, ok := sliceAtoi(numsStr)
		if ok != nil {
			fmt.Println(err)
			return
		}
		matrix = append(matrix, numsInt)
	}

	var A [][]int
	var B [][]int
	var C [5][5]int

	counterFirstMatrixLines := matrix[0][1]
	counterSecondMatrixLines := matrix[counterFirstMatrixLines+1][1]

	for k, v := range matrix {
		if k <= counterFirstMatrixLines && k != 0 {
			A = append(A, v)
		}
		if k > counterFirstMatrixLines+1 && k <= counterSecondMatrixLines+counterFirstMatrixLines {
			B = append(B, v)
		}
	}

	aInput := matrix[0]
	bInput := matrix[counterFirstMatrixLines+1]

	n := aInput[0]
	p := bInput[1]
	m := aInput[1]
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < m; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j]
			}
		}
	}
	fmt.Println(C)

}
