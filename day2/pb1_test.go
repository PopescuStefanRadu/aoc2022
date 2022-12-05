package day2_test

import (
	"aoc2022/day2"
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed testdata/rockpaperscissors
var testData string

var gesturesMapping = map[string]day2.Gesture{
	"A": day2.GestureRock,
	"B": day2.GesturePaper,
	"C": day2.GestureScissors,
	"X": day2.GestureRock,
	"Y": day2.GesturePaper,
	"Z": day2.GestureScissors,
}

var inputMappings = map[string]any{
	"A": day2.GestureRock,
	"B": day2.GesturePaper,
	"C": day2.GestureScissors,
	"X": day2.OutcomeLoss,
	"Y": day2.OutcomeEqual,
	"Z": day2.OutcomeWin,
}

func TestEvaluateStrategy(t *testing.T) {
	rawRounds := strings.Split(testData, "\n")
	rounds := make([][2]day2.Gesture, len(rawRounds))

	for i, round := range rawRounds {
		gestures := strings.Split(round, " ")
		rounds[i][0] = gesturesMapping[gestures[0]]
		rounds[i][1] = gesturesMapping[gestures[1]]
	}

	result := day2.EvaluateStrategy(rounds)

	fmt.Println(result)
}

func TestEvaluatePb2Strategy(t *testing.T) {
	rawRounds := strings.Split(testData, "\n")

	gestures := make([]day2.Gesture, len(rawRounds))
	outcomes := make([]day2.Outcome, len(rawRounds))

	for i, round := range rawRounds {
		inputs := strings.Split(round, " ")
		gestures[i] = inputMappings[inputs[0]].(day2.Gesture)
		outcomes[i] = inputMappings[inputs[1]].(day2.Outcome)
	}

	result := day2.EvaluatePb2Strategy(gestures, outcomes)

	fmt.Println(result)
}
