package util

func If[T any](condition bool, onTrue T, onFalse T) T {
	if condition {
		return onTrue
	}
	return onFalse
}
