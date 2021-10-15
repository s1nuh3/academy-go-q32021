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

//GoRoutineService - Implementacion for the Contrac to call the workpool
type GoRoutineService struct {
	file *os.File
}

//LoadBalancer - Struck to pass the worker ID and load ( items to proccess)
type LoadBalancer struct {
	worker int
	load   int
}

//New - Creates an new instance to a access Users, receive repo csv
func New(file *os.File) *GoRoutineService {
	return &GoRoutineService{file: file}
}

//WorkPool - Creates an concrete workpool of go routines to procces the read lines in a cvs file using channels
func (gs GoRoutineService) WorkPool(filter, items, itemsPerWorker, workers int) (*[]model.Users, error) {
	fcsv := csv.NewReader(gs.file)
	_, err := gs.file.Seek(0, 0)
	if err != nil {
		log.Printf("An error happend at file: %v", err)
	}
	rs := make(chan *model.Users)
	lines := make(chan []string)
	var u []model.Users
	var wg sync.WaitGroup

	loads := calculateWorkLoad(workers, itemsPerWorker, items)
	wg.Add(workers)
	fmt.Printf("Starting [%d] Workers  \n", workers)
	for w := 0; w < workers; w++ {
		go func(ID, task int, lines <-chan []string, rs chan<- *model.Users, wg *sync.WaitGroup) {
			GetRecords(loads[ID], filter, lines, rs, wg)
		}(w, itemsPerWorker, lines, rs, &wg)
	}
	go ReadCSVLines(fcsv, lines)

	go func() {
		wg.Wait()
		close(rs)
		fmt.Printf("Workerpool finished \n")
	}()

	for user := range rs {
		u = append(u, *user)
	}
	return &u, nil
}

func calculateWorkLoad(workers int, itemsPerWorker int, items int) []LoadBalancer {
	extraItems := (workers * itemsPerWorker) - items
	var loads []LoadBalancer
	for i := 0; i < workers; i++ {
		if i == 0 && extraItems != 0 {
			loads = append(loads, LoadBalancer{
				worker: i,
				load:   itemsPerWorker - extraItems,
			})
		} else {
			loads = append(loads, LoadBalancer{
				worker: i,
				load:   itemsPerWorker,
			})
		}
	}
	return loads
}

//ReadCSVLines - Reads a cvs file and pushes it to a channel , This is meant to be run a a go routine
func ReadCSVLines(fcsv *csv.Reader, lines chan []string) {
	for {
		rStr, err := fcsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("ERROR: ", err.Error())
			break
		}
		lines <- rStr
	}
	close(lines)
}

//GetRecords - This function acts as a worker in the workpool, fetch the channel for a string
// parses it as a model user, validates if comply with the filter, and then pushes it to a result channel,
// Also controls the item per worked limit
func GetRecords(load LoadBalancer, filter int, lines <-chan []string, rs chan<- *model.Users, wg *sync.WaitGroup) {
	defer wg.Done()
	completed := 0
	for job := range lines {
		if completed == load.load {
			fmt.Printf(" Worker [%d] Completed [%d] of [%d] \n", load.worker, completed, load.load)
			break
		}
		aux, err := user.ParseUserRecord(job)
		if err != nil {
			log.Printf("Failed to Parse User Record %v, error %v\n", job, err)
		} else {
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
		}
	}
}
