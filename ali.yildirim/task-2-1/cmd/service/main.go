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
	MinTemp      = 15
	MaxTemp      = 30
	NoSolution   = -1
	ExpectedArgs = 2
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
	scanner.Scan()

	firstLine := strings.Fields(scanner.Text())
	if len(firstLine) < ExpectedArgs {
		return
	}

	departments, _ := strconv.Atoi(firstLine[0])
	employees, _ := strconv.Atoi(firstLine[1])

	for range departments {
		temp := NewTemperatureRange()

		for range employees {
			scanner.Scan()
			line := strings.TrimSpace(scanner.Text())

			if line == "" {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) != ExpectedArgs {
				continue
			}

			operator := Operator(parts[0])
			val, _ := strconv.Atoi(parts[1])

			err := temp.Update(operator, val)
			if err != nil {
				fmt.Println(temp.Optimal())

				continue
			}

			fmt.Println(temp.Optimal())
		}
	}
}
