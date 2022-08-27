package c2

type Ordered interface {
	~float64 | ~int | ~string
}

func quicksort[T Ordered](data []T, left, right int) {
	if left < right {
		var pivot = patition(data, left, right)

		quicksort(data, left, pivot)
		quicksort(data, pivot, right)
	}
}


func patition[T Ordered](data []T, left, right int) int {
	pivot := data[left]
	var i = left
	var j = right

	for i < j {
		for data[i] <= pivot && i < right {
			i++
		}

		for data[j] >= pivot && j > left {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}

	}

	data[left] = data[j]
	data[j] = pivot
	return j
}