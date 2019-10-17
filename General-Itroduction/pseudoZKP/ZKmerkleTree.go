package main

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"reflect"
)

type ZKMerkleTree struct {
	data []big.Int
	tree [][32]byte
}

func (zk ZKMerkleTree) init(data []int) {
	for i := range data {
		zk.data = append(zk.data, *big.NewInt(int64(data[i])))
	}
	nextPowOf2 := int(math.Pow(2, math.Ceil(math.Log2(float64(len(data))))))

	for i := 0; i < (nextPowOf2 - len(data)); i++ {
		zk.data = append(zk.data, *big.NewInt(0))
	}

	var randList []big.Int
	for i := 0; i < len(data); i++ {
		bigInt := big.NewInt(1<<32 - 1)
		bigRand, _ := rand.Int(rand.Reader, bigInt)

		randList = append(randList, *bigRand)
	}

	for i := 0; i < len(zk.data); i++ {
		zk.data = append(zk.data, zk.data[i])
		zk.data = append(zk.data, randList[i])
	}

	var leftPart [][32]byte
	var rightPart [][32]byte

	for index := range zk.data {
		leftPart[index][0] = '0'
	}

	for index, val := range zk.data[:] {
		var byteSlice []byte
		byteSlice = append(byteSlice, []byte(val))
		rightPart[index] = hashString(byteSlice)
	}

	zk.tree = append(leftPart, rightPart...)

	for i := len(zk.data) - 1; i >= 0; i-- {
		hash := hashString(append(zk.tree[i*2][:], zk.tree[i*2+1][:]...))
		zk.tree[i] = hash
	}
}

func (zk ZKMerkleTree) getRoot() [32]byte {
	return zk.tree[1]
}

func (zk ZKMerkleTree) getValAndPath(id int) ([]byte, [][32]byte) {
	id = id * 2
	value := zk.data[id]
	var authPath [][32]byte

	id = id + len(zk.data)

	for {
		if id <= 1 {
			break
		}
		authPath = append(authPath, zk.tree[id^1])
		id = id / 2
	}

	return value, authPath
}

func VerifyZKMerklePath(dataSize, valueId int, root, value []byte, path [][32]byte) (bool, error) {
	cur := hashString(value)
	treeNodeID := valueId*2 + int(math.Pow(2, math.Ceil(math.Log2(float64(dataSize*2)))))

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
