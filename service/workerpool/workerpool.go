package workerpool

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sync"

	"github.com/s1nuh3/academy-go-q32021/model"
	"github.com/s1nuh3/academy-go-q32021/service/user"
)

type GoRoutineService struct {
	file *os.File
}

//New - Creates an new instance to a access Users, receive repo csv
func New(file *os.File) *GoRoutineService {
	return &GoRoutineService{file: file}
}

//WorkPool - Creates an concrete workpool of go routines to procces the read lines in a cvs file using channels
func (gs GoRoutineService) WorkPool(filter, items, itemsPerWorker, workers int) (*[]model.Users, error) {
	gs.file.Seek(0, 0)
	fcsv := csv.NewReader(gs.file)
	rs := make(chan *model.Users)
	lines := make(chan []string)
	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		fmt.Printf("Launch Worker [%d] \n", w)
		go func(ID, task int, lines <-chan []string, rs chan<- *model.Users, wg *sync.WaitGroup) {
			GetRecords(ID, filter, itemsPerWorker, lines, rs, wg)
		}(w, itemsPerWorker, lines, rs, &wg)
	}

	go ReadCSVLines(fcsv, lines)

	go func() {
		wg.Wait()
		close(rs)
	}()
	var u []model.Users
	for i := 1; i <= items; i++ {
		user, ok := <-rs
		if !ok {
			break
		}
		u = append(u, *user)
		//fmt.Printf("Job %v has been finished with result ID %v\n", i, res.ID)
	}
	return &u, nil
}

//ReadCSVLines - Reads a cvs file and pushes it to a channel , This is meant to be run a a go routine
func ReadCSVLines(fcsv *csv.Reader, lines chan []string) {
	for {
		rStr, err := fcsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		lines <- rStr
	}
	close(lines)
}

//GetRecords - This function acts as a worker in the workpool, fetch the channel for a string
// parses it as a model user, validates if comply with the filter, and then pushes it to a result channel,
// Also controls the item per worked limit
func GetRecords(ID, filter, task int, lines <-chan []string, rs chan<- *model.Users, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Printf("Worker %d waiting for job \n", ID)
	completed := 1
	if completed <= task {
		for job := range lines {
			fmt.Printf("Worker [%d] started inner job  [%d] \n", ID, completed)
			aux, err := user.ParseUserRecord(job)
			if err != nil {
				log.Printf("Failed to Parsr User Record %v, error %v\n", job, err)
			}
			switch filter {
			case 0:
				if math.Remainder(float64(aux.ID), 2) == 0 {
					rs <- aux
					completed++
				}
			case 1:

				if math.Remainder(float64(aux.ID), 2) != 0 {
					rs <- aux
					completed++
				}
			}

			if completed > task {
				break
			}

		}
	}
}
