package p

func Map[T, R any](a []T, f func(T) R) []R {
	var r []R
	for _, e := range a {
		r = append(r, f(e))
	}
	return r
}

func Filter[T any](a []T, p func(T) bool) []T {
	var r []T
	for _, e := range a {
		if p(e) {
			r = append(r, e)
		}
	}
	return r
}
