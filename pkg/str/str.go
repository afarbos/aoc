package str

func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Permutations call f on all possible combination of the set.
func Permutations(set map[string]struct{}, f func([]string)) {
	var (
		list = make([]string, len(set))
		i    = 0
	)

	for v := range set {
		list[i] = v
		i++
	}

	perm(list, f, 0)
}
