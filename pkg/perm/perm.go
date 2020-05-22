package perm

func str(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}

	str(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		str(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Strings call f on all possible combination of the set.
func Strings(set map[string]struct{}, f func([]string)) {
	var (
		list = make([]string, len(set))
		i    = 0
	)

	for v := range set {
		list[i] = v
		i++
	}

	str(list, f, 0)
}

func ints(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}

	ints(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		ints(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Ints call f on all possible combination of the set.
func Ints(set map[int]struct{}, f func([]int)) {
	var (
		list = make([]int, len(set))
		i    = 0
	)

	for v := range set {
		list[i] = v
		i++
	}

	ints(list, f, 0)
}
