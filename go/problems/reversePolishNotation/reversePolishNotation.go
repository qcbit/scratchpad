package reversepolishnotation

import (
	"regexp"
	"strconv"
	"strings"
)

type Calculator struct {
	field []float64
	size  int
}

type RPN interface {
	Calculator
	pop()
	push(x float64)
	calculate() float64
}

func NewCalculator(k int) *Calculator {
	return &Calculator{
		field: make([]float64, 0, k),
		size:  k,
	}
}

func (c *Calculator) push(x float64) bool {
	if len(c.field) == c.size {
		return false
	}
	c.field = append(c.field, x)
	return true
}

func (c *Calculator) pop() (float64, bool) {
	if len(c.field) == 0 {
		return 0, false
	}
	x := c.field[len(c.field)-1]
	c.field = c.field[:len(c.field)-1]
	return x, true
}

func (calc *Calculator) calculate(rpnInput string) float64 {
	for _, c := range strings.Split(rpnInput, " ") {
		re := regexp.MustCompile(`^-?\d+(?:\.\d+)?$`)
		if re.MatchString(c) {
			f, _ := strconv.ParseFloat(c, 64)
			if ok := calc.push(f); !ok {
				panic("Pushing number caused a panic. Stack likely full.")
			}
		} else {
			op2, ok := calc.pop()
			if !ok {
				panic("Stack empty. Nothing to pop.")
			}
			op1, ok := calc.pop()
			if !ok {
				panic("Stack empty. No second operand found.")
			}
			switch c {
			case "-":
				calc.push(op1 - op2)
			case "+":
				calc.push(op1 + op2)
			case "/":
				calc.push(op1 / op2)
			case "*":
				calc.push(op1 * op2)
			default:
				panic("unrecognized operator")
			}
		}
	}
	ans, _ := calc.pop()
	return ans
}
