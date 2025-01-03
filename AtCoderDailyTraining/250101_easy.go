package main

import (
	"fmt"
)

func main() {
	var x, a, d, n int64
	fmt.Scanf("%d %d %d %d", &x, &a, &d, &n)
	// 先頭とのdiffを格納
	var diff int64
	if x > 0 {
		diff = x - a
	} else {
		diff = a - x
	}
	diff = abs(diff)

	// 差がなければそのまま終了
	if diff == 0 {
		fmt.Println(0)
		return
	}

	for i := int64(1); i < n; i++ {

		// 等差数列S_nの値と整数Xのdiffを確認
		sum := a + i*d

		var tmpDif int64
		if x > 0 {
			tmpDif = x - sum
		} else {
			tmpDif = sum - x
		}
		fmt.Println(tmpDif, x, sum)

		absDif := abs(tmpDif)

		// absDifの方が小さければ処理継続
		if diff > absDif {
			diff = absDif
		} else { // S_nとXの差が前回の物より広がっていればその時点で終了
			break
		}
	}

	fmt.Println(diff)
}

// 絶対値確認
func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func call3() {
	var l, r int
	var str string

	fmt.Scanf("%d %d", &l, &r)
	fmt.Scan(&str)

	// 反転対象取得
	tmp := str[l-1 : r]

	// 反転実行
	leng := len(tmp)
	tmp2 := ""
	for i := 0; i < leng; i++ {
		tmp2 = tmp2 + string(tmp[leng-1-i])
	}

	newStr := str[:l-1] + tmp2 + str[r:]
	fmt.Println(newStr)
}

func call2() {
	var N, modi, n int64
	n = 998244353
	fmt.Scan(&N)
	// debug
	//N = -9982443534

	// nは倍数、rは倍数にかける数値とする
	// 式 N - x = rn
	// -x = rn -N
	modi = N % n

	// 解が0以上の条件を満たしていれば終了
	if modi >= 0 {
		fmt.Println(modi)
		return
	}

	// 0未満であるなら 余り数値に加算し、倍数になるように修正する
	modi = modi + n
	fmt.Println(modi)
}

func call1() {
	var str string
	fmt.Scan(&str)

	// 同一文字の比較
	var sAry []string
	for i := 0; i < 3; i++ {
		tmp := str[i]

		// 値の比較 無ければ文字をパターンに追加
		if !checkChar(sAry, string(tmp)) {
			sAry = append(sAry, string(tmp))
		}
	}

	// 異なる文字の出現回数から、全3文字の総パターン数を求める
	// 3! / (文字1の個数! * 文字2の個数!.....文字kの個数!)の式
	upper := getFactorial(3)

	lower := 1
	switch len(sAry) {
	case 1: // 1種しかないパターンはそれ1つで3文字固定 (3,0,0)
		lower = getFactorial(3)
	case 2: // 2種のパターンは(1,2,0)、(2,1,0)等、[0]の位置は代わるが全て同じなので 2*1
		lower = getFactorial(2) * getFactorial(1)
	case 3: // 3種のパターンは(1,1,1)のみ
		lower = getFactorial(1) * getFactorial(1) * getFactorial(1)
	}
	result := upper / lower
	fmt.Println(result)
}

func getFactorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func checkChar(sAray []string, c string) bool {

	var res bool
	for _, v := range sAray {
		if v == c {
			res = true
			break
		}
	}
	return res
}
