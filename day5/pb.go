package day5

type CraneMovement struct {
	Count int
	From  int
	To    int
}

func ApplyCraneMovementsStacked(state [][]string, movements []CraneMovement) [][]string {
	for _, movement := range movements {
		source := state[movement.From]
		mostBottomPickedUpCrane := len(source) - movement.Count

		state[movement.From] = source[:mostBottomPickedUpCrane]
		state[movement.To] = append(state[movement.To], source[mostBottomPickedUpCrane:]...)
	}

	return state

}

func ApplyCraneMovementsOneByOne(state [][]string, movements []CraneMovement) [][]string {
	for _, movement := range movements {
		source := state[movement.From]
		target := state[movement.To]
		for i := 0; i < movement.Count; i++ {
			target = append(target, source[len(source)-1])
			source = source[:len(source)-1]
		}
		state[movement.From] = source
		state[movement.To] = target
	}

	return state
}

func getTopCranes(state [][]string) (res []string) {
	for _, boxes := range state {
		if len(boxes) != 0 {
			res = append(res, boxes[len(boxes)-1])
		}
	}
	return res
}

func ApplyCraneMovementsAndGetTopCranes(state [][]string, movements []CraneMovement, applyMovement func(state [][]string, movement []CraneMovement) [][]string) string {
	newState := applyMovement(state, movements)
	cranes := getTopCranes(newState)
	result := ""
	for _, crane := range cranes {
		result += crane
	}
	return result
}
