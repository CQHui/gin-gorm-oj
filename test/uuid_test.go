package test

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	u := uuid.New()
	println(u.String())
}
