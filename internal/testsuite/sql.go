package testsuite

import (
	"github.com/stretchr/testify/suite"
)

// SQLTestSuite is the base test suite for SQL tests.
type SQLTestSuite struct {
	suite.Suite

	Helper
}

// AfterTest is called after each test.
func (s *SQLTestSuite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.Require().NoError(err)
}

// BeforeTest is called before each test.
func (s *SQLTestSuite) BeforeTest(suiteName, testName string) {
	err := s.SetUp()
	s.Require().NoError(err)
}

