package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufReader := bufio.NewReader(ch.input)
	bufWriter := bufio.NewWriter(ch.output)

	postfix, err := bufReader.ReadString('\n')
	if err != nil {
		return err
	}

	//TODO change function to desired
	prefix, err := PrefixToPostfix(postfix)
	if err != nil {
		return err
	}

	_, err = bufWriter.WriteString(prefix)
	if err != nil {
		return err
	}

	err = bufWriter.Flush()
	if err != nil {
		return err
	}

	return nil
}
