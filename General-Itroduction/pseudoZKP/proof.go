package main

func getProof(problem, assignment []int, numQueries int) {
	var proof []byte
	randomnessSeed := problem[:]

	for i := 0; i < numQueries; i++ {
		zkmerkle := ZKMerkleTree{}
		witness, _ := getWitness(problem, assignment)

		tree = zkmerkle.init(witness)
	}
}
