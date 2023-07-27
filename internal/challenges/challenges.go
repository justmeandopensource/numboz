package challenges

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/justmeandopensource/numboz/internal/common"
)

type challenge struct {
	Question     string
	YourAnswer   int
	ActualAnswer int
	Result       string
}

var challengeReport []challenge

func Start(challengeType string, digits int) {

	var question string
	var yourAnswer, actualAnswer int
	result := "FAIL"

	switch challengeType {

	case "addition":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(digits)
		question = fmt.Sprintf("%v + %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&yourAnswer)
		actualAnswer = op1 + op2

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
		fmt.Scan(&yourAnswer)
		actualAnswer = num1 - num2

	case "multiplication":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v x %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&yourAnswer)
		actualAnswer = op1 * op2

	case "division":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v / %v = ", op1, op2)
		fmt.Print(question)
		fmt.Scan(&yourAnswer)
		actualAnswer = op1 / op2
	}

	if yourAnswer == actualAnswer {
		result = "PASS"
	}

	c := challenge{
		Question:     question,
		YourAnswer:   yourAnswer,
		ActualAnswer: actualAnswer,
		Result:       result,
	}
	challengeReport = append(challengeReport, c)
}

func Mixed(digits int) {
	Start(*common.PickRandomChallenge(), digits)
}

func PrintChallengeReport() {

	common.ClearTerminal()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "#"},
			{Text: "QUESTION"},
			{Text: "ACTUAL ANSWER"},
			{Text: "YOUR ANSWER"},
			{Text: "RESULT"},
		},
	}

	for i, record := range challengeReport {

		r := []*simpletable.Cell{
			{Text: strconv.Itoa(i + 1)},
			{Text: strings.Replace(record.Question, "=", "", -1)},
			{Text: strconv.Itoa(record.YourAnswer)},
			{Text: strconv.Itoa(record.ActualAnswer)},
			{Text: record.Result},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
