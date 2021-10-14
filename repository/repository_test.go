package repository

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// CSVServiceTestSuite - Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type CSVServiceTestSuite struct {
	suite.Suite
	repository CSVService
}

// before each test
func (suite *CSVServiceTestSuite) SetupTest() {
	file, err := os.Create("./files/testfile.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	repository := New(file)
	suite.repository = *repository
}

// after each test
func (suite *CSVServiceTestSuite) TearDownSuite() {
	suite.repository.file.Close()
	os.Remove(suite.repository.file.Name())
}

//TestGetData_Positive - Test correct access to file
func (suite *CSVServiceTestSuite) TestGetData_Positive() {

	_, err := suite.repository.GetData()
	suite.NoError(err, "no error when getting data")
}

//TestGetData_Negative - Test no access to file
func (suite *CSVServiceTestSuite) TestGetData_Negative() {

	suite.repository.file.Close()
	_, err := suite.repository.GetData()
	suite.Error(err, "show error message")
}

//TestAccessFileSuite - Runs the test on the suite
func TestAccessFileSuite(t *testing.T) {
	suite.Run(t, new(CSVServiceTestSuite))
}
