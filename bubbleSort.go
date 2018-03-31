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

func bubblesort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > i; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
}
