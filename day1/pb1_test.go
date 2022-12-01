package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Mancarici = []int

//go:embed testdata/input1
var testInput string

func TestElfWithMostFood(t *testing.T) {
	testData, err := CreateTestData(testInput)
	if err != nil {
		t.Fatal(err.Error())
	}

	result, err := ElfWithMostFood(testData)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(result)
}

func TestTop3ElvesWithMostFood(t *testing.T) {
	testData, err := CreateTestData(testInput)
	if err != nil {
		t.Fatal(err.Error())
	}

	food := TopNElvesWithMostFood(testData, 3)
	if food == 0 {
		t.Fatal("top 3 elves have 0 food")
	}
	fmt.Printf("top 3 calories: %v \n", food)
}

func CreateTestData(textInput string) ([][]int, error) {
	elfNo := 0
	arrOfStrs := strings.Split(textInput, "\n")

	mancariciByElfNo := []Mancarici{Mancarici{}}
	for _, str := range arrOfStrs {
		if str == "" {
			elfNo++
			mancariciByElfNo = append(mancariciByElfNo, nil)
			continue
		}
		food, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		mancariciByElfNo[elfNo] = append(mancariciByElfNo[elfNo], food)
	}

	return mancariciByElfNo, nil
}
