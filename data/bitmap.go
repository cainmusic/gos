package data

import (
	"fmt"
)

// 2^6=64, int64用Bitn=6
// 2^5=32, int32用Bitn=5
// 2^4=16, int16用Bitn=4
// 2^3=8, int8用Bitn=3

const (
	Bitn   = 6
	BitLen = 1 << Bitn
	Bits   = BitLen - 1
)

var BitAnd, BitOr [64]uint64

func init() {
	for i := 0; i < BitLen; i++ {
		BitAnd[i] = 1 << i
		BitOr[i] = ^BitAnd[i]
	}
}

func Test() {
	fmt.Println("BitAnd", BitAnd)
	fmt.Println("BitOr", BitOr)
}

type BitMap struct {
	L int64
	M []uint64
	N int64
	R map[int64]int64
}

func NewBitMap(length int64, startLine int64) *BitMap {
	tl := length/BitLen + 1
	return &BitMap{
		L: length,
		M: make([]uint64, tl),
		N: startLine,
	}
}

/*
  应该提前做好x的合法性验证，否则会panic。
  x的值，应处于[bm.N, bm.N+bm.L]。
*/
func (bm *BitMap) XToPos(x int64) (i, j int64) {
	x -= bm.N
	if x < 0 || x > bm.L {
		panic("wrong x, should between [N, N+L]")
	}
	return x >> Bitn, x & Bits
}

func (bm *BitMap) PosToX(i, j int64) int64 {
	return (i << Bitn) + j + bm.N
}

func (bm *BitMap) SetByPos(i, j int64) {
	bm.M[i] |= BitAnd[j]
}

func (bm *BitMap) Set(x int64) {
	bm.SetByPos(bm.XToPos(x))
}

func (bm *BitMap) SetSlice(xs []int64) {
	for _, x := range xs {
		bm.Set(x)
	}
}

func (bm *BitMap) Set0(x int64) {
	i, j := bm.XToPos(x)
	bm.M[i] &= BitOr[j]
}

func (bm *BitMap) GotByPos(i, j int64) bool {
	return bm.M[i]&BitAnd[j] > 0
}

func (bm *BitMap) Got(x int64) bool {
	return bm.GotByPos(bm.XToPos(x))
}

func (bm *BitMap) Output() []int64 {
	r := []int64{}
	for i, l := range bm.M {
		for j := 0; j < BitLen; j++ {
			if l&BitAnd[j] > 0 {
				r = append(r, bm.PosToX(int64(i), int64(j)))
			}
		}
	}
	return r
}

func (bm *BitMap) Sort() []int64 {
	return bm.Output()
}
