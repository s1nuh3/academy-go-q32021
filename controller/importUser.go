package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/s1nuh3/academy-go-q32021/model"
)

type CtrlImportInterface interface {
	ImportUserCtrl(id int) (*model.Users, error)
}

type ImportHandler struct {
	us CtrlImportInterface
}

func NewImportHandler(u CtrlImportInterface) ImportHandler {
	return ImportHandler{u}
}

// GetUsersbyId - Look up for a user id
func (c ImportHandler) ImportHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		returnError(w, r, errors.New("ID provided is not valid"), 400)
		return
	}
	u, err := c.us.ImportUserCtrl(id)
	if err != nil {
		returnError(w, r, err, 500)
		return
	}
	if u.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}

}
