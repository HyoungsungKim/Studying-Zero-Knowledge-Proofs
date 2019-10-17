package main

import (
	"crypto/sha256"
	"errors"
	"math"
	"math/big"
	"reflect"
)

func hashString(s big.Int) [32]byte {
	return sha256.Sum256(s)
}

type MerkleTree struct {
	data []byte
	tree [][32]byte
}

func (m MerkleTree) init(data []byte) {
	m.data = data
	nextPowOf2 := int(math.Pow(2, math.Ceil(math.Log2(float64(len(data))))))

	for i := 0; i < (nextPowOf2 - len(data)); i++ {
		m.data = append(m.data, '0')
	}

	var leftPart [][32]byte
	var rightPart [][32]byte

	for index := range m.data {
		leftPart[index][0] = '0'
	}

	for index, val := range m.data {
		var byteSlice []byte
		byteSlice = append(byteSlice, val)
		rightPart[index] = hashString(byteSlice)
	}

	m.tree = append(leftPart, rightPart...)

	for i := len(m.data) - 1; i >= 0; i-- {
		hash := hashString(append(m.tree[i*2][:], m.tree[i*2+1][:]...))
		m.tree[i] = hash
	}
}

func (m MerkleTree) getRoot() [32]byte {
	return m.tree[1]
}

func (m MerkleTree) getValAndPath(id int) (byte, [][32]byte) {
	value := m.data[id]
	var authPath [][32]byte

	id = id + len(m.data)

	for {
		if id <= 1 {
			break
		}
		authPath = append(authPath, m.tree[id^1])
		id = id / 2
	}

	return value, authPath
}

func VerifyMerklePath(dataSize, valueId int, root, value []byte, path [][32]byte) (bool, error) {
	cur := hashString(value)
	treeNodeID := valueId + int(math.Pow(2, math.Ceil(math.Log2(float64(dataSize)))))

	for _, sibling := range path {
		if treeNodeID <= 0 {
			return false, errors.New("Wrong tree node id")
		}

		if treeNodeID%2 == 0 {
			cur = hashString(append(cur[:], sibling[:]...))
		} else {
			cur = hashString(append(sibling[:], cur[:]...))
		}
		treeNodeID = treeNodeID / 2
	}

	if treeNodeID != 1 {
		return false, errors.New("Wrong tree node id")
	}

	return reflect.DeepEqual(root[:], cur[:]), nil
}
