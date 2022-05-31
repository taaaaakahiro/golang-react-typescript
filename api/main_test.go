package main

import (
	"testing"
	"api/practice"
	"fmt"
)

func TestPractice(t *testing.T) {
	got := practice.MathTest(5, 7)
    if got != 35 {
        fmt.Println(got)
    }
}