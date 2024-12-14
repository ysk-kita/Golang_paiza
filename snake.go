package main

import (
	"bufio"
	"fmt"
)

var area [][]rune
var checkedArea [][]int
var mineswiper [][]int

func SnakeArea(sc *bufio.Scanner, tate int, yoko int) ([][]rune, [][]int) {

	var area [][]rune

	for i := 0; i < tate; i++ {
		//sc.Scan()
		//strs := sc.Text()
		strs := scanDummy(i)
		var areaYoko []rune
		var checkedAreaYoko []int
		var mineswiperYoko []int

		for _, val := range strs {
			areaYoko = append(areaYoko, val)
			checkedAreaYoko = append(checkedAreaYoko, 0)
			mineswiperYoko = append(mineswiperYoko, 0)
		}
		area = append(area, areaYoko)
		checkedArea = append(checkedArea, checkedAreaYoko)
		mineswiper = append(mineswiper, mineswiperYoko)
	}

	fmt.Println(tate, yoko)
	fmt.Println(area)

	return area, checkedArea
}

// /---
// / エリアチェック開始
// /---
func AreaCheck(tate int, yoko int) int {

	// 人のエリア
	var areaCnt int = 0

	for i := 0; i < tate; i++ {

		// 行に','が存在するかチェックする
		for j := 0; j < yoko; j++ {

			// runeが46(',')であるなら上下左右を探索しにいく。
			// また、ルーンとしてチェックしたエリアは探索対象に含めない
			if checkedArea[i][j] == 0 && area[i][j] == 46 {

				//

			}

		}
	}

	return areaCnt
}

func JyougeSayuCheck(i, j int) {

}

// 上のマスをチェックする
func UeCheck(i, j int) int {
	// 1行目の場合はチェック無し
	if i == 0 {
		return 0
	}

	// 上のマスが’,'なら再度上下左右チェックへ
	if area[i-1][j] == 46 {
		return 1
	}
	return 0
}

// 下のマスをチェックする
func SitaCheck(i, j, tate int) int {
	// 最終行目の場合はチェック無し
	if i+1 == tate {
		return 0
	}

	// 下のマスが’,'なら再度上下左右チェックへ
	if area[i+1][j] == 46 {
		return 1
	}
	return 0
}

// 右のマスをチェックする
func MigiCheck(i, j, yoko int) int {
	// 行の最後の場合はチェック無し
	if j+1 == yoko {
		return 0
	}

	// 右のマスが’,'なら再度上下左右チェックへ
	if area[i][j+1] == 46 {
		return 1
	}
	return 0
}

// 　左のマスをチェックする
func HidariCheck(i, j int) int {
	// 行の先頭の場合はチェック無し
	if j-1 == 0 {
		return 0
	}

	// 左のマスが’,'なら再度上下左右チェックへ
	if area[i][j-1] == 46 {
		return 1
	}
	return 0
}

func scanDummy(i int) string {
	strs := ""
	if i == 0 {
		strs = "######"
	}

	if i == 1 {
		strs = ".....#"
	}

	if i == 2 {
		strs = "##.###"
	}

	if i == 3 {
		strs = ".###.."
	}
	return strs
}
