package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compareFiles(outputFileName, expectedOutputFileName string) error {

	outputFile, err := os.Open(outputFileName)
	if err != nil {
		return fmt.Errorf("error opening output file: %v", err)
	}
	defer outputFile.Close()

	expectedOutputFile, err := os.Open(expectedOutputFileName)
	if err != nil {
		return fmt.Errorf("error opening expected output file: %v", err)
	}
	defer expectedOutputFile.Close()

	outputScanner := bufio.NewScanner(outputFile)
	expectedScanner := bufio.NewScanner(expectedOutputFile)

	lineNum := 1
	for outputScanner.Scan() {
		if !expectedScanner.Scan() {
			return fmt.Errorf("mismatch at line %d: expected more lines in expected output", lineNum)
		}

		outputLine := outputScanner.Text()
		expectedLine := expectedScanner.Text()

		if outputLine != expectedLine {
			return fmt.Errorf("mismatch at line %d: expected '%s', got '%s'", lineNum, expectedLine, outputLine)
		}
		lineNum++
	}

	if expectedScanner.Scan() {
		return fmt.Errorf("mismatch at line %d: expected more lines in output", lineNum)
	}

	return nil
}

func test(inputFileName string, outputFileName string, expectedOutputFileName string) {

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	intStack, _ := New[int](16)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Fields(line)

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "Push":
			if len(command) > 1 {
				val, _ := strconv.Atoi(command[1])
				err := intStack.Push(val)
				if err != nil {
					fmt.Fprintln(outputFile, err.Error())
				}
			} else {
				fmt.Fprintf(outputFile, "Invalid number: %s\n", command[1])
			}
		case "Pop":
			if err := intStack.Pop(); err != nil {
				fmt.Fprintln(outputFile, err)
			}
		case "Top":
			topElement, err := intStack.Top()
			if err != nil {
				fmt.Fprintln(outputFile, err)
			} else {
				fmt.Fprintln(outputFile, *topElement)
			}
		case "Size":
			size := intStack.Size()
			fmt.Fprintln(outputFile, size)
		case "IsEmpty":
			isEmpty := intStack.IsEmpty()
			fmt.Fprintln(outputFile, isEmpty)
		default:
			fmt.Fprintf(outputFile, "Unknown command: %s\n", command[0])
		}
	}

	if err := compareFiles(outputFileName, expectedOutputFileName); err != nil {
		fmt.Println("Test failed:", err)
	} else {
		fmt.Println("Test passed!")
	}
}
