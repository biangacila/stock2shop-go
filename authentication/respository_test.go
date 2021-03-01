package authentication

import (
	"stock2shop-go/io"
	"testing"
)

func TestRepository_New(t *testing.T) {
	hub := Repository{}
	hub.Keyspace = io.DB_NAME
	hub.HasFile = true
	hub.New()
}
