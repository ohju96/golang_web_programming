package memberships

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestServiceTestSuit(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *ServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
}

func (s *ServiceTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("BeforeTest")
}

func (s *ServiceTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest")
}

func (s *ServiceTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

func (s *ServiceTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *ServiceTestSuite) TestCreate() {
	// TODO
}

func (s *ServiceTestSuite) TestUpdate() {
	// TODO
}
