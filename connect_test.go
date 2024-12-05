package redditmongo

import (
	"testing"
)

func TestNew(t *testing.T) {
	err := ConnectUsingEnv()

	if err != nil {
		t.Error(err)
	}
}
