package challenge

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	common "../common"
)

// PerformChallenge -- Perform challenge
func PerformChallenge(maxNumber int) (resultNumber []int) {

	const (
		maxPossibilities  = 4
		resultExpected    = 21
		timeDurantionFunc = "30s"
	)

	var (
		counter      int
		digits       string
		durationFunc time.Duration
		elapsed      time.Duration
	)

	durationFunc, err := time.ParseDuration(timeDurantionFunc)

	common.CheckError(err)

	validation(maxNumber, maxPossibilities)

	start := time.Now()

	for {

		counter++

		digits += fmt.Sprint(counter)

		if len(digits) == maxPossibilities {

			sum := checkSumOfDigits(digits, maxNumber)

			if sum == resultExpected {

				number, err := strconv.Atoi(digits)

				common.CheckError(err)

				resultNumber = append(resultNumber, number)

			} else {

				digits = ""

			}

		} else if len(digits) > maxPossibilities {

			digits = ""

		}

		if len(resultNumber) >= maxNumber {

			break
		}

		elapsed = time.Since(start)

		if elapsed.Minutes() >= durationFunc.Minutes() {
			fmt.Printf("Timout of %s exceeded\n", elapsed.String())
			break
		}

	}

	sort.Ints(resultNumber)

	fmt.Print(`
			----- [Show quantity execution to generate result] -----
			`, counter,
		`
			---------------------------`,
		`
			----- [Time of execution] -----
			`, elapsed.String(),
		`
			---------------------------
			Result: `)

	return

}

func checkSumOfDigits(digits string, maxNumber int) (total int) {

	for _, digit := range digits {

		number, err := strconv.Atoi(string(digit))

		common.CheckError(err)

		if number > maxNumber {

			total = 0

			break
		}

		total += number

	}

	return

}

func validation(maxNumber, maxPossibilities int) {

	if maxNumber < maxPossibilities {

		err := fmt.Errorf(
			"the execution is less than max of possibilities. Number max of execution: %d | Max of possibilities %d",
			maxNumber,
			maxPossibilities)

		common.CheckError(err)
	}
}
