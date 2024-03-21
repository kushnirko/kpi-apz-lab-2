package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCorrectInput(t *testing.T) {
	input := strings.NewReader("6 3 ^ 3 4 ^ - 5 +")
	output := bytes.NewBuffer(nil)
	expected := "+ - ^ 6 3 ^ 3 4 5"

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()

	assert.Nil(t, err)
	assert.Equal(t, expected, output.String())
}

func TestIncorrectExpressions(t *testing.T) {
	type testCase struct {
		name       string
		expression string
	}

	testCases := []testCase{
		{name: "No operands", expression: "2 4 8 16"},
		{name: "Incorrect operator", expression: "5 a + b x - *"},
		{name: "Incorrect operands", expression: "2 5 | 3 4 & |"},
		{name: "Expression as word", expression: "expression"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.expression)
			output := bytes.NewBuffer(nil)

			handler := ComputeHandler{
				Input:  input,
				Output: output,
			}
			err := handler.Compute()

			assert.NotNil(t, err)
		})
	}
}
