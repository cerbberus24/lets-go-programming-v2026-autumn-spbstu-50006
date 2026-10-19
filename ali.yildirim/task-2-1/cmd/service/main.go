package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	var departments int

	var employees int

	_, err := fmt.Fscan(reader, &departments, &employees)
	if err != nil {
		return
	}

	for range departments {
		temp := NewTemperatureRange()

		for range employees {
			var opStr string

			var val int

			_, err := fmt.Fscan(reader, &opStr, &val)
			if err != nil {
				return
			}

			operator := Operator(opStr)

			err = temp.Update(operator, val)
			if err != nil {
				fmt.Println(NoSolution)

				continue
			}

			fmt.Println(temp.Optimal())
		}
	}
}
