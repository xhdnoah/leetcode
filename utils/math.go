package utils

import "math"

const Mod int = 1e9 + 7

type comparable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Min[T comparable](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Max[T comparable](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Minimum(nums ...int) int {
	m := math.MaxInt32
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func Maximum(nums ...int) int {
	m := math.MinInt32
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
