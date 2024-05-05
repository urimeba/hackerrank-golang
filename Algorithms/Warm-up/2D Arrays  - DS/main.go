package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.

Input #1:
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Expected output:
10
Explanation:
The 16 ascendent sorted hourglass sums are:
[0 2 3 3 4 4 6 6 7 7 8 8 9 10 14 19]
The highest hourglass sum is 19

Input #2:
0 9 2 -4 -4 0
0 0 0 -2 0 0
0 0 -1 -2 -4 0
Expected output:
13
Explanation:
The 16 ascendent sorted hourglass sums are:
[-15 -14 -8 -6 -6 -5 0 2 2 4 4 7 9 10 12 13]
The highest hourglass sum is 13

Input #3:
-9 -9 -9 1 1 1
0 -9 0 4 3 2
-9 -9 -9 1 2 3
0 0 8 6 6 0
0 0 0 -2 0 0
0 0 1 2 4 0
Expected output:
28
Explanation:
The 16 ascendent sorted hourglass sums are:
[-63 -34 -27 -11 -10 -9 -2 0 9 10 12 17 18 23 25 28]
The highest hourglass sum is 28
*/

func hourglassSum(arr [][]int32) int32 {
	lenArray := len(arr)
	if lenArray < 3 { // <- If the array is less than 3, we should return 0 as you can't form an hourglass
		return 0
	}

	var allOfTheSum []int // <- This will store all of the sums of the hourglasses

	// We should iterate over the array,
	for x := 0; x < lenArray; x++ {
		if x-1 < 0 || x+1 >= lenArray { // <- We should skip the arrays that don't have an UPPER array or a LOWER array, as the array is always 6x6
			continue
		}

		// We should iterate over the array of the array
		for y := 0; y < len(arr[x]); y++ {
			if y-1 < 0 || y+1 >= len(arr[x]) { // <- We should skip the arrays that don't have a LEFT or RIGHT value, as the array is always 6x6
				continue
			}

			someSum := arr[x-1][y-1] + // <- This is the sum of the hourglass
				arr[x-1][y] +
				arr[x-1][y+1] +
				arr[x][y] +
				arr[x+1][y-1] +
				arr[x+1][y] +
				arr[x+1][y+1]

			allOfTheSum = append(allOfTheSum, int(someSum)) // <- We append the sum to the array
		}
	}

	sort.Slice(allOfTheSum, func(x, y int) bool {
		return allOfTheSum[x] < allOfTheSum[y] // <- We sort the array ascendent
	})

	return int32(allOfTheSum[len(allOfTheSum)-1]) // <- We return the last value of the array (the highest value)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
