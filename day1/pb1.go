package day1

import "errors"

func TopNElvesWithMostFood(foodByElfNumber [][]int, n int) int {
	if n == 0 {
		return 0
	}
	totalFoodByElfNo := computeTotalFoodByElfNumber(foodByElfNumber)
	topNCalories := make([]int, n, n) // []int{1, 2, 3}, la stanga este cel mai mic, top3Calories[0]= cel mai mic din top 3

	for _, totalCaloriesByElf := range totalFoodByElfNo {
		posToInsert := -1
		for pos, topCalorie := range topNCalories {
			if totalCaloriesByElf > topCalorie {
				posToInsert = pos
			} else if totalCaloriesByElf < topCalorie {
				break
			}
		}

		if posToInsert == -1 {
			continue
		}

		for i := 1; i <= posToInsert; i++ { // shift values that are smaller than totalCaloriesByElf to the left
			topNCalories[i-1] = topNCalories[i]
		}

		topNCalories[posToInsert] = totalCaloriesByElf
	}

	sum := 0
	for i := 0; i < len(topNCalories); i++ {
		sum += topNCalories[i]
	}

	return sum
}

func ElfWithMostFood(foodByElfNumber [][]int) (int, error) {
	if len(foodByElfNumber) < 1 {
		return 0, errors.New("there are no elves")
	}

	// 1231241 1231231 515151 6968855
	// 11 1231 55343
	// 8686 4040

	// 12312314
	// 412413132
	// 1321314141

	// 1321314141

	totalFoodByElfNo := computeTotalFoodByElfNumber(foodByElfNumber)

	maxVal := totalFoodByElfNo[0]
	for i := 1; i < len(totalFoodByElfNo); i++ {
		if maxVal < totalFoodByElfNo[i] {
			maxVal = totalFoodByElfNo[i]
		}
	}

	return maxVal, nil
}

func computeTotalFoodByElfNumber(foodByElfNumber [][]int) []int {
	totalFoodByElfNo := make([]int, len(foodByElfNumber))

	for elfNo, foods := range foodByElfNumber { // ne uitam la un elf anume
		sumForOneElf := 0

		for _, food := range foods { // ne uitam la o mancare anume a acelui elf
			sumForOneElf += food
		}

		totalFoodByElfNo[elfNo] = sumForOneElf
	}
	return totalFoodByElfNo
}
