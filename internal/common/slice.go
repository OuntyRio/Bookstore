package common

func SliceMap[T1, T2 any](f func(*T1) T2, items []T1) []T2 {
	if len(items) == 0 {
		return []T2{}
	}

	s := make([]T2, len(items))
	for i := range items {
		s[i] = f(&items[i])
	}
	return s
}

func SliceMapPtr[T1, T2 any](f func(*T1) *T2, items []T1) []*T2 {
	if len(items) == 0 {
		return []*T2{}
	}

	s := make([]*T2, len(items))
	for i := range items {
		s[i] = f(&items[i])
	}
	return s
}
