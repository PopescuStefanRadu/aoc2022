package day5_test

import (
	"aoc2022/day5"
	_ "embed"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

//go:embed testdata/cranes
var testData string

func TestApplyCraneMovements(t *testing.T) {
	allLines := strings.Split(testData, "\n")

	stateSeparatorPosition, err := getStateSeparatorPosition(allLines)
	require.NoError(t, err)

	stateLines := allLines[:stateSeparatorPosition]

	state, err := day5.ReadStateInput(stateLines)

	craneMovementLines, err := day5.ReadCraneMovementLines(allLines[stateSeparatorPosition+1:])
	require.NoError(t, err)

	result := day5.ApplyCraneMovementsAndGetTopCranes(state, craneMovementLines, day5.ApplyCraneMovementsOneByOne)
	fmt.Println("one by one", result)

	state, err = day5.ReadStateInput(stateLines)
	require.NoError(t, err)

	result2 := day5.ApplyCraneMovementsAndGetTopCranes(state, craneMovementLines, day5.ApplyCraneMovementsStacked)
	fmt.Println("stacked", result2)
}

func getStateSeparatorPosition(lines []string) (int, error) {
	for pos, line := range lines {
		if line == "" {
			return pos, nil
		}
	}

	return 0, errors.New("could not get position of separator between state and movement")
}
