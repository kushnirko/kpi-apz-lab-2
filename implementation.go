package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PostfixToPrefix(postfixExpression string) (string, error) {
	if postfixExpression == "" {
		return "", fmt.Errorf("expression not found")
	}

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
		if token == "" {
			return "", fmt.Errorf("postfix expression entered incorrectly")
		}
		if !operators[token] && !isNumeric(token) {
			return "", fmt.Errorf("invalid character entered")
		}
		if operators[token] {
			if len(stack) < 2 {
				return "", fmt.Errorf("postfix expression entered incorrectly")
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, fmt.Sprintf("%s %s %s", token, operand1, operand2))
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("postfix expression entered incorrectly")
	}

	return stack[0], nil
}

func isNumeric(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}
