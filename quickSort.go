package main

/*
+------------------------------------------+---------------------+
|    Time complexity                       |   Space Complexity  |
+-------------+---------------+------------+---------------------+
| Best        |  Average      |   Worst    |        Worst        |
+-------------+---------------+------------+---------------------+
| Ω(n log(n)) |  Θ(n log(n))  |   O(n^2)   |        O(log(n))    |
+-------------+---------------+------------+---------------------+
*/

func quickSort(array []int) {
	quickSortInternal(array, 0, len(array)-1)
}

func quickSortInternal(array []int, left int, right int) {
	if left < right {
		p := reorder(array, left, right)
		quickSortInternal(array, left, p-1)
		quickSortInternal(array, p, right)
	}
}

func reorder(array []int, left int, right int) int {
	p := array[left+(right-left)/2]
	i := left
	j := right

	for i <= j {
		for ; array[i] < p; i++ {
		}
		for ; array[j] > p; j-- {
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	return i
}
