package testsuite

import (
	"github.com/stretchr/testify/suite"
)

// GenericTestSuite is the base test suite for generic database tests.
type GenericTestSuite struct {
	suite.Suite

	Helper
}

// AfterTest is called after each test.
func (s *GenericTestSuite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.Require().NoError(err)
}

// BeforeTest is called before each test.
func (s *GenericTestSuite) BeforeTest(suiteName, testName string) {
	err := s.SetUp()
	s.Require().NoError(err)
}

