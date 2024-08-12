package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ip
type NumberData struct {
	numbers []int
}
type ProcessedData struct {
	numbers []int
}

func NewNumberData() *NumberData {
	return &NumberData{numbers: []int{}}
}

func (nd *NumberData) AddNumber(num int) {
	nd.numbers = append(nd.numbers, num)
}

func (nd *NumberData) Process() ProcessedData {
	processed := ProcessedData{numbers: make([]int, len(nd.numbers))}
	for i, num := range nd.numbers {
		processed.numbers[i] = num * 2
	}
	return processed
}

type FileProcesser struct {
	inputFile  string
	outputFile string
}

func NewFileProcessor(input, output string) *FileProcesser {
	return &FileProcesser{
		inputFile:  input,
		outputFile: output,
	}
}

func (fp *FileProcesser) ReadAndDeserialize() (*NumberData, error) {
	file, err := os.Open(fp.inputFile)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}
	defer file.Close()
	data := NewNumberData()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error converting to number: %w", err)
		}
		data.AddNumber(num)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %w", err)
	}

	return data, nil
}

func (fp *FileProcesser) Process() error {
	// Read and deserialize input data
	inputData, err := fp.ReadAndDeserialize()
	if err != nil {
		return err
	}

	// Process data
	processedData := inputData.Process()

	// Serialize and write output data
	err = fp.SerializeAndWrite(processedData)
	if err != nil {
		return err
	}

	return nil
}

func (fp *FileProcesser) SerializeAndWrite(data ProcessedData) error {
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, num := range data.numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}
	return nil
}

func main() {
	processor := NewFileProcessor("input.txt", "output.txt")
	if err := processor.Process(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Processing completed successfully.")
}
