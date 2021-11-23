package repository

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	ErrWritingFile = "an error ocurred while writing the csv file"
	ErrReadingFile = "an error ocurred while reading the csv file"
)

//CSVService - Struc to implement the interfaces that access csv file
type CSVService struct {
	file *os.File
}

//New - Creates a new instance to access a csv file, recieves the os file
func New(file *os.File) *CSVService {
	return &CSVService{file: file}
}

// GetData - Reads file from a given path, returns the slice of records
func (c CSVService) GetData() ([][]string, error) {
	return readFile(c.file)
}

//WriteALLData - Appends a range of rows of new data to CSV File
func (c CSVService) WriteALLData(records [][]string) error {
	w := csv.NewWriter(c.file)
	defer w.Flush()
	var data [][]string
	data = append(data, records...)
	err := w.WriteAll(data)
	if err != nil {
		return fmt.Errorf(ErrWritingFile+" %w", err)
	}
	return nil
}

//WriteRowData - Appends a row of new data to CSV File
func (c CSVService) WriteRowData(record []string) error {
	w := csv.NewWriter(c.file)
	defer w.Flush()
	if err := w.Write(record); err != nil {
		return fmt.Errorf(ErrWritingFile+" %w", err)
	}
	return nil
}

//readFile Reads the file and retunrs the data
func readFile(file *os.File) ([][]string, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return [][]string{}, fmt.Errorf(ErrReadingFile+" %w", err)
	}
	Records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, fmt.Errorf(ErrReadingFile+" %w", err)
	}
	return Records, nil
}
