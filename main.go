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

/*

Complete a seguinte função para que a mesma devolva todos os
possíveis números de 4 dígitos, onde cada um seja menor ou igual
a <maxDigit>, e a soma dos dígitos de cada número gerado seja 21

Exemplo com maxDigit=6: 3666, 4566

*/

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
