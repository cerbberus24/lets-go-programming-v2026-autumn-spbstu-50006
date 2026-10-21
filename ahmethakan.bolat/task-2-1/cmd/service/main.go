package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	minTemp    = 15
	maxTemp    = 30
	noSolution = -1
)

type Operator string

const (
	GreaterOrEqual Operator = ">="
	LessOrEqual    Operator = "<="
)

var errUnknownOperator = errors.New("unknown operator")

type Temperature struct {
	low  int
	high int
}

func NewTemperature() *Temperature {
	return &Temperature{
		low:  minTemp,
		high: maxTemp,
	}
}

func (t *Temperature) Update(operator Operator, value int) error {
	switch operator {
	case GreaterOrEqual:
		t.low = max(t.low, value)
	case LessOrEqual:
		t.high = min(t.high, value)
	default:
		return fmt.Errorf("%w: %q", errUnknownOperator, operator)
	}

	return nil
}

func (t *Temperature) Optimal() int {
	if t.low > t.high {
		return noSolution
	}

	return t.low
}

func readCount(reader *bufio.Reader) (int, error) {
	var count int

	_, err := fmt.Fscan(reader, &count)
	if err != nil {
		return 0, fmt.Errorf("read count: %w", err)
	}

	return count, nil
}

func readCondition(reader *bufio.Reader) (Operator, int, error) {
	var (
		operator string
		value    int
	)

	_, err := fmt.Fscan(reader, &operator, &value)
	if err != nil {
		return "", 0, fmt.Errorf("read condition: %w", err)
	}

	return Operator(operator), value, nil
}

func processDepartment(reader *bufio.Reader, employees int) error {
	temperature := NewTemperature()

	for range employees {
		operator, value, err := readCondition(reader)
		if err != nil {
			return err
		}

		err = temperature.Update(operator, value)
		if err != nil {
			return fmt.Errorf("update temperature: %w", err)
		}

		fmt.Println(temperature.Optimal())
	}

	return nil
}

func run(reader *bufio.Reader) error {
	departments, err := readCount(reader)
	if err != nil {
		return fmt.Errorf("departments: %w", err)
	}

	for range departments {
		employees, err := readCount(reader)
		if err != nil {
			return fmt.Errorf("employees: %w", err)
		}

		err = processDepartment(reader, employees)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
