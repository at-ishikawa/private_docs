package internal

import (
	"unicode"
)

func eval(result int, operator byte, sign int, number int) int {
	right := sign * number
	// fmt.Printf("%d %c %d\n", result, operator, right)
	switch operator {
	case '+':
		return result + right
	case '-':
		return result - right
	case '*':
		return result * right
	}
	return result
}

type Token struct {
	value    int
	operator byte
	sign     int
}

func calculate(s string) int {
	stack := make([]Token, 0)

	// the first operand should be added to the result
	operator := byte('+')

	result := 0
	currentNumber := 0
	sign := 1
	for index := 0; index < len(s); index++ {
		sCh := byte(s[index])
		if unicode.IsDigit(rune(s[index])) {
			currentNumber = currentNumber*10 + int(sCh-'0')
			continue
		}
		switch sCh {
		case '+', '-', '*', '/':
			result = eval(result, operator, sign, currentNumber)
			currentNumber = 0
			if sCh == '-' {
				operator = '+'
				sign = -1
			} else {
				operator = sCh
				sign = 1
			}

		case '(':
			stack = append(stack, Token{value: result}, Token{operator: operator}, Token{sign: sign})
			sign = 1
			result = 0
			operator = '+'
		case ')':
			result = eval(result, operator, sign, currentNumber)
			currentNumber = 0

			parenthesisSign := stack[len(stack)-1].sign // sign
			operator = stack[len(stack)-2].operator
			left := stack[len(stack)-3].value // operand

			result = eval(left, operator, parenthesisSign, result)
			stack = stack[:len(stack)-3]
		}
	}
	if currentNumber != 0 {
		result = eval(result, operator, sign, currentNumber)
	}
	return result
}

/*
import (
	"fmt"
	"unicode"
)

func sum(array []int) int {
	ans := 0
	for _, elem := range array {
		ans = ans + elem
	}
	return ans
}

func eval(s string) (int, int) {
	index := 0
	currentNumber := 0
	stack := make([]int, 0)
	sign := byte('+')

	update := func(operator byte, num int) {
		fmt.Printf("%d %c\n", num, operator)
		switch operator {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		}
	}

	for ; index < len(s); index++ {
		sCh := s[index]
		if unicode.IsDigit(rune(sCh)) {
			currentNumber = currentNumber*10 + int(sCh-'0')
			continue
		}
		switch sCh {
		case '+', '-':
			update(sign, currentNumber)
			currentNumber = 0
			sign = sCh
		case '(':
			currentNumber, index = eval(s[index+1:])
		case ')':
			update(sign, currentNumber)
			return sum(stack), index
		}
	}
	update(sign, currentNumber)
	return sum(stack), 0
}

func calculate(s string) int {
	res, _ := eval(s)
	return res
}
*/
/*
func eval(left int, operator byte, right int) int {
	fmt.Printf("%d %c %d\n", left, operator, right)

	switch operator {
	case '+':
		return left + right
	case '-':
		return left - right
	}

	return right
}

func calculate(s string) int {
	result := 0
	var operator byte
	currentNumber := 0
	sign := 1
	isOperand := true

	for index := 0; index < len(s); index++ {
		sCh := byte(s[index])
		switch sCh {
		case '+', '-':
			if isOperand {
				sign = -1
			} else {
				result = sign * currentNumber
				currentNumber = 0
				sign = 1
				operator = sCh
			}

		case '(':
			answer := calculate(s[index+1:])
			result = sign * eval(result, operator, answer)
			sign = 1
			currentNumber = 0
			operator = '0'
			for ; s[index] != ')'; index++ {
			}

		case ')':
			return eval(result, operator, sign*currentNumber)

		default:
			digit := int(sCh - '0')
			currentNumber = currentNumber*10 + digit
		}
	}

	return eval(result, operator, currentNumber)
}
*/

/*
type Expression struct {
	operator    byte // + or -
	value       int
	sign        int
	parenthesis bool
}


func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
	// return ch <= '0' && ch >= '9'
}

func calculateExpression(stack []Expression) ([]Expression, int) {
	if len(stack) == 1 {
		return nil, stack[0].value
	}

	endIndex := len(stack) - 1
	right := stack[endIndex].value
	if stack[endIndex-1].parenthesis {
		return stack[:endIndex-1], right
	}

	operator := stack[endIndex-1].operator
	left := stack[endIndex-2].value

	ans := 0
	if operator == '+' {
		ans = left + right
	} else if operator == '-' {
		ans = left - right
	}
	fmt.Printf("%d %c %d = %d\n", left, operator, right, ans)

	return stack[:endIndex-2], ans
}

// Calculate recursively
func calculate(s string) int {
	stack := make([]Expression, 0)

	isOperand := true
	currentNumber := 0
	sign := 1
	for index, sCh := range s {
		if index > 0 && isDigit(rune(s[index-1])) && !isDigit(sCh) {
			stack = append(stack, Expression{
				value: sign * currentNumber,
			})
			currentNumber = 0
			sign = 1
			isOperand = false
		}

		switch sCh {
		case ' ':

		case '+':
			stack = append(stack, Expression{
				operator: '+',
			})
			isOperand = true

		case '-':

			if isOperand {
				sign = -1
			} else {
				stack = append(stack, Expression{
					operator: '+',
				})
				isOperand = true
			}

		case '(':
			isOperand = true
			stack = append(stack, Expression{
				parenthesis: true,
				sign:        sign,
			})
			sign = 1

		case ')':
			ans := stack[len(stack)-1].value
			for !stack[len(stack)-2].parenthesis {
				stack, ans = calculateExpression(stack)
				stack = append(stack, Expression{
					value: ans,
				})
			}
			sign = stack[len(stack)-2].sign
			stack = stack[:len(stack)-2]
			stack = append(stack, Expression{
				value: sign * ans,
			})
			isOperand = false
			sign = 1

		default:
			digit := int(sCh - '0')
			currentNumber = currentNumber*10 + digit
		}
	}
	if isOperand {
		stack = append(stack, Expression{
			value: sign * currentNumber,
		})
	}

	_, ans := calculateExpression(stack)
	return ans
}
*/
