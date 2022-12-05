package day3

import "unicode"

func FindPrioritiesOfCommonTypes(leftCompartment, rightCompartment []rune) int {
	commonTypes := FindCommonTypes(leftCompartment, rightCompartment)

	sum := 0
	for _, commonType := range commonTypes {
		sum += GetPriority(commonType)
	}
	return sum
}

func GetPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - int('A') + 27
	}
	return int(r) - int('a') + 1
}

func FindCommonTypes(leftCompartment, rightCompartment []rune) (commonRunes []rune) {
	runeOccurences := map[rune]struct{}{}

	for _, r := range leftCompartment {
		runeOccurences[r] = struct{}{}
	}

	for _, r := range rightCompartment {
		if _, ok := runeOccurences[r]; ok {
			commonRunes = append(commonRunes, r)
			delete(runeOccurences, r)
		}
	}
	return commonRunes
}
