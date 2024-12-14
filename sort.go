package main

import (
	"fmt"
)

/*
バブルソート
*/
func BubbleSort() {
	il := []int{4, 5, 32, 2, 1}
	sl := []string{"A", "B", "C", "D", "E"}
	/*
		sl := []string{"A","B","C","D","E"}

		sort.Ints(il)
		fmt.Println(il)

		sort.Strings(sl)
		fmt.Println(sl)

	*/

	for i := 0; i < len(il)-1; i++ {
		for j := 0; j < len(il)-1-i; j++ {
			if il[j] < il[j+1] {
				il[j], il[j+1] = il[j+1], il[j]
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}

	fmt.Print(il)
	fmt.Print(sl)
}
