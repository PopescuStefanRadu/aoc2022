package day3

import (
	"fmt"
)

type AllTrueSlice struct {
	trueCount int
	slice     []bool
}

func (s AllTrueSlice) IsAllTrue() bool {
	return s.trueCount == len(s.slice)
}

func (s *AllTrueSlice) SetTrue(pos int) {
	if s.slice[pos] == false {
		s.trueCount++
	}
	s.slice[pos] = true
}

func NewAllTrueSlice(sliceSize int) *AllTrueSlice {
	return &AllTrueSlice{
		trueCount: 0,
		slice:     make([]bool, sliceSize),
	}
}

func FindCommonRune(runesSlices ...[]rune) (rune, error) {
	runeOccurences := map[rune]*AllTrueSlice{}

	for i, slice := range runesSlices {
		for _, r := range slice {
			allTrueSlice, exists := runeOccurences[r]
			if !exists {
				allTrueSlice = NewAllTrueSlice(3)
				runeOccurences[r] = allTrueSlice
			}
			allTrueSlice.SetTrue(i)
			if allTrueSlice.IsAllTrue() {
				return r, nil
			}
		}
	}

	testData := ""
	for _, slice := range runesSlices {
		testData += string(slice) + "\n"
	}

	return 0, fmt.Errorf("could not find common rune for testData: \n %v", testData)
}

func FindPrioritiesForCommonRunes(runeSlices [][3][]rune) (int, error) {
	sum := 0
	for _, slices := range runeSlices {
		commonRune, err := FindCommonRune(slices[:]...)
		if err != nil {
			return 0, err
		}
		sum += GetPriority(commonRune)
	}

	return sum, nil
}
