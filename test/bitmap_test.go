package main

import (
	"testing"

	"github.com/cainmusic/gos/data"
)

var bm1, bm2 *data.BitMap

func init() {
	bm1 = data.NewBitMap(200, 0)
	bm2 = data.NewBitMap(100, 0)
}

func TestBM1(t *testing.T) {
	t.Run("Test BitMap1 New", testBitMap1New)
	t.Run("Test BitMap1 SetGot", testBitMap1SetGet)
	t.Run("Test BitMap1 XAndPos", testBitMap1XAndPos)
	t.Run("Test BitMap1 SetSliece", testBitMap1SetSlice)
	t.Run("Test BitMap1 Output", testBitMap1Output)
	t.Run("Test BitMap1 Panic", testBitMap1Panic)
	t.Run("Test BitMap1 Set0", testBitMap1Set0)
}

func testBitMap1New(t *testing.T) {
	if bm1.L != 200 {
		t.Fatal("wrong bm.L")
	}
	if bm1.N != 0 {
		t.Fatal("wrong bm.N")
	}
	if len(bm1.M) != 4 {
		t.Fatal("wrong len(bm.M)")
	}
}

func testBitMap1SetGet(t *testing.T) {
	bm1.Set(50)
	if bm1.Got(50) != true {
		t.Fatal("wrong set")
	}
	if bm1.Got(100) != false {
		t.Fatal("wrong got")
	}
}

func testBitMap1XAndPos(t *testing.T) {
	if bm1.PosToX(bm1.XToPos(80)) != 80 {
		t.Fatal("wrong x and pos")
	}
}

func testBitMap1SetSlice(t *testing.T) {
	bm1.SetSlice([]int64{0, 50, 70, 100, 150, 200})
	if bm1.Got(0) != true ||
		bm1.Got(50) != true ||
		bm1.Got(70) != true ||
		bm1.Got(100) != true ||
		bm1.Got(150) != true ||
		bm1.Got(200) != true {
		t.Fatal("wrong set slice")
	}
}

func testBitMap1Output(t *testing.T) {
	r := bm1.Output()
	for i, v := range []int64{0, 50, 70, 100, 150, 200} {
		if r[i] != v {
			t.Fatal("wrong output")
		}
	}
}

func testBitMap1Panic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("recover from panic: %v", err)
		} else {
			t.Fatal("wrong, should panic")
		}
	}()
	bm1.Set(-1)
}

func testBitMap1Set0(t *testing.T) {
	bm1.Set0(50)
	if bm1.Got(50) != false {
		t.Fatal("wrong set0")
	}
}

func TestBM2(t *testing.T) {
	t.Run("Test BitMap2 Repeat", testBitMap2Repeat)
}

func testBitMap2Repeat(t *testing.T) {
	t.Log("this test case is actually a use case for storing multi times only")
	xs := []int64{0, 0, 0, 1, 1, 1, 1, 2, 5, 8, 10, 100, 100}
	repeatMap := map[int64]int{}
	for _, x := range xs {
		if bm2.Got(x) {
			if repeatMap[x] == 0 {
				repeatMap[x] = 2
			} else {
				repeatMap[x]++
			}
		} else {
			bm2.Set(x)
		}
	}
	if repeatMap[0] != 3 ||
		repeatMap[1] != 4 ||
		repeatMap[100] != 2 {
		t.Fatal("wrong repeat check")
	}
	t.Log(repeatMap)
}
