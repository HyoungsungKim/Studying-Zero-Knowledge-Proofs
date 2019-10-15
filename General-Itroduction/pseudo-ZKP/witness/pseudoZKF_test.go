package witness

import (
	"testing"
)

func TestGetWitness(t *testing.t) {
	problem := []int{4, 11, 8, 1}
	assignment := []int{1, -1, 1, -1}
	witness, err := getWitness(problem, assignment)

	if err == nil {
		t.Logf("witness : %v", witness)
	} else {
		t.Errorf("%v", err)
	}
}
