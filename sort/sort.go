package sort

//QuickSort 快排
func QuickSort(src []int, left, right int) {
	if len(src) == 0 {
		return
	}
	if left >= right {
		return
	}
	head := left
	tail := right
	key := src[(head+tail)/2]
	for head < tail {
		if src[head] < key {
			head++
		}
		if src[tail] > key {
			tail--
		}
		src[head], src[tail] = src[tail], src[head]
	}
	QuickSort(src, left, head-1)
	QuickSort(src, tail+1, right)
	return
}
