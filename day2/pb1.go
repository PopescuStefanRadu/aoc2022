package day2

type Gesture string

type Outcome string

const (
	GestureRock     Gesture = "rock"
	GesturePaper    Gesture = "paper"
	GestureScissors Gesture = "scissors"
	OutcomeWin      Outcome = "win"
	OutcomeEqual    Outcome = "equal"
	OutcomeLoss     Outcome = "loss"
)

var GestureToPoints = map[Gesture]int{
	GestureRock:     1,
	GesturePaper:    2,
	GestureScissors: 3,
}

func (g Gesture) DeducePlay(expectedOutcome Outcome) Gesture {
	switch expectedOutcome {
	case OutcomeWin:
		switch g {
		case GestureRock:
			return GesturePaper
		case GesturePaper:
			return GestureScissors
		case GestureScissors:
			return GestureRock
		}
	case OutcomeEqual:
		return g
	case OutcomeLoss:
		switch g {
		case GestureRock:
			return GestureScissors
		case GesturePaper:
			return GestureRock
		case GestureScissors:
			return GesturePaper
		}
	}
	return ""
}

func (g Gesture) Beats(other Gesture) bool {
	if g == GestureRock && other == GestureScissors {
		return true
	}
	if g == GesturePaper && other == GestureRock {
		return true
	}
	if g == GestureScissors && other == GesturePaper {
		return true
	}
	return false
}

func EvaluateStrategy(strategy [][2]Gesture) int {
	sum := 0
	for _, gestures := range strategy {
		sum += GestureToPoints[gestures[1]]
		if gestures[0] == gestures[1] {
			sum += 3
			continue
		}
		if gestures[1].Beats(gestures[0]) {
			sum += 6
		}
	}

	return sum
}

func EvaluatePb2Strategy(gestures []Gesture, outcomes []Outcome) int {
	sum := 0
	for i := 0; i < len(gestures); i++ {
		gesture := gestures[i]
		outcome := outcomes[i]

		switch outcome {
		case OutcomeWin:
			sum += 6
		case OutcomeEqual:
			sum += 3
		}

		ownGesture := gesture.DeducePlay(outcome)
		sum += GestureToPoints[ownGesture]
	}
	return sum
}
