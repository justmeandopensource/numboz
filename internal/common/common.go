package common

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var challengeTypes = []string{"addition", "subtraction", "multiplication", "division"}

func ValidateParams(challengeType string, digits int) {

	challengeTypes = append(challengeTypes, "mixed")
	validChallengeType := false

	for _, item := range challengeTypes {
		if challengeType == item {
			validChallengeType = true
			break
		}
	}

	if !validChallengeType {
		fmt.Println("[ERR] not a valid challenge type")
		os.Exit(1)
	}

	if digits < 2 || digits > 4 {
		fmt.Println("[ERR] valid value for digits is 2, 3 or 4")
		os.Exit(1)
	}
}

func GenerateRandomNumber(digits int) int {
	rand.Seed(time.Now().UnixNano())
	ceil := math.Pow(10, float64(digits)) - 1
	floor := math.Pow(10, float64(digits-1))
	randNumber := int(floor) + rand.Intn(int(ceil-floor))
	return randNumber
}

func PickRandomChallenge() *string {
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(len(challengeTypes))
	randomChallenge := challengeTypes[randNumber]
	return &randomChallenge
}

func ClearTerminal() {
	var clearCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
