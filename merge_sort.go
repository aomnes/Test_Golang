package main

import "fmt"

// MergeSort recursion
func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merge left right
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

/*
input: Stdin
n_element
a1, a2, a3, ..., an
*/
func main() {
	var valeur int
	var tmp int
	fmt.Scanf("%d", &valeur)
	s := make([]int, valeur)
	for i := int(0); i < valeur; i++ {
		fmt.Scanf("%d", &tmp)
		s[i] = tmp
	}
	s = MergeSort(s)
	fmt.Println(s)
}
