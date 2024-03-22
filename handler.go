package lab2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufReader := bufio.NewReader(ch.Input)
	bufWriter := bufio.NewWriter(ch.Output)

	postfix, err := bufReader.ReadString('\n')

	//TODO change function to desired
	prefix, err := PostfixToPrefix(postfix)
	if err != nil {
		return err
	}

	bufWriter.WriteString(prefix)
	bufWriter.Flush()

	return nil
}
