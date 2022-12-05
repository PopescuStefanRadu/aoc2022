package day3_test

import (
	"aoc2022/day3"
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

//go:embed testdata/rucksacks
var inputData string

func TestFindPrioritiesOfCommonTypes(t *testing.T) {
	rucksacks := strings.Split(inputData, "\n")

	sum := 0
	for _, rucksack := range rucksacks {
		runes := []rune(rucksack)
		sum += day3.FindPrioritiesOfCommonTypes(runes[0:len(runes)/2], runes[len(runes)/2:])
	}
	fmt.Println(sum)
}

func TestFindPrioritiesForCommonRunes(t *testing.T) {
	rucksacks := strings.Split(inputData, "\n")
	groupCount := len(rucksacks) / 3

	groups := make([][3][]rune, groupCount)

	for i := 0; i < groupCount; i = i + 1 {
		pos := i * 3
		groups[i] = [3][]rune{
			[]rune(rucksacks[pos]),
			[]rune(rucksacks[pos+1]),
			[]rune(rucksacks[pos+2]),
		}
	}
	sum, err := day3.FindPrioritiesForCommonRunes(groups)
	require.NoError(t, err)
	fmt.Println(sum)
}
