package vector

import (
	"encoding/binary"
	"fmt"
	"hash/maphash"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2025/sliceutils"
)

var seed = maphash.MakeSeed()

type Vector struct {
	elements []int
}

func New(elements []int) Vector {
	return Vector{sliceutils.ShallowCopy(elements)}
}

func Fill(size int, val int) Vector {
	elements := make([]int, size)
	for idx := range size {
		elements[idx] = val
	}
	return New(elements)
}

func (v Vector) Size() int {
	return len(v.elements)
}

func (v Vector) Clone() Vector {
	elements := sliceutils.ShallowCopy(v.elements)
	return New(elements)
}

func (v Vector) At(idx int) (int, error) {
	if idx < 0 || idx >= v.Size() {
		return 0, indexError
	}

	return v.elements[idx], nil
}

func (v Vector) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for idx, val := range v.elements {
		if idx == v.Size()-1 {
			builder.WriteString(strconv.Itoa(val))
		} else {
			fmt.Fprintf(&builder, "%d,", val)
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (v Vector) Hash() uint64 {
	var hash maphash.Hash
	hash.SetSeed(seed)
	var buffer [8]byte
	for _, val := range v.elements {
		binary.LittleEndian.PutUint64(buffer[:], uint64(val))
		hash.Write(buffer[:])
	}
	return hash.Sum64()
}

func (v Vector) IsAll(cond func(int) bool) bool {
	for _, val := range v.elements {
		if !cond(val) {
			return false
		}
	}
	return true
}
