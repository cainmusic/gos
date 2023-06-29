package main

import (
	"testing"
	"time"

	"github.com/cainmusic/gos/timeo"
)

var day int64 = 86400

// global

func TestGlobal(t *testing.T) {
	t.Run("Test Global SetGet", testSetGet)
	t.Run("Test Global SetTime", testSetTime)
	t.Run("Test Global GetDurationString", testGetDurationString)
	t.Run("Test Global Now", testNow)
	t.Run("Test Global Clear", testClear)
}

func testSetGet(t *testing.T) {
	timeo.SetOffset(day)
	if timeo.GetOffset() != day {
		t.Fatal("Global SetGet Wrong")
	}
}

func testSetTime(t *testing.T) {
	now := time.Now()
	timeo.SetTime(now)
	if timeo.GetOffset() != 0 {
		t.Fatal("Global SetTime Wrong")
	}
}

func testGetDurationString(t *testing.T) {
	timeo.SetOffset(day)
	if timeo.GetOffsetDurationString() != "24h0m0s" {
		t.Fatal("Global GetOffsetDurationString Wrong")
	}
}

func testNow(t *testing.T) {
	timeoNow := timeo.Now()
	timeNow := time.Now()
	if timeoNow.Unix() != timeNow.Unix()+day {
		t.Fatal("Global Now Wrong")
	}
}

func testClear(t *testing.T) {
	timeo.ClearOffset()
	if timeo.GetOffset() != 0 {
		t.Fatal("Global Clear Wrong")
	}
}

// object

var pos = timeo.NewOffset(0)

func TestObject(t *testing.T) {
	t.Run("Test Object New", testObjectNew)
	t.Run("Test Object SetGet", testObjectSetGet)
	t.Run("Test Object SetTime", testObjectSetTime)
	t.Run("Test Object GetDurationString", testObjectGetDurationString)
	t.Run("Test Object Now", testObjectNow)
	t.Run("Test Object Clear", testObjectClear)
}

func testObjectNew(t *testing.T) {
	if pos.GetOffset() != 0 {
		t.Fatal("Object SetGet Wrong")
	}
}

func testObjectSetGet(t *testing.T) {
	pos.SetOffset(day)
	if pos.GetOffset() != day {
		t.Fatal("Object SetGet Wrong")
	}
}

func testObjectSetTime(t *testing.T) {
	now := time.Now()
	pos.SetTime(now)
	if pos.GetOffset() != 0 {
		t.Fatal("Object SetTime Wrong")
	}
}

func testObjectGetDurationString(t *testing.T) {
	pos.SetOffset(day)
	if pos.GetOffsetDurationString() != "24h0m0s" {
		t.Fatal("Object GetOffsetDurationString Wrong")
	}
}

func testObjectNow(t *testing.T) {
	posNow := pos.Now()
	timeNow := time.Now()
	if posNow.Unix() != timeNow.Unix()+day {
		t.Fatal("Object Now Wrong")
	}
}

func testObjectClear(t *testing.T) {
	pos.ClearOffset()
	if pos.GetOffset() != 0 {
		t.Fatal("Object Clear Wrong")
	}
}
