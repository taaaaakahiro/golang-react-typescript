package main

import (
	"testing"
)

func TestPractice(t *testing.T) {
	post, err := testPractice(3, 6)
	if err != nil {
		t.Error(err)
	}
}