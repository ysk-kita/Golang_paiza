package main

import (
	"fmt"
	"math"
	"strconv"
)

func saidaikoyakusuCall1() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	// a_1 = 1, a_2=0, a_3=-1で、N,Kの余りがそのまま値を取り出せるようにリストを定義
	var aSuretsu []int = []int{-1, 1, 0}

	// 数列は 1, 0 ,-1の繰り返しなので、NとKの間の繰り返し相殺で0になる箇所は省略
	start := N % 3
	end := K % 3
	sum := 0
	if start == end { // 始点と終点が同じなら和ではない
		sum = aSuretsu[start]
	} else {
		if start < end {
			for i := start; i < end; i++ {
				sum += aSuretsu[i]
			}
		} else {
			sum += aSuretsu[start]
			// N%3のが大きくなるパターンは K%3=0, K%3=1のどちらか
			if end == 0 {
				sum += aSuretsu[0]
			} else {
				sum += aSuretsu[0] + aSuretsu[1]
			}
		}
	}
	fmt.Println(sum)
}

func saidaikoyakusuCall2() {
	var n int64
	fmt.Scan(&n)

	bin := strconv.FormatInt(int64(n), 2)

	// 2の階乗を求めるので底Nに2をセット
	N := 2
	ans, pow := 1, N
	// 2進数の桁数分処理を繰り返す
	length := len(bin)
	for i := 0; i < length; i++ {

		// bit値下1桁を取得
		if bin[len(bin)-1] == '1' {
			ans = ans * pow % 1000003
		}
		// i乗を求める
		pow = pow * pow % 1000003
		bin = bin[:len(bin)-1]
	}
	fmt.Println(ans)
}

func saidaikoyakusuCall3() {
	var N float64
	fmt.Scan(&N)
	fmt.Println(isSosu(N))
}

func isSosu(n float64) string {

	// 0, 1は素数ではない
	if n < 2 {
		return "NO"
	}

	// ルートnまでのループで良い
	// 合成数かどうかを判定し、合成数でないなら素数である。
	var isGouseisu bool
	for i := 2; i < int(math.Sqrt(n)); i++ {
		// 自身の値までに割り切れる数値があるならOut
		if int(n)%i == 0 {
			isGouseisu = true
			break
		}
	}
	if isGouseisu {
		return "NO"
	}
	return "Yes"
}

func SaidaiKoyakusu(a, b int) int {
	fmt.Println(a, b)
	for b != 0 {
		a, b = b, a%b // aをbに、bをa % bに置き換え
		fmt.Println(a, b)
	}
	return a // bが0になった時、aが最大公約数
}
