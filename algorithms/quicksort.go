package algorithms

func QuickSort(a []int, n int) {
	quickSort(a, 0, n-1)
}

func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}

func partition(a []int, p, r int) int {
	pivot := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}

	a[i], a[r] = a[r], a[i]

	return i
}
