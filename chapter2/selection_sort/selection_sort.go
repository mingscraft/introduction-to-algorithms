package selection_sort

// SelectionSort sort int slice with selection sort
// Consider sorting n numbers stored in array A by first finding the smallest element of
// A, and exchange it with the element in A[1]. Then find the second smallest element of A,
// And exchange it with A[2]. Continue in this manner for the first n - 1 elements of A.
func SelectionSort(a []int) []int {
	// i start with 1
	// loop invariant: A[i..n], smallest element will be in position of A[i]
	// maintain: i ++
    // termination: i = n -1, we only need to run n - 1 element the 6th element would always be biggest number.
    // the best-case and worst case running time of selection sort are the same: n + (n - 1) + (n - 2) ...
   	for i := 0; i < len(a); i++ {
		var smallest = i
   		for j, val := range a[i:] {
   			if val < a[smallest] {
   				smallest = j + i
			}
		}
   		a[i], a[smallest] = a[smallest], a[i]
	}
   	return a
}
