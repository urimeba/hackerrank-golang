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
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.

Example #1: [0, 0, 1, 0, 0, 0, 0, 1, 0, 0]
Expected output: 6
Explanation: From position 0, jumps to 1
			From position 1, jumps to 3
			From position 3, jumps to 5
			From position 5, jumps to 6
			From position 6, jumps to 8
			From position 8, jumps to 9

Example #2: [0, 0, 1, 0, 0, 1, 0]
Expected output: 4
Explanation: From position 0, jumps to 1
			From position 1, jumps to 3
			From position 3, jumps to 4
			From position 4, jumps to 6

 Example #3: [0, 0, 0, 1, 0, 0]
 Expected output: 3
 Explanation: From position 0, jumps to 2
			From position 2, jumps to 4
			From position 4, jumps to 5

*/

func jumpingOnClouds(clouds []int32) int32 {
	if len(clouds) == 0 {
		return 0
	}

	jumps := 0
	for x := 0; x < len(clouds); x++ {
		currentPositionPlusTwo := x + 2
		// Check if the position exists in the array, by checking the lenght of it. Then, we check if the position, if exists, is zero. If so, we can jump 2 positions
		if len(clouds) > currentPositionPlusTwo && clouds[currentPositionPlusTwo] == 0 {
			jumps++
			x = x + 1
			continue
		}

		// Check if the position exists in the array, by checking the lenght of it. Then, we check if the position, if exists, is zero. If so, we can jump only 1 position
		currentPositionPlusOne := x + 1
		if len(clouds) > currentPositionPlusOne && clouds[currentPositionPlusOne] == 0 {
			jumps++
			continue
		}
	}

	return int32(jumps)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
