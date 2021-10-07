package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type CSVService struct {
	file string
}

func New(filename string) CSVService {
	return CSVService{file: filename}
}

// GetData - Reads file from a given path, returns the slice of records
func (c CSVService) GetData() ([][]string, error) {
	file, err := os.Open(c.file)
	if err != nil {
		log.Println(err.Error())
		return [][]string{}, errors.New("an error ocurred while opening the csv file")
	}
	defer file.Close()

	return readFile(file)
}

func (c CSVService) WriteALLData(records [][]string) error {
	file, err := os.OpenFile(c.file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
		return errors.New("an error ocurred while opening the csv file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	// Using WriteAll
	var data [][]string
	data = append(data, records...)
	err = w.WriteAll(data)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("an error ocurred while writing the csv file")
	}
	return nil
}

func (c CSVService) WriteRowData(record []string) error {
	file, err := os.OpenFile(c.file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
		return errors.New("an error ocurred while opening the csv file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	if err := w.Write(record); err != nil {
		log.Println(err.Error())
		return errors.New("an error ocurred while writing the csv file")
	}
	return nil
}

// ValidateFile - Reads file from a given path, returns the if  to the file or error
func (c CSVService) ValidateFile() {
	file, err := os.Open(c.file)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
}

// readFile Reads the file and retunrs the data
func readFile(file *os.File) ([][]string, error) {
	Records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		return [][]string{}, errors.New("an error ocurred while reading the csv file")
	}
	return Records, nil
}
