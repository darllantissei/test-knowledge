package challenge

import (
	challenge "../challenge"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

const (
	paramMaxNumber   = 6
	paramSumExpected = 21
)

var paramListNumbersExpecteds = []int{
	3666, 4566, 4656, 4665, 5466, 6465,
}

func Test_ExpectedValues(t *testing.T) {
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name             string
		args             args
		wantResultNumber []int
	}{
		{
			name: "Expected values",
			args: args{
				maxNumber: paramMaxNumber,
			},
			wantResultNumber: paramListNumbersExpecteds,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResultNumber := challenge.PerformChallenge(tt.args.maxNumber); !reflect.DeepEqual(gotResultNumber, tt.wantResultNumber) {
				t.Errorf("ExpectedValues() = %v, want %v", gotResultNumber, tt.wantResultNumber)
			}
		})
	}
}

func Test_Has4Digits(t *testing.T) {
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name       string
		args       args
		has4Digits int
	}{
		{
			name: "Has 4 digits",
			args: args{
				maxNumber: paramMaxNumber,
			},
			has4Digits: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultNumber := challenge.PerformChallenge(tt.args.maxNumber)

			for _, number := range gotResultNumber {

				if len(fmt.Sprint(number)) != tt.has4Digits {

					t.Errorf("Has4Digits() = %v, want %v", len(fmt.Sprint(number)), tt.has4Digits)
				}

			}
		})
	}
}

func Test_EachIsEqualOrLessParamMaxNumber(t *testing.T) {
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name          string
		args          args
		isEqualOrLess int
	}{
		{
			name: "Digits of numbers are equal or less than paramMaxNumber",
			args: args{
				maxNumber: paramMaxNumber,
			},
			isEqualOrLess: paramMaxNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotResultNumber := challenge.PerformChallenge(tt.args.maxNumber)

			for _, number := range gotResultNumber {

				digits := fmt.Sprint(number)

				for _, digit := range digits {

					value, err := strconv.Atoi(string(digit))

					if err != nil {
						t.Fatalf("EachIsEqualOrLessParamMaxNumber() error. Details: %v", err)
					}

					if value > tt.isEqualOrLess {
						t.Errorf("EachIsEqualOrLessParamMaxNumber() = %v, want %v", value, tt.isEqualOrLess)
					}
				}

			}

		})
	}
}

func Test_SumExpected(t *testing.T) {
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name        string
		args        args
		sumExpected int
	}{
		{
			name: "Sum expected is",
			args: args{
				maxNumber: paramMaxNumber,
			},
			sumExpected: paramSumExpected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotResultNumber := challenge.PerformChallenge(tt.args.maxNumber)

			for _, number := range gotResultNumber {

				digits := fmt.Sprint(number)

				sum := 0
				for _, digit := range digits {

					value, err := strconv.Atoi(string(digit))

					if err != nil {
						t.Fatalf("SumExpected() error. Details: %v", err)
					}

					sum += value

				}
				if sum != tt.sumExpected {
					t.Errorf("SumExpected() = %v, want %v", sum, tt.sumExpected)
				}

			}

		})
	}
}
