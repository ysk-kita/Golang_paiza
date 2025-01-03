package main

import (
	"fmt"
)

// mainから起動する用
func call_ruisekiwaRuisekiwa() {

}

// 課題1
func call_ruisekiwa1() {
	var N, a, sum int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		sum += a
		fmt.Println(sum)
	}
}

func call_ruisekiwa2() {
	var N, a, n, sum, l, u int
	var list []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &l, &u)
		// 累積和計算
		sum = 0

		for j := l; j <= u; j++ {
			sum += list[j]
		}
		fmt.Println(sum)
	}
}

func call_ruisekiwa3() {
	var N, M, a int
	fmt.Scanf("%d %d", &N, &M)
	var list []int

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}

	fmt.Println(SyakutoriMethodMin(list, M))
}

// 尺取り法で区間の長さを求める（最小）
func SyakutoriMethodMin(list []int, checkVal int) int {
	head, tail := 0, 0
	sum, ans := 0, len(list)+1
	for {
		// 区間の和が条件値を超えた場合
		if sum >= checkVal {
			// 条件を満たしているので、現在の長さと超えた瞬間の長さを比較
			ans = min(ans, head-tail)
			sum -= list[tail]
			tail += 1
		} else {
			if head == len(list) {
				break
			}
			sum += list[head]
			head += 1
		}
	}
	// 区間長が配列全体のままなら条件を満たせていない
	if ans == len(list)+1 {
		ans = -1
	}
	return ans
}

func call_ruisekiwa4() {
	var N, M, a int
	fmt.Scanf("%d %d", &N, &M)
	var list []int

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}

	fmt.Println(SyakutoriMethodMax(list, M))
}

// 尺取り法で区間の長さを求める（最大）
// 0....Nまで計測する際、Nの方に左記に向かうのがhead,後から追いかけるのがtail
func SyakutoriMethodMax(list []int, checkVal int) int {
	head, tail := 0, 0
	sum, ans := 0, -1

	for {
		// 区間の和が長さを超えなかった場合、
		if sum <= checkVal {
			// 条件を満たしているので、現在の長さと超えた瞬間の長さを比較
			ans = max(ans, head-tail)
			if head == len(list) {
				break
			}
			// より長くするために先頭を１歩前に
			sum += list[head]
			head += 1

		} else {
			// 区間の和が超えてしまったので、最後尾を切り捨てる
			sum -= list[tail]
			tail += 1
		}
	}
	return ans
}

func call_ruisekiwa5() {
	var N, M, A, l, u, a int
	var list, addList []int
	// N M のスキャン
	fmt.Scanf("%d %d", &N, &M)

	// A_1 ... A_Nのスキャン
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		list = append(list, A)
	}

	/// n回のループを2層でやっているのでO(n^2)、データが増えると処理量が膨大になる
	// l_1 u_1 a_1 ... _nまでのスキャン
	// for i := 0; i < M; i++ {
	// 	fmt.Scanf("%d %d %d ", &l, &u, &a)
	// 	for j := l-1; j < u; j++ {
	// 		list[j] += a
	// 	}
	// }

	/// IMOS法を使って解決する
	// 長さNの加算用リストを作成
	for i := 0; i < N+1; i++ {
		addList = append(addList, 0)
	}
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d %d ", &l, &u, &a)
		//
		addList[l-1] += a
		if u < N {
			addList[u] -= a
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(list[i] + addList[i])
		if i < N {
			addList[i+1] += addList[i]
		}
	}
}

func call_ruisekiwa6() {
	var N, K, a int
	fmt.Scanf("%d %d", &N, &K)
	var list []int

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}

	fmt.Println(SyakutoriMethodMinByByMultiple(list, K))
}

// 尺取り法で区間積の長さを求める（最小）
func SyakutoriMethodMinByByMultiple(list []int, checkVal int) int {
	// 区間積を求めるためにprodに最初からlist[0]を入れる
	right, left := 1, 0
	prod, ans := list[0], len(list)+1

	for {
		// 区間の積が条件値を超えた場合
		if prod >= checkVal {
			// 条件を満たしているので、現在の長さと超えた瞬間の長さを比較
			ans = min(ans, right-left)
			prod /= list[left]
			left += 1

			// 区間長が"1"になったらそれ以上短くなりようがないので終了
			if ans == 1 {
				break
			}
		} else {
			// 端まで到着したら終了
			if right == len(list) {
				break
			}
			// 次の要素が"0"の場合、prodが強制的に0になるので区間をスキップ
			if list[right] == 0 {
				prod = list[right+1]
				left = right + 1
				right = right + 2
			} else { // そうでなければ乗算
				prod *= list[right]
				right += 1
			}
		}
	}
	return ans
}

func call_ruisekiwa7() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	// 累積和を求める
	var ruisekiwa []int = make([]int, n+1)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Scan(&a)
		ruisekiwa[i] = a + ruisekiwa[i-1]
	}

	// 最大の期間を求める
	var maxCount, kouhoDay, kouhoSu = 0, 1, 0
	for i := k; i <= n; i++ {
		// i日目時点の累積和 - からk日分より過去の累積和を除外
		ninzu := ruisekiwa[i] - ruisekiwa[i-k]
		// 最大人数が多ければ最短候補日、候補数をリセットして更新
		if maxCount < ninzu {
			maxCount = ninzu
			kouhoDay = i - k + 1
			kouhoSu = 1
		} else if maxCount == ninzu { // 最大人数が同じなら候補数を追加
			kouhoSu++
		}
	}
	fmt.Println(kouhoSu, kouhoDay)
}
