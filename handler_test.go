package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHandlerCorrect(t *testing.T) {
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

func TestHandlerIncorrect(t *testing.T) {
	for _, tc := range []struct {
		name       string
		expression string
	}{
		{name: "No operators", expression: "2 4 8 16"},
		{name: "Incorrect operands", expression: "5 a + b x - *"},
		{name: "Incorrect operators", expression: "2 5 | 3 4 & |"},
		{name: "Expression as word", expression: "expression"},
	} {
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
