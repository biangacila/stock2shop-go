package util

import (
	"testing"
)

func TestRepository_New(t *testing.T) {
	hub := Repository{}
	hub.Keyspace = DB_NAME
	hub.HasFile = true
	hub.New()
}
