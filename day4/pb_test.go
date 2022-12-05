package day4_test

import (
	"aoc2022/day4"
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/pairContains
var testData string

func TestPairsContainCount(t *testing.T) {
	intervalPairs := GetIntervals(t)

	count := day4.PairsContainCount(intervalPairs)
	fmt.Println(count)
}

func TestPairsOverlapCount(t *testing.T) {
	intervalPairs := GetIntervals(t)

	count := day4.PairsOverlapCount(intervalPairs)
	fmt.Println(count)
}

func GetIntervals(t *testing.T) [][2]day4.Interval {
	pairs := strings.Split(testData, "\n")

	intervalPairs := make([][2]day4.Interval, len(pairs))
	for i, pair := range pairs {
		splitPairs := strings.Split(pair, ",")
		leftInterval, err := parsePair(splitPairs[0])
		require.NoError(t, err)
		rightInterval, err := parsePair(splitPairs[1])
		require.NoError(t, err)
		intervalPairs[i] = [2]day4.Interval{leftInterval, rightInterval}
	}
	return intervalPairs
}

func parsePair(in string) (day4.Interval, error) {
	leftAndRight := strings.Split(in, "-")
	if len(leftAndRight) != 2 {
		return day4.Interval{}, fmt.Errorf("interval does not have only lower and upper bounds for %v", in)
	}
	left, err := strconv.Atoi(leftAndRight[0])
	if err != nil {
		return day4.Interval{}, fmt.Errorf("could not parse lower for %v: %w", in, err)
	}
	right, err := strconv.Atoi(leftAndRight[1])
	if err != nil {
		return day4.Interval{}, fmt.Errorf("could not parse upper for %v: %w", in, err)
	}
	return day4.Interval{
		Lower: left,
		Upper: right,
	}, nil
}
