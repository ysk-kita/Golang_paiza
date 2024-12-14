package main

import "fmt"

// 五目並べの勝ち判定
func CheckMain() {
	var str string
	var arr []string

	var yokoRes string
	var yokoWinFlg bool = false

	for i := 0; i < 5; i++ {
		fmt.Scan(&str)
		arr = append(arr, str)

		// 横行チェックを行う
		yokoRes = YokoCheck(str)
		if yokoRes != "D" {
			yokoWinFlg = true
			break
		}
	}
	// 横でかってれば終了
	if yokoWinFlg {
		fmt.Println(yokoRes)
		return
	}

	// 縦でかってれば終了
	tateRes := TateCheck(arr)
	if tateRes != "D" {
		fmt.Println(tateRes)
		return
	}

	// 最後は斜め
	nanameRes := NanameCheck(arr)
	if nanameRes != "D" {
		fmt.Println(nanameRes)
		return
	}

	// 勝ちなしならドロー表示
	fmt.Println("D")
}

// 左右の文字チェック実施
func YokoSameCheck(str string, i, W int) bool {
	var leftSame, rightSame bool
	// 左側確認 ただし座標が左端('0') の場合は強制的にTrue
	if i == 0 {
		leftSame = true
	} else {
		leftSame = str[i] == str[i-1]
	}

	// 右側確認　ただし座標が右端(W)の場合は強制的にtrue
	if i == W-1 {
		rightSame = true
	} else {
		rightSame = str[i] == str[i+1]
	}

	return leftSame && rightSame

}

// 上下の文字が指定した文字と一緒か判定
func TateSameAsPerfixStrCheck(strs []string, i, j, H int, rn string) bool {

	var upperSame, bottomSame bool
	// 上確認、ただし1行目であれば強制でTrue
	if i == 0 {
		upperSame = true
	} else {
		upperSame = rn == string(strs[i-1][j])
	}

	// 下側確認　ただし座標が右端(W)の場合は強制的にtrue
	if i == H-1 {
		bottomSame = true
	} else {
		bottomSame = rn == string(strs[i+1][j])
	}
	return upperSame && bottomSame
}

// 左右の文字が指定した文字と一緒か判定
func YokoSameAsPerfixStrCheck(str string, i, W int, rn string) bool {

	var leftSame, rightSame bool
	// 左側確認 ただし座標が左端('0') の場合は強制的にTrue
	if i == 0 {
		leftSame = true
	} else {
		leftSame = rn == string(str[i-1])
	}

	// 右側確認　ただし座標が右端(W)の場合は強制的にtrue
	if i == W-1 {
		rightSame = true
	} else {
		rightSame = rn == string(str[i+1])
	}
	return leftSame && rightSame
}

func NanameCheck(arr []string) string {
	// 左上⇒右下ルート
	var flg bool = true
	checkStr := arr[0][0]
	for i := 1; i < 5; i++ {
		// 不一致なら最後まで見る必要がないので終了
		if checkStr != arr[i][i] {
			flg = false
			break
		}
	}
	if flg {
		return string(checkStr)
	}

	// 不一致なら 左下⇒左上ルートへ
	flg2 := true
	checkStr = arr[4][0]
	for i := 1; i < 5; i++ {
		// 不一致なら最後まで見る必要がないので終了
		if checkStr != arr[4-i][i] {
			flg2 = false
			break
		}
	}
	if flg2 {
		return string(checkStr)
	}
	// 不一致2回なら引き分け
	return "D"
}

// 縦1列が全て同じ文字か判定
// 同じならその文字、そうでなければ"D"を返す
func TateCheck(arr []string) string {

	var res string = "D"

	for i := 0; i < 5; i++ {
		checkStr := arr[0][i]
		var isBreak bool = false

		for j := 1; j < 5; j++ {

			if checkStr != arr[j][i] {
				isBreak = true
			}
		}
		// 縦探索でbreakしない = 縦が全て一致なので答えを保存
		if !isBreak && arr[0][i] != '.' {
			res = string(arr[0][i])
			break
		}
	}
	return res
}

// 横1行が全て同じ文字か判定
// 同じならその文字、そうでなければDを返す
func YokoCheck(str string) string {
	var flg bool = true

	if str[0] == '.' {
		return "D"
	}

	for i := 1; i < len(str); i++ {
		if str[0] != str[i] {
			flg = false
			break
		}
	}
	if flg {
		return str[0:1]
	} else {
		return "D"
	}
}
