package challenges

import (
	"fmt"
	"strings"

	"github.com/justmeandopensource/numboz/internal/common"
)

type challenge struct {
	Question       string
	ExpectedResult int
	ActualResult   int
	Result         string
}

var challengeReport []challenge

func Start(challengeType string, digits int) {

	var question string
	var expectedResult, actualResult int
	result := "FAIL"

	switch challengeType {

	case "addition":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(digits)
		question = fmt.Sprintf("%v + %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&expectedResult)
		actualResult = op1 + op2

	case "subtraction":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(digits)
		var num1, num2 int
		if op1 > op2 {
			num1, num2 = op1, op2
		} else {
			num1, num2 = op2, op1
		}
		question = fmt.Sprintf("%v - %v = ", num1, num2)
		fmt.Print(question)
		fmt.Scan(&expectedResult)
		actualResult = num1 - num2

	case "multiplication":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v x %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&expectedResult)
		actualResult = op1 * op2

	case "division":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v / %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&expectedResult)
		actualResult = op1 / op2
	}

	if expectedResult == actualResult {
		result = "PASS"
	}

	c := challenge{
		Question:       question,
		ExpectedResult: expectedResult,
		ActualResult:   actualResult,
		Result:         result,
	}
	challengeReport = append(challengeReport, c)
}

func Mixed(digits int) {
	Start(*common.PickRandomChallenge(), digits)
}

func PrintChallengeReport() {
	common.ClearTerminal()
	for i, record := range challengeReport {
		fmt.Printf("Q%v\t%v\t%v\t%v\t%v\n",
			i+1,
			strings.Replace(record.Question, "=", "", -1),
			record.ExpectedResult,
			record.ActualResult,
			record.Result,
		)
	}
}
