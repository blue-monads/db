package testsuite

import (
	db "github.com/upper/db/v4"
)

// Helper is the interface that test helpers must implement.
type Helper interface {
	Session() db.Session
	Adapter() string
	SetUp() error
	TearDown() error
}

