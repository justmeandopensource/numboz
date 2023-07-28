package challenges

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/justmeandopensource/numboz/internal/common"
)

type challenge struct {
	Question     string
	YourAnswer   int
	ActualAnswer int
	Result       string
	Duration     time.Duration
}

var challengeReport []challenge

func Start(challengeType string, digits int) {

	var (
		question     string
		yourAnswer   int
		actualAnswer int
		result       string
		duration     time.Duration
	)

	switch challengeType {

	case "addition":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(digits)
		question = fmt.Sprintf("%v + %v = ", op1, op2)
		fmt.Print(question)
		startTime := time.Now()
		fmt.Scan(&yourAnswer)
		duration = time.Since(startTime).Round(time.Second)
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
		startTime := time.Now()
		fmt.Scan(&yourAnswer)
		duration = time.Since(startTime).Round(time.Second)
		actualAnswer = num1 - num2

	case "multiplication":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v x %v = ", op1, op2)
		fmt.Print(question)
		startTime := time.Now()
		fmt.Scan(&yourAnswer)
		duration = time.Since(startTime).Round(time.Second)
		actualAnswer = op1 * op2

	case "division":
		op1 := common.GenerateRandomNumber(digits)
		op2 := common.GenerateRandomNumber(1) + 1
		question = fmt.Sprintf("%v / %v = ", op1, op2)
		fmt.Print(question)
		startTime := time.Now()
		fmt.Scan(&yourAnswer)
		duration = time.Since(startTime).Round(time.Second)
		actualAnswer = op1 / op2
	}

	if yourAnswer == actualAnswer {
		result = "PASS"
	} else {
		result = "FAIL"
	}

	c := challenge{
		Question:     question,
		YourAnswer:   yourAnswer,
		ActualAnswer: actualAnswer,
		Result:       result,
		Duration:     duration,
	}
	challengeReport = append(challengeReport, c)
}

func Mixed(digits int) {
	Start(*common.PickRandomChallenge(), digits)
}

func PrintChallengeReport() {

	common.ClearTerminal()

	allPass := true

	for _, record := range challengeReport {
		if record.Result == "FAIL" {
			allPass = false
		}
	}

	table := simpletable.New()

	generateReportTableHeader(allPass, table)
	generateReportTableBody(allPass, table)

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

	totalCorrect, totalTimeTaken := getScoreStats()

	fmt.Println(common.ColorizeBlue(fmt.Sprintf("%v correct in %v", totalCorrect, totalTimeTaken)))
}

func generateReportTableHeader(allpass bool, table *simpletable.Table) {

	if allpass {
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Text: "#"},
				{Text: "RESULT"},
				{Text: "TIME TAKEN"},
				{Text: "QUESTION"},
			},
		}
	} else {
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Text: "#"},
				{Text: "RESULT"},
				{Text: "TIME TAKEN"},
				{Text: "QUESTION"},
				{Text: "YOUR ANSWER"},
				{Text: "ACTUAL ANSWER"},
			},
		}
	}
}

func generateReportTableBody(allpass bool, table *simpletable.Table) {
	if allpass {
		for i, record := range challengeReport {

			r := []*simpletable.Cell{
				{Text: strconv.Itoa(i + 1)},
				{Text: common.ColorizeGreen(record.Result)},
				{Text: record.Duration.String()},
				{Text: strings.Replace(record.Question, "=", "", -1)},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}
	} else {
		for i, record := range challengeReport {

			var (
				result       string
				yourAnswer   string
				actualAnswer string
			)

			if record.Result == "FAIL" {
				result = common.ColorizeRed(record.Result)
				yourAnswer = common.ColorizeRed(strconv.Itoa(record.YourAnswer))
				actualAnswer = common.ColorizeGreen(strconv.Itoa(record.ActualAnswer))
			} else {
				result = common.ColorizeGreen(record.Result)
			}

			r := []*simpletable.Cell{
				{Text: strconv.Itoa(i + 1)},
				{Text: result},
				{Text: record.Duration.String()},
				{Text: strings.Replace(record.Question, "=", "", -1)},
				{Text: yourAnswer},
				{Text: actualAnswer},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}
	}
}

func getScoreStats() (int, string) {

	var (
		totalCorrect   int
		totalTimeTaken time.Duration
	)

	for _, record := range challengeReport {
		if record.Result == "PASS" {
			totalCorrect += 1
		}
		totalTimeTaken += record.Duration
	}

	return totalCorrect, totalTimeTaken.String()
}
