package day6

// abcabcd
// 4
// a,b,c,a,b,c,d

func FindIndexOfNConsecutiveDistinctCharactersOptimal(in string, n int) int {
	// trebuie sa tina intr-un map ultimele N caractere
	// in map numaram pentru fiecare caracter de cate ori l-am intalnit in ultimele N caractere
	//

	runes := []rune(in)
	occurrenceMap := make(map[rune]int, n)

	for i := 0; i < n; i++ {
		r := runes[i]
		occurrenceMap[r] = occurrenceMap[r] + 1
	}

	if len(occurrenceMap) == n {
		return n
	}

	// edzez
	// d -> 1
	// e -> 2
	// z -> 1

	// -e

	// d -> 1
	// e -> 1
	// z -> 1

	// +z

	// d -> 1
	// e -> 1
	// z -> 2

	// must have 4 different
	// abcdde
	// 4 -> 0
	// 5 -> 1
	// 6 -> 2
	// abcdefg
	// 0123456
	for i := n; i < len(runes); i++ {
		elemToRemove := runes[i-n]                                // elemToRemove ia valoarea runei de la indexul i-n, adica o sa fie caracterul de acu N caractere
		occurrencesForElemToRemove := occurrenceMap[elemToRemove] // occurrencesForElemToRemove tine numarul de intalniri al elemToRemove
		if occurrencesForElemToRemove == 1 {                      // daca am intalnit elemToRemove o singura data
			delete(occurrenceMap, elemToRemove)
		} else { // daca am intalnit elemToRemove de mai multe ori
			occurrenceMap[elemToRemove] = occurrencesForElemToRemove - 1
		}

		elemToAdd := runes[i] //  'z'
		occurrenceMap[elemToAdd] = occurrenceMap[elemToAdd] + 1

		if len(occurrenceMap) == n {
			return i + 1
		}
	}

	return 0
}

