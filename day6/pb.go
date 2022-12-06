package day6

import (
	"fmt"
)

// your subroutine needs to identify the first position where
// the four most recently received characters were all different.

// functia trebuie sa intoarca pozitia pentru care ultimele 4 caractere citite sunt diferite intre ele

// in: abcacccc
// a -> 1
//
// a -> 1, b -> 1
// a -> 1, b -> 1, c-> 1
// a -> 2, b -> 1, c-> 1
// a -> 1, b -> 1, c-> 2
// a -> 1, c-> 3
// a -> 0, c-> 4
// a -> 0, c-> 1, d, e, f

func FindIndexOfNConsecutiveDistinctCharacters(in string, n int) int {
	characters := []rune(in) // pentru ca vrem sa citim caracterele vizibile UTF

	// pentru ca vrem sa urmarim ultimele 4 caractere citite si sa putem sa vedem daca sunt toate diferite
	queue := &MappedQueue{}
	for i := 0; i < n; i++ {
		queue.Enqueue(characters[i])
	}
	if queue.AreAllDifferent() {
		return n
	}

	for i := n; i < len(characters); i++ { // ca sa parcurgem `characters`
		queue.Dequeue()
		queue.Enqueue(characters[i])
		if queue.AreAllDifferent() {
			return i + 1
		}
	}

	return 0
}

func FindIndexOfNConsecDisctincCharsSuboptimal(in string, n int) (int, error) {
	characters := []rune(in)

	validateAllDifferent := func(runes []rune) bool {
		occurencesByRune := map[rune]struct{}{}

		for _, r := range runes {
			// val, exists := occurencesByRune[r]
			// val := occurencesByRune[r]
			// _, exists := occurencesByRune[r]

			_, exists := occurencesByRune[r]
			if exists {
				return false
			}

			occurencesByRune[r] = struct{}{}
		}
		return true
	}

	// 0, 1, 2, 3, 4
	lastNChars := make([]rune, n)
	for i := n - 1; i < len(characters); i++ {

		for j, k := i, n-1; j > i-n; j, k = j-1, k-1 {
			lastNChars[k] = characters[j]
		}

		if validateAllDifferent(lastNChars) {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("did not find %v distinct chars", n)
}

type MappedQueue struct {
	countsByRune map[rune]int
	head, tail   *Node // queue
}

func (q *MappedQueue) Enqueue(in rune) {
	if q.countsByRune == nil {
		q.countsByRune = make(map[rune]int, 4)
	}
	q.countsByRune[in] = q.countsByRune[in] + 1
	if q.head == nil {
		q.head = &Node{
			Val: in,
		}
		q.tail = q.head
		return
	}

	prevHead := q.head

	newNode := &Node{
		Val:  in,
		Next: nil,
		Prev: prevHead,
	}
	prevHead.Next = newNode
	q.head = newNode
}

func (q *MappedQueue) Dequeue() rune {
	if q.tail == nil {
		return 0
	}

	oldValue := q.tail.Val

	if q.countsByRune == nil {
		return 0
	} else {
		newCountForKey := q.countsByRune[oldValue] - 1
		if newCountForKey == 0 {
			delete(q.countsByRune, oldValue)
		} else {
			q.countsByRune[oldValue] = newCountForKey
		}
	}

	newTail := q.tail.Next
	if newTail == nil {
		q.head = nil
		q.tail = nil
		return oldValue
	}

	newTail.Prev = nil
	q.tail.Next = nil
	q.tail = newTail
	return oldValue
}

func (q MappedQueue) AreAllDifferent() bool {
	for _, v := range q.countsByRune {
		if v != 1 {
			return false
		}
	}
	return true
}

type Node struct {
	Val  rune
	Next *Node
	Prev *Node
}
