package day6_test

import (
	"aoc2022/day6"
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed testdata/in
var in string

func TestQueue(t *testing.T) {
	queue := &day6.MappedQueue{}
	queue.Enqueue('a')
	queue.Enqueue('b')
	queue.Enqueue('c')
	dequeued := queue.Dequeue()
	require.Equalf(t, dequeued, 'a', "a vs %c", dequeued)
	dequeued = queue.Dequeue()
	require.Equalf(t, dequeued, 'b', "b vs %c", dequeued)
	dequeued = queue.Dequeue()
	require.Equalf(t, dequeued, 'c', "c vs %c", dequeued)
}

func TestFindIndexOf4ConsecutiveDistinctCharacters(t *testing.T) {
	position := day6.FindIndexOfNConsecutiveDistinctCharacters(in, 4)
	fmt.Println("for 4:", position)

	v2, err := day6.FindIndexOfNConsecDisctincCharsSuboptimal(in, 4)
	require.NoError(t, err)
	require.Equal(t, position, v2)

	///////////////

	position = day6.FindIndexOfNConsecutiveDistinctCharacters(in, 14)
	fmt.Println("for 14:", position)

	v2, err = day6.FindIndexOfNConsecDisctincCharsSuboptimal(in, 14)
	require.NoError(t, err)
	require.Equal(t, position, v2)
}

func BenchmarkFindIndexOf4ConsecutiveDistinctCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.FindIndexOfNConsecutiveDistinctCharacters(in, 14)
	}
}

func BenchmarkFindIndexOf4ConsecutiveDistinctCharactersV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.FindIndexOfNConsecDisctincCharsSuboptimal(in, 14)
	}
}
