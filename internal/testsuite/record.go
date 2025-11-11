package testsuite

import (
	"github.com/stretchr/testify/suite"
)

// RecordTestSuite is the base test suite for record tests.
type RecordTestSuite struct {
	suite.Suite

	Helper
}

// AfterTest is called after each test.
func (s *RecordTestSuite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.Require().NoError(err)
}

// BeforeTest is called before each test.
func (s *RecordTestSuite) BeforeTest(suiteName, testName string) {
	err := s.SetUp()
	s.Require().NoError(err)
}

