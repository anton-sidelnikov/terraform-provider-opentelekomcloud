package acceptance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CommonTestSuite struct {
	suite.Suite
	unos int
}

// SetupSuite this function executes before the test suite begins execution
func (suite *CommonTestSuite) SetupSuite() {
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite")
	suite.unos = 1
}

// TearDownSuite this function executes after all tests executed
func (suite *CommonTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// SetupTest this function executes before each test case
func (suite *CommonTestSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
	suite.unos = 1
}

// TearDownTest this function executes after each test case
func (suite *CommonTestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *CommonTestSuite) TestAddOne() {
	fmt.Println("From TestAddOne")
	suite.unos += 1
	suite.Equal(2, suite.unos)
}

func (suite *CommonTestSuite) TestSubtractOne() {
	fmt.Println("From TestSubtractOne")
	suite.unos -= 1
	suite.Equal(0, suite.unos)
}

func TestCommonTestSuite(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}
