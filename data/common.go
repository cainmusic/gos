package data

// max and min
const (
	// uint
	MinUint uint = 0        // int的位数，所有位数为0
	MaxUint uint = ^MinUint // int的位数，所有位数为1

	// uint
	MaxInt int = int(MaxUint >> 1) // int的位数，第一位为0，其他为1
	MinInt int = ^MaxInt           // int的位数，第一位为1，其他为0

	// uint32
	MinUint32 uint32 = 0
	MaxUint32 uint32 = ^MinUint32

	// int32
	MaxInt32 int32 = int32(MaxUint32 >> 1)
	MinInt32 int32 = ^MaxInt32

	// uint64
	MinUint64 uint64 = 0
	MaxUint64 uint64 = ^MinUint64

	// int64
	MaxInt64 int64 = int64(MaxUint64 >> 1)
	MinInt64 int64 = ^MaxInt64
)

// 下面的两个常量慎用
const ConstMaxUint32 = 1<<32 - 1
const ConstMaxUint64 = 1<<64 - 1 // 暂时不清楚32位机器上下面的ConstMaxUint64会如何表现

const IntSize = 32 << (^uint(0) >> 63)
