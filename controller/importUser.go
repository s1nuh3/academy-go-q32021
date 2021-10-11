package controller

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusBadRequest)
		json, err := json.Marshal("ID provided is not valid")
		if err != nil {
			returnError(w, r, err, 500)
			return
		}
		w.Write(json)
		return
	}
	u, err := ih.uci.ImportUserUC(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json, err := json.Marshal(err.Error())
		if err != nil {
			returnError(w, r, err, 500)
			return
		}
		w.Write(json)
		return
	}
	if u.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json, err := json.Marshal("ID not found in external API")
		if err != nil {
			returnError(w, r, err, 500)
			return
		}
		w.Write(json)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json, err := json.Marshal(u)
		if err != nil {
			returnError(w, r, err, 500)
			return
		}
		w.Write(json)
	}

}
