package main

/**
の式を求める
・ a_1 = 1
・ a_2 = 1
・ a_n = a_{n-2} + a_{n-1} (n ≧ 3)
*/
func Fibonati(n int) int {

	// n = 1, 2 の場合は1で解が固定
	if n < 3 {
		return 1
	}

	// n が 3以上ならa_{n-2} + a_{n-1} を算出

	// a_n-2, a_n-1を n まで格納 & 計算するための配列
	var aList []int

	// 固定値 1, 2 をセット
	aList = append(aList, 1)
	aList = append(aList, 1)

	// 1,2 は固定なので3番目から開始
	for i := 2; i < n; i++ {
		// a_{n-2} + a_{n-1}を計算
		x := aList[i-2] + aList[i-1]
		aList = append(aList, x)
	}

	return aList[n-1]
}
