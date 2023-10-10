package util

import "github.com/spf13/cast"

func SliceElementToString[T any](s []T) []string {
	res := make([]string, 0, len(s))
	for _, v := range s {
		res = append(res, cast.ToString(v))
	}
	return res
}
