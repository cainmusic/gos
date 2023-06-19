package main

import (
	"fmt"
	"time"

	"github.com/cainmusic/gos/timeo"
)

func main() {
	test1()
	fmt.Println("===============")
	test2()
}

func test1() {
	now := timeo.Now()
	fmt.Println(now)

	day := int64(86400)
	timeo.SetOffset(day)
	now = timeo.Now()
	fmt.Println(now)

	now = timeo.Unix(now.Unix())
	fmt.Println(now)

	timeo.SetTime(time.Unix(86400, 0))
	now = timeo.Now()
	fmt.Println(now)

	timeo.SetTimeNTD(time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*24*3))
	now = timeo.Now()
	fmt.Println(now)
	fmt.Println(timeo.GetOffsetDuration())

	timeo.ClearOffset()
	now = timeo.Now()
	fmt.Println(now)
	fmt.Println(timeo.GetOffsetDuration())

	timeo.SetOffset(timeo.Day*3 + timeo.Hour*12)
	now = timeo.Now()
	fmt.Println(now)
	fmt.Println(timeo.GetOffsetDuration())
	fmt.Println(timeo.GetOffsetDurationString())
	fmt.Println(timeo.GetOffset())

	timeo.SetTime(now.Add(-timeo.GetOffsetDuration()))
	now = timeo.Now()
	fmt.Println(now)
	fmt.Println(timeo.GetOffsetDurationString())
}

func test2() {
	os1 := timeo.NewOffset(0)
	t1 := os1.Now()
	fmt.Println(t1)

	day := int64(86400)
	os1.SetOffset(day)
	t1 = os1.Now()
	fmt.Println(t1)

	t1 = os1.Unix(t1.Unix())
	fmt.Println(t1)

	os1.SetTime(time.Unix(86400, 0))
	t1 = os1.Now()
	fmt.Println(t1)

	os1.SetTimeNTD(time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*24*3))
	t1 = os1.Now()
	fmt.Println(t1)
	fmt.Println(os1.GetOffsetDuration())

	os1.ClearOffset()
	t1 = os1.Now()
	fmt.Println(t1)
	fmt.Println(os1.GetOffsetDuration())

	os1.SetOffset(timeo.Day*3 + timeo.Hour*12)
	t1 = os1.Now()
	fmt.Println(t1)
	fmt.Println(os1.GetOffsetDuration())
	fmt.Println(os1.GetOffsetDurationString())
	fmt.Println(os1.GetOffset())

	os1.SetTime(t1.Add(-os1.GetOffsetDuration()))
	t1 = os1.Now()
	fmt.Println(t1)
	fmt.Println(os1.GetOffsetDurationString())
}
