package repository

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type CSVServiceTestSuite struct {
	suite.Suite
	repository CSVService
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *CSVServiceTestSuite) SetupTest() {
	file, err := os.Create("./files/testfile.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	repository := New(file)
	suite.repository = *repository
}

func (suite *CSVServiceTestSuite) TearDownSuite() {
	suite.repository.file.Close()
	os.Remove(suite.repository.file.Name())
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *CSVServiceTestSuite) TestGetData_Positive() {

	_, err := suite.repository.GetData()
	suite.NoError(err, "no error when getting data")
}

func (suite *CSVServiceTestSuite) TestGetData_Negative() {

	suite.repository.file.Close()
	_, err := suite.repository.GetData()
	suite.Error(err, "show error message")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(CSVServiceTestSuite))
}
