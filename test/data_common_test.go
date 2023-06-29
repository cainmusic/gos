package main

import (
	"strconv"
	"testing"

	"github.com/cainmusic/gos/data"
)

const f32 = 0xffffffff
const f64 = 0xffffffffffffffff

const (
	maxUint64 uint64 = uint64(f64)
	minUint64 uint64 = ^maxUint64

	maxInt64 int64 = int64(f64 >> 1)
	minInt64 int64 = ^maxInt64

	maxUint32 uint32 = uint32(f32)
	minUint32 uint32 = ^maxUint32

	maxInt32 int32 = int32(f32 >> 1)
	minInt32 int32 = ^maxInt32
)

func TestMaxMin(t *testing.T) {
	t.Run("Test Int Uint", testIntUint)
	t.Run("Test Int32 Uint32", testInt32Uint32)
	t.Run("Test Int64 Uint64", testInt64Uint64)
}

func testIntUint(t *testing.T) {
	if strconv.IntSize == 32 {
		if data.MinUint != uint(minUint32) ||
			data.MaxUint != uint(maxUint32) ||
			data.MinInt != int(minInt32) ||
			data.MaxInt != int(maxInt32) {
			t.Fatal("IntSize 32, Test Int Uint Wrong")
		}
	} else if strconv.IntSize == 64 {
		if data.MinUint != uint(minUint64) ||
			data.MaxUint != uint(maxUint64) ||
			data.MinInt != int(minInt64) ||
			data.MaxInt != int(maxInt64) {
			t.Fatal("IntSize 64, Test Int Uint Wrong")
		}
	} else {
		t.Fatal("Unknown Int Size")
	}
}

func testInt32Uint32(t *testing.T) {
	if data.MinUint32 != minUint32 ||
		data.MaxUint32 != maxUint32 ||
		data.MinInt32 != minInt32 ||
		data.MaxInt32 != maxInt32 {
		t.Fatal("Test Int32 Uint32 Wrong")
	}
}

func testInt64Uint64(t *testing.T) {
	if data.MinUint64 != minUint64 ||
		data.MaxUint64 != maxUint64 ||
		data.MinInt64 != minInt64 ||
		data.MaxInt64 != maxInt64 {
		t.Fatal("Test Int64 Uint64 Wrong")
	}
}

func TestIntSize(t *testing.T) {
	// 看了下源代码，尴尬的地方在于，strconv.IntSize和data.IntSize用的是同一个算法
	if strconv.IntSize != data.IntSize {
		t.Fatal("IntSize Wrong")
	}
}
