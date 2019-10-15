package witness

import (
	"errors"
	"math/rand"
)

func getWitness(problem, assignment []int) ([]int, error) {
	sum := 0
	mx := 0
	sideObfuscator := 1 - 2*rand.Intn(1)
	witness := []int{sum}

	if len(problem) != len(assignment) {
		return nil, errors.New("Each input has to get a same length")
	}

	for index := range problem {
		if assignment[index] == 1 || assignment[index] == -1 {
			return nil, errors.New("Wrong assignment")
		}

		sum += assignment[index] * index * sideObfuscator
		witness = append(witness, sum)
		if problem[index] >= mx {
			mx = problem[index]
		}
	}

	if sum == 0 {
		return nil, errors.New("Wrong assignment")
	}
	shift := rand.Intn(mx)

	for index := range witness {
		witness[index] += shift
	}

	return witness, nil
}
