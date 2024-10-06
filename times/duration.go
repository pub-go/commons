package times

import (
	"time"

	"code.gopub.tech/commons/nums"
)

func Nanoseconds[T nums.Number](ns T) time.Duration {
	return time.Duration(int64(ns)) * time.Nanosecond
}

func Microseconds[T nums.Number](us T) time.Duration {
	return time.Duration(int64(us)) * time.Microsecond
}

func Milliseconds[T nums.Number](ms T) time.Duration {
	return time.Duration(int64(ms)) * time.Millisecond
}

func Seconds[T nums.Number](s T) time.Duration {
	return time.Duration(int64(s)) * time.Second
}

func Minutes[T nums.Number](min T) time.Duration {
	return time.Duration(int64(min)) * time.Minute
}

func Hours[T nums.Number](hour T) time.Duration {
	return time.Duration(int64(hour)) * time.Hour
}

func Days[T nums.Number](day T) time.Duration {
	return time.Duration(int64(day)) * 24 * time.Hour
}
