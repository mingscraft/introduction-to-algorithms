/*
Our first algorithm, insertion sort, solves the sorting problem introduced in chapter1:
Input: A sequence of n numbers <a1, a2, ...., an>.
Output: A permutation(reordering) <a1', a2', ....,an'> of the input sequence such that a1', <= a2' <= ...<=an'.
 */
package insertion_sort

type SortOrder int

const (
	Ascending SortOrder = iota + 1
	Descending
)

// InsertionSort sort array using insertion sort.
func InsertionSort(arr []int, sortOrder SortOrder) []int {

	j := 1
	for j < len(arr) {
		key := arr[j]

		// insert arr[j]  into the sorted sequence A[1..j-1].
		i := j - 1
		for i >= 0 && shouldMoveRight(arr[i], key, sortOrder) {
			arr[i + 1] = arr[i]
			i = i - 1
		}

		arr[i + 1] = key
		j ++
	}

	return arr
}

// shouldMoveRight should move i (j - 1) to right
func shouldMoveRight(a int, b int, sortOrder SortOrder) bool {
	if sortOrder == Ascending {
		return a > b
	}

	return b > a
}
