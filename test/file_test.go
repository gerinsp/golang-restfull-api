package test

import (
	"belajar-golang-restfull-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectioon(t *testing.T) {
	conn, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, conn)

	cleanup()
}
