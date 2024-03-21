package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixToPrefix(t *testing.T) {
	result, err := PostfixToPrefix("")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "expression not found")

	result, err = PostfixToPrefix("4 2 7 8 ! ? @")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "invalid character entered")

	result, err = PostfixToPrefix("8 5+")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "invalid character entered")

	result, err = PostfixToPrefix("4 2 - + -")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "postfix expression entered incorrectly")

	result, err = PostfixToPrefix("4 2 5")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "postfix expression entered incorrectly")

	result, err = PostfixToPrefix("4.1 2.5 +")
	assert.Equal(t, "+ 4.1 2.5", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("6 2 * 3 -")
	assert.Equal(t, "- * 6 2 3", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("41 23 10 / +")
	assert.Equal(t, "+ 41 / 23 10", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("4 2 - 3 * 5 +")
	assert.Equal(t, "+ * - 4 2 3 5", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("6 5 4 - * 2 /")
	assert.Equal(t, "/ * 6 - 5 4 2", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("9 7 + 1 3 + / 2 * 3 - 8 +")
	assert.Equal(t, "+ - * / + 9 7 + 1 3 2 3 8", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("7 2 * 3 - 8 + 1 - 5 4 * 2 / +")
	assert.Equal(t, "+ - + - * 7 2 3 8 1 / * 5 4 2", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("2 5 ^ 9 - 12 + 6 2 / - 4 + 2 - 9 +")
	assert.Equal(t, "+ - + - + - ^ 2 5 9 12 / 6 2 4 2 9", result)
	assert.NoError(t, err)

	result, err = PostfixToPrefix("6 2 / 2 * 2 - 8 2 ^ + 2 3 * - 7 3 ^ +")
	assert.Equal(t, "+ - + - * / 6 2 2 2 ^ 8 2 * 2 3 ^ 7 3", result)
	assert.NoError(t, err)
}

func ExamplePostfixToPrefix() {
	prefixExpression, _ := PostfixToPrefix("6 3 ^ 3 4 ^ - 5 +")
	fmt.Println(prefixExpression)

	// Output:
	// + - ^ 6 3 ^ 3 4 5
}
