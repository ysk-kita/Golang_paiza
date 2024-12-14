package main

import (
	"fmt"
	"slices"
)

func main() {

	//fmt.Println(strings.ToUpper(s[:1]) + s[1:])
	//var s string = "aaaAAAaa"
	//	fmt.Println(strings.Count("rrrrrrrrrr", "rr"))
	// var arr []string = []string{"HND", "NRT", "KIX", "NGO", "NGO"}
	// fmt.Println(slices.Contains(arr, "NRT"))
	// BubbleSort()

	// var mp map[string]int = map[string]int{"key": 2}

	// str := "key2"
	// mp[str] = 1
	// fmt.Println(mp)

	// if val, result := mp["key"]; result {
	// 	fmt.Println(result, val)
	// } else {
	// 	fmt.Println(result, val)
	// }

	// var s []int = []int{1, 21, 3, 4, 5}
	// sort.Sort(sort.Reverse(sort.IntSlice(s)))
	// fmt.Println(s)

	// str := "GO"
	// Print(str)

	// Print(str[0])
	// Print(str[1])

	// start := str[0]
	// end := str[1]

	// for i := start; i <= end; i++ {

	// 	fmt.Println(string(i))
	// }

	// 	var H, W, N int
	// 	fmt.Scanf("%d %d %d", &H, &W, &N)
	// 	var strs []string
	// 	for i := 0; i < H; i++ {
	// 		var str string
	// 		fmt.Scan(&str)

	// 		strs = append(strs, str)
	// 	}

	// 	for i := 0; i < N; i++ {
	// 		var x, y int
	// 		fmt.Scanf("%d %d", &x, &y)

	// 		fmt.Println(strs[x][y])
	// 	}
	// }

	// //fmt.Println(UzumakiIdouTokeiMawari(38, 47, 27))

	// var H, W, y, x, n int = 7, 3, 2, 1, 5
	// //var leftOrRight, muki string
	// var leftOrRight string = "L"

	// ary := []string{"..#", "...", "...", "...", "..#", ".#.", "##."}
	// fmt.Println("start pos:", x, y)
	// nowX, nowY, nowMuki := x, y, "N"
	// for i := 0; i < n; i++ {

	// 	// スピンできるなら移動を実施
	// 	nowX, nowY, nowMuki = IsMovableToSpin2(ary, nowX, nowY, H, W, leftOrRight, nowMuki)

	// 	if nowMuki == "Stop" {
	// 		fmt.Println(nowMuki)
	// 		break
	// 	} else {
	// 		fmt.Println("NowPos:", nowY, nowX, nowMuki)
	// 	}

	// }
	//SnakeSpinOnTime2()

	hs := []Test{
		{1, 2},
		{3, 4},
	}
	fmt.Println(hs)

	fmt.Println(slices.Contains(hs, Test{1, 2}))
	fmt.Println(slices.Contains(hs, Test{1, 3}))

}

type Test struct {
	x, y int
}

func Print(val interface{}) {
	fmt.Println(val)
}
