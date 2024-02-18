package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDefaultDB(t *testing.T) {
	db := OpenDB()
	assert.NotNil(t, db)
}
