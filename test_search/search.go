package test_search

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 06/03/2023;
 * Time : 1:48;
**/

// Find returns the smallest index i at which x <= a[i].
// If there is no such index, it returns len(a).
// The slice must be sorted in ascending order.
//{[]int{1, 2, 3, 3}, 3, 3},
func Find(a []int, x int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		if x <= a[0] {
			return 0
		}
		return 1
	}
	mid := len(a) / 2
	if x <= a[mid-1] {
		return Find(a[:mid], x)
	}
	return mid + Find(a[mid:], x)
}
