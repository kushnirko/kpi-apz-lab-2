package main

import (
	"flag"
	"fmt"
	lab2 "github.com/kushnirko/kpi-apz-lab-2.git"
	"io"
	"os"
	"strings"
)

var (
	inputExpression    = flag.String("e", "", "Expression to compute")
	expressionLocation = flag.String("f", "", "Location of the expression to compute")
	resultOutput       = flag.String("o", "", "Where to place conversion results")
)

func GetInput() (io.Reader, error) {
	var inp io.Reader
	var err error

	if *inputExpression != "" {
		inp = strings.NewReader(*inputExpression)
	}

	if *expressionLocation != "" {
		inp, err = os.Open(*expressionLocation)
		if err != nil {
			return nil, err
		}
	}

	return inp, nil
}

func GetOutput() (io.Writer, error) {
	if *resultOutput == "" {
		return os.Stdout, nil
	}

	_, err := os.Stat(*resultOutput)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(*resultOutput)
			if err != nil {
				return os.Stdout, err
			}
			return file, nil
		} else {
			return os.Stdout, err
		}
	}

	file, err := os.OpenFile(*resultOutput, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return os.Stdout, err
	}

	return file, nil
}

func main() {
	flag.Parse()
	input, err := GetInput()
	if err != nil {
		fmt.Println("File error:", err)
		return
	}

	output, err := GetOutput()
	if err != nil {
		fmt.Println("File error:", err)
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err = handler.Compute()
	if err != nil {
		fmt.Println("Args error:", err)
	}
}
