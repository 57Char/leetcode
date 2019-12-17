package algorithms

func binarySearch(a []int, n, v int) int {
	if n == 0 {
		return -1
	}

	low, high := 0, n-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == v {
			return mid
		}
		if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}