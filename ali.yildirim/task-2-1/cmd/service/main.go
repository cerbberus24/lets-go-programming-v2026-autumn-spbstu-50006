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

func (tr *TemperatureRange) Update(op Operator, val int) error {
	switch op {
	case OpGreaterOrEqual:
		if val > tr.low {
			tr.low = val
		}
	case OpLessOrEqual:
		if val < tr.high {
			tr.high = val
		}
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperator, op)
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

	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	if len(firstLine) < 2 {
		return
	}

	n, _ := strconv.Atoi(firstLine[0])
	k, _ := strconv.Atoi(firstLine[1])

	for i := 0; i < n; i++ {
		temp := NewTemperatureRange()

		for j := 0; j < k; j++ {
			scanner.Scan()
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) != 2 {
				continue
			}

			op := Operator(parts[0])
			val, _ := strconv.Atoi(parts[1])

			err := temp.Update(op, val)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}

			fmt.Println(temp.Optimal())
		}
	}
}
