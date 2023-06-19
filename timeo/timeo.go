package timeo

import (
	"time"
)

const (
	Second int64 = 1
	Minute int64 = Second * 60
	Hour   int64 = Minute * 60
	Day    int64 = Hour * 24
)

// 全局式

var offsetSecond int64 = 0

func SetOffset(os int64) {
	offsetSecond = os
}

func ClearOffset() {
	offsetSecond = 0
}

func SetTime(t time.Time) {
	SetTimeNTD(time.Now(), t)
}

func SetTimeNTD(now time.Time, destination time.Time) {
	SetOffset(int64(destination.Sub(now) / time.Second))
}

func GetOffset() int64 {
	return offsetSecond
}

func GetOffsetDuration() time.Duration {
	return time.Duration(GetOffset()) * time.Second
}

func GetOffsetDurationString() string {
	return GetOffsetDuration().String()
}

func Now() time.Time {
	return Unix(time.Now().Unix())
}

func Unix(ts int64) time.Time {
	return time.Unix(ts+GetOffset(), 0)
}

// 对象式

type Offset struct {
	offsetSecond int64
}

func NewOffset(os int64) *Offset {
	return &Offset{
		offsetSecond: os,
	}
}

func (pos *Offset) SetOffset(os int64) {
	pos.offsetSecond = os
}

func (pos *Offset) ClearOffset() {
	pos.offsetSecond = 0
}

func (pos *Offset) SetTime(t time.Time) {
	pos.SetTimeNTD(time.Now(), t)
}

func (pos *Offset) SetTimeNTD(now time.Time, destination time.Time) {
	pos.SetOffset(int64(destination.Sub(now) / time.Second))
}

func (pos *Offset) GetOffset() int64 {
	return pos.offsetSecond
}

func (pos *Offset) GetOffsetDuration() time.Duration {
	return time.Duration(pos.GetOffset()) * time.Second
}

func (pos *Offset) GetOffsetDurationString() string {
	return pos.GetOffsetDuration().String()
}

func (pos *Offset) Now() time.Time {
	return pos.Unix(time.Now().Unix())
}

func (pos *Offset) Unix(ts int64) time.Time {
	return time.Unix(ts+pos.GetOffset(), 0)
}
