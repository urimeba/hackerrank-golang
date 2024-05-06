package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d


Input #1:
5 4
1 2 3 4 5
Expected output:
5 1 2 3 4
Explanation:
Rotating 4 times to the left, the array will be: [5, 1, 2, 3, 4]

Input #2:
20 10
41 73 89 7 10 1 59 58 84 77 77 97 58 1 86 58 26 10 86 51
Expected output:
77 97 58 1 86 58 26 10 86 51 41 73 89 7 10 1 59 58 84 77
Explanation:
Rotating 10 times to the left, the array will be [77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77]

Input #3:
15 13
33 47 70 37 8 53 13 93 71 72 51 100 60 87 97
Expected output:
87 97 33 47 70 37 8 53 13 93 71 72 51 100 60
Explanation:
Rotating 13 times to the left, the array will be [87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60]

*/

func rotLeft(a []int32, d int32) []int32 {
	if d == 0 { // <- If no swap is needed,	just return the array
		return a
	}

	return append(a[d:], a[:d]...) // <- Append the array from "d" to the end and then append the array from 0 to "d"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
