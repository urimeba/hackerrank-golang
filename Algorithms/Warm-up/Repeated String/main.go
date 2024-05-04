package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const checkValue = "a"

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// If the string only has 1 value, and is "a", we should return 1 (n)
	if len(s) == 1 && string(s[0]) == checkValue {
		return n
	}

	timesToRepeat := n / int64(len(s))   // <- This is the number of times we should multiply the repeated character in the string
	residuoDivision := n % int64(len(s)) // <- This is the remaining characters that we should check

	timesRepeatedNormalString := strings.Count(s, checkValue)                   // <- This is the number of times the character is repeated in the string
	timesRepeatedJoinedString := timesRepeatedNormalString * int(timesToRepeat) // <- This is the number of times the character is repeated in the string multiplied by the number of times we should repeat the string

	if residuoDivision != 0 {
		timesRepeatedJoinedString += strings.Count(s[:residuoDivision], checkValue) // <- This is the number of times the character is repeated in the remaining characters
	}

	return int64(timesRepeatedJoinedString)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
