package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

type UseCaseGoRoutine interface {
	ReadConcurrent(filter, items, itemsPerWorker int) (*[]model.Users, error)
}

type GoRoutineHandler struct {
	ur UseCaseGoRoutine
}

//NewUser - Creates a new instance of handlers for the user paths
func NewGoRoutine(ur UseCaseGoRoutine) GoRoutineHandler {
	return GoRoutineHandler{ur}
}

//ConcurrentRead - Read concurrently the users file
func (gr GoRoutineHandler) GoRoutineHdl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filterType := vars["type"]
	items := vars["items"]
	itemsPerWorker := vars["items_per_worker"]

	w.Header().Set("Content-Type", "application/json")

	if filterType == "" {
		returnError(w, r, errors.New("type provided is not valid"), http.StatusBadRequest)
		return
	}

	filter := 0
	if filterType == "odd" {
		filter = 1
	}

	itms, err := strconv.Atoi(items)
	if err != nil {
		itms = 100
	}
	ipw, err := strconv.Atoi(itemsPerWorker)
	if err != nil {
		ipw = 50
	}

	if ipw > itms {
		returnError(w, r, errors.New("items per worker can't be higher than items"), http.StatusBadRequest)
		return
	}

	u, err := gr.ur.ReadConcurrent(filter, itms, ipw)
	if err != nil {
		returnError(w, r, err, http.StatusInternalServerError)
	}

	if cnt := len(*u); cnt != 0 {
		w.Header().Add("TotalResults", strconv.Itoa(cnt))
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			returnError(w, r, err, http.StatusInternalServerError)
		}
		return
	} else {
		jso, err := json.Marshal([]int{})
		if err != nil {
			returnError(w, r, err, http.StatusInternalServerError)
		}
		w.Write(jso)
	}
}
