package day5

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strconv"
)

type LineReaderState string

const (
	LineReaderInsideName  LineReaderState = "inside"
	LineReaderOutsideName LineReaderState = "outside"
)

type LineReader struct {
	State     LineReaderState
	CrateName string
	io.RuneReader
}

func ReadStateInput(lines []string) ([][]string, error) {
	state, presentRowNums, err := initializeStacksFromNumbers(lines[len(lines)-1])
	if err != nil {
		return nil, err
	}

	crateLineCount := len(lines) - 1
	readers := make([]*LineReader, crateLineCount)
	for i := 0; i < crateLineCount; i++ {
		// we have the readers reversed because we want to read the stacks from base to top
		readers[crateLineCount-i-1] = &LineReader{
			State:      LineReaderOutsideName,
			RuneReader: bytes.NewBuffer([]byte(lines[i])),
		}
	}
	maxWidth := 0
	for _, line := range lines {
		currLen := len([]rune(line))
		if maxWidth < currLen {
			maxWidth = currLen
		}
	}

	presentRowIndex := 0
	currentRowNum := presentRowNums[0]
	var readingNames bool
	for i := 0; i < maxWidth; i++ {
		for _, reader := range readers {
			r, _, err := reader.ReadRune()
			if errors.Is(err, io.EOF) {
				continue
			}
			if err != nil {
				return nil, err
			}
			switch r {
			case '[':
				reader.CrateName = ""
				reader.State = LineReaderInsideName
				readingNames = true
			case ' ':
				reader.CrateName = reader.CrateName + " "
			case ']':
				reader.State = LineReaderOutsideName
				state[currentRowNum] = append(state[currentRowNum], reader.CrateName)
				reader.CrateName = ""
			default:
				reader.State = LineReaderInsideName
				reader.CrateName = reader.CrateName + string(r)
			}
		}

		if !readingNames {
			// we need to continue until we start reading a crate name for any of our line readers
			continue
		}

		readingNames = false
		for _, reader := range readers {
			if reader.State == LineReaderInsideName {
				readingNames = true
				break
			}
		}

		if !readingNames {
			presentRowIndex++
			if presentRowIndex < len(presentRowNums) {
				// our stopping condition doesn't prevent out of bounds error here
				currentRowNum = presentRowNums[presentRowIndex]
			}
		}
	}
	return state, err
}

// initializeStacksFromNumbers returns the stack slice, a slice with the present stack numbers and an error
func initializeStacksFromNumbers(in string) ([][]string, []int, error) {
	digits, err := regexp.Compile("\\d+")
	if err != nil {
		return nil, nil, err
	}
	numbers := digits.FindAllString(in, -1)
	var presentRowNums []int
	maxNumber := 0
	for _, number := range numbers {
		maxNumber, err = strconv.Atoi(number)
		presentRowNums = append(presentRowNums, maxNumber)
		if err != nil {
			return nil, nil, err
		}
	}
	return make([][]string, maxNumber+1), presentRowNums, nil
}

func ReadCraneMovementLines(in []string) ([]CraneMovement, error) {
	reg, err := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	if err != nil {
		return nil, err
	}

	movements := make([]CraneMovement, len(in))
	for i, s := range in {
		groups := reg.FindStringSubmatch(s)
		count, err := strconv.Atoi(groups[1])
		if err != nil {
			return nil, err
		}
		from, err := strconv.Atoi(groups[2])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(groups[3])
		if err != nil {
			return nil, err
		}
		movements[i] = CraneMovement{
			Count: count,
			From:  from,
			To:    to,
		}
	}
	return movements, nil
}
