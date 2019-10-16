package main

import (
	"testing"
)

func TestGetWitness(t *testing.T) {
	// GO111MODULE=off go test -timeout 1500s -run TestGetWitness -v
	problem := []int{4, 11, 8, 1}
	assignment := []int{1, -1, 1, -1}
	witness, err := getWitness(problem, assignment)

	if err == nil {
		t.Logf("witness : %v", witness)
	} else {
		t.Errorf("%v", err)
	}
}
