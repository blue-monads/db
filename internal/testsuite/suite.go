package testsuite

import (
	"github.com/stretchr/testify/suite"
)

// Suite is a basic test suite with a Helper.
type Suite struct {
	suite.Suite

	Helper
}

// AfterTest is called after each test.
func (s *Suite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.Require().NoError(err)
}

// BeforeTest is called before each test.
func (s *Suite) BeforeTest(suiteName, testName string) {
	err := s.SetUp()
	s.Require().NoError(err)
}
