package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type CSVService struct {
	file *os.File
}

func New(file *os.File) *CSVService {
	return &CSVService{file: file}
}

// GetData - Reads file from a given path, returns the slice of records
func (c CSVService) GetData() ([][]string, error) {
	return readFile(c.file)
}

func (c CSVService) WriteALLData(records [][]string) error {
	w := csv.NewWriter(c.file)
	defer w.Flush()
	// Using WriteAll
	var data [][]string
	data = append(data, records...)
	err := w.WriteAll(data)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("an error ocurred while writing the csv file")
	}
	return nil
}

func (c CSVService) WriteRowData(record []string) error {
	w := csv.NewWriter(c.file)
	defer w.Flush()
	if err := w.Write(record); err != nil {
		log.Println(err.Error())
		return errors.New("an error ocurred while writing the csv file")
	}
	return nil
}

// readFile Reads the file and retunrs the data
func readFile(file *os.File) ([][]string, error) {
	file.Seek(0, 0)
	Records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		return [][]string{}, errors.New("an error ocurred while reading the csv file")
	}
	return Records, nil
}
