package day4

type Interval struct {
	Lower, Upper int
}

func (i Interval) Contains(other Interval) bool {
	return other.Lower >= i.Lower && other.Upper <= i.Upper
}

func (i Interval) Overlaps(other Interval) bool {
	return !(other.Lower > i.Upper || other.Upper < i.Lower) // not outside
}

func AnyContains(left, right Interval) bool {
	return left.Contains(right) || right.Contains(left)
}

func PairsContainCount(intervals [][2]Interval) (sum int) {
	for _, pairs := range intervals {
		if AnyContains(pairs[0], pairs[1]) {
			sum++
		}
	}
	return
}

func PairsOverlapCount(intervals [][2]Interval) (sum int) {

	for _, pairs := range intervals {
		if pairs[0].Overlaps(pairs[1]) {
			sum++
		}
	}
	return

}
