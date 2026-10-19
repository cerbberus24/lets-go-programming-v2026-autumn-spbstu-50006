package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinTemp    = 15
	MaxTemp    = 30
	NoSolution = -1
)

type Operator string

const (
	OpGreaterOrEqual Operator = ">="
	OpLessOrEqual    Operator = "<="
)

var ErrUnknownOperator = errors.New("unknown operator")

type TemperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		low:  MinTemp,
		high: MaxTemp,
	}
}

func (tr *TemperatureRange) Update(operator Operator, val int) error {
	switch operator {
	case OpGreaterOrEqual:
		if val > tr.low {
			tr.low = val
		}
	case OpLessOrEqual:
		if val < tr.high {
			tr.high = val
		}
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperator, operator)
	}

	return nil
}

func (tr *TemperatureRange) Optimal() int {
	if tr.low <= tr.high {
		return tr.low
	}

	return NoSolution
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var tokens []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		tokens = append(tokens, fields...)
	}

	if len(tokens) < 2 {
		return
	}

	tokenIdx := 0
	nextInt := func() int {
		if tokenIdx >= len(tokens) {
			return 0
		}
		val, _ := strconv.Atoi(tokens[tokenIdx])
		tokenIdx++

		return val
	}

	nextStr := func() string {
		if tokenIdx >= len(tokens) {
			return ""
		}
		val := tokens[tokenIdx]
		tokenIdx++

		return val
	}

	departments := nextInt()
	employees := nextInt()

	for range departments {
		temp := NewTemperatureRange()

		for range employees {
			if tokenIdx+1 >= len(tokens) {
				return
			}

			opStr := nextStr()
			valStr := nextStr()

			operator := Operator(opStr)
			val, _ := strconv.Atoi(valStr)

			err := temp.Update(operator, val)
			if err != nil {
				fmt.Println(NoSolution)

				continue
			}

			fmt.Println(temp.Optimal())
		}
	}
}
