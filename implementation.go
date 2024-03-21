package lab2

import (
	"fmt"
	"strings"
)

func PostfixToPrefix(postfixExpression string) (string, error) {

	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
	}
	var stack []string

	tokens := strings.Split(postfixExpression, " ")
	for _, token := range tokens {
		if operators[token] {
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, fmt.Sprintf("%s %s %s", token, operand1, operand2))
		} else {
			stack = append(stack, token)
		}
	}

	return stack[0], nil
}
