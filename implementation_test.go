package lab2

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixToPrefix(t *testing.T) {
	for _, tc := range []struct {
		name        string
		expression  string
		expectedRes string
		expectedErr error
	}{
		{
			name:        "Empty string",
			expression:  "",
			expectedRes: "",
			expectedErr: errors.New("😈"),
		},
		{
			name:        "Invalid character",
			expression:  "4 2 7 8 ! ? @",
			expectedRes: "",
			expectedErr: errors.New("invalid character entered"),
		},
		{
			name:        "No whitespace",
			expression:  "8 5+",
			expectedRes: "",
			expectedErr: errors.New("invalid character entered"),
		},
		{
			name:        "Too many operators",
			expression:  "4 2 - + -",
			expectedRes: "",
			expectedErr: errors.New("postfix expression entered incorrectly"),
		},
		{
			name:        "No operators",
			expression:  "4 2 5",
			expectedRes: "",
			expectedErr: errors.New("postfix expression entered incorrectly"),
		},
		{
			name:        "Fractional numbers",
			expression:  "4.1 2.5 +",
			expectedRes: "+ 4.1 2.5",
			expectedErr: nil,
		},
		{
			name:        "Short expression",
			expression:  "41 23 10 / +",
			expectedRes: "+ 41 / 23 10",
			expectedErr: nil,
		},
		{
			name:        "Medium expression",
			expression:  "9 7 + 1 3 + / 2 * 3 - 8 +",
			expectedRes: "+ - * / + 9 7 + 1 3 2 3 8",
			expectedErr: nil,
		},
		{
			name:        "Long expression",
			expression:  "6 2 / 2 * 2 - 8 2 ^ + 2 3 * - 7 3 ^ +",
			expectedRes: "+ - + - * / 6 2 2 2 ^ 8 2 * 2 3 ^ 7 3",
			expectedErr: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := PostfixToPrefix(tc.expression)
			assert.Equal(t, tc.expectedRes, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("6 3 ^ 3 4 ^ - 5 +")
	fmt.Println(res)

	// Output:
	// + - ^ 6 3 ^ 3 4 5
}
