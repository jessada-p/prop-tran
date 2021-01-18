package auth_test

import (
	"testing"

	"github.com/pallat/todos/auth"
)

func TestAutnJWT(t *testing.T) {
	given, _ := auth.Token()

	if ok, err := auth.AuthKey(given); !ok {
		t.Errorf("token %v, validate %v , error %v \n", given, ok, err)
	}
}
