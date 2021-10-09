package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

//UseCaseImportUser - Interface to be implemented on usecase layer
type UseCaseImportUser interface {
	ImportUserUC(id int) (*model.Users, error)
}

//ImportHandler - Struc to implement the user handlers
type ImportHandler struct {
	uci UseCaseImportUser
}

//New - Creates a new instance of handlers for the import user paths
func NewImportHandler(ui UseCaseImportUser) ImportHandler {
	return ImportHandler{ui}
}

// ImportHandler - Handles the call to import a new user
func (ih ImportHandler) ImportHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		returnError(w, r, errors.New("ID provided is not valid"), 400)
		return
	}
	u, err := ih.uci.ImportUserUC(id)
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
