package main

/*
+--------------------------+---------------------+
|    Time complexity       |   Space Complexity  |
+------+----------+--------+---------------------+
| Best |  Average | Worst  |        Worst        |
+------+----------+--------+---------------------+
| Ω(n) |  Θ(n^2)  | O(n^2) |        O(1)         |
+------+----------+--------+---------------------+
*/

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		k := array[i]
		j := i - 1
		for j >= 0 && array[j] > k {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = k
	}
}
