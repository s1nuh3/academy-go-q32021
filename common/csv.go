package common

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

// GetData - Reads file from a given path, returns the slice of records
func GetData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return [][]string{}, errors.New("an error ocurred while opening the csv file")
	}
	defer file.Close()

	return ReadFile(file)
}

func WriteALLData(filename string, records [][]string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
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

func WriteRowData(filename string, record []string) error {
	fmt.Println("Record: ", record)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("an error ocurred while opening the csv file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	if err := w.Write(record); err != nil {
		fmt.Println(err.Error())
		return errors.New("an error ocurred while writing the csv file")
	}
	w.Flush()
	return nil
}

// ReadFile Reads the file and retunrs the data
func ReadFile(file *os.File) ([][]string, error) {
	fmt.Println("Open File", file.Name())
	Records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return [][]string{}, errors.New("an error ocurred while reading the csv file")
	}
	return Records, nil
}

// ValidateFile - Reads file from a given path, returns the if  to the file or error
func ValidateFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()
	return nil
}
