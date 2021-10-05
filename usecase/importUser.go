package usecase

import "github.com/s1nuh3/academy-go-q32021/model"

//CtrlImpUser - To implement the ConsumeAPI interface
type CtrlImpUser struct {
	client ConsumeAPI
}

//NewImport create service pass it to Controller
func NewConsumeAPI(c ConsumeAPI) *CtrlImpUser {
	return &CtrlImpUser{
		client: c,
	}
}

// GetExternalUserCtrl - Returns a user by id if it's found in a csv file
func (s *CtrlImpUser) ImportUserCtrl(id int) (*model.Users, error) {
	return s.client.ImportUser(id)
}
