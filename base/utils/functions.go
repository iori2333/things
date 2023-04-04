package utils

type Cloneable[T any] interface {
	Clone() T
}

func CloneMap[K comparable, T interface{ Clone() T }](m map[K]T) map[K]T {
	ret := make(map[K]T)
	for k, v := range m {
		ret[k] = v.Clone()
	}
	return ret
}
