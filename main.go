package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	common "./common"

	challenge "./challenge"
)

func main() {

	fmt.Println("Please, enter with information of interaction of execution")

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)

	common.CheckError(err)

	maxDigit := int32(nTemp)

	result := challenge.PerformChallenge(int(maxDigit))

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
