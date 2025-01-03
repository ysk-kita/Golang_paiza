package algorithm

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	linearSearch3()
}

// スキャナ共通定義
var sc = bufio.NewScanner(os.Stdin)

// スキャナ定義(分割読込)
var sc2 *bufio.Scanner

// スキャナ定義(大規模バッファサイズ)
var sc3 *bufio.Scanner

// Reader定義
var reader = bufio.NewReader(os.Stdin)

func setSanner2() {
	sc2 = bufio.NewScanner(os.Stdin)
	// 読込を円滑にするために単語単位で読み込めるように
	sc2.Split(bufio.ScanWords)
	// デフォルトの64KBバッファから1MBバッファに更新
	const maxScanTokenSize = 1024 * 1024
	sc2.Buffer(make([]byte, maxScanTokenSize), maxScanTokenSize)
}

func setSanner3() {
	sc3 = bufio.NewScanner(os.Stdin)
	// デフォルトの64KBバッファから1MBバッファに更新
	const maxScanTokenSize = 1024 * 1024
	sc3.Buffer(make([]byte, maxScanTokenSize), maxScanTokenSize)
}

// 1行スキャン関数, 区切り文字指定
func getIntList(separator string, length int64) []int64 {
	sc.Scan()
	stringArray := strings.Split(sc.Text(), separator)

	// 変換して配列に格納
	var list []int64 = make([]int64, length)
	for i, v := range stringArray {
		a, _ := strconv.ParseInt(v, 10, 64)
		list[i] = a
	}
	return list
}

// 1行スキャン関数, 分割スキャン
func getIntListBySplitScan(length int64) []int64 {
	// 変換して配列に格納
	var list []int64 = make([]int64, length)
	i := 0
	for sc2.Scan() {
		a, _ := strconv.ParseInt(sc2.Text(), 10, 64)
		list[i] = a
		i++
	}
	return list
}

// 1行スキャン 大規模バッファ
func getIntListByLargeBufferScan(separator string, length int64) []int64 {
	sc3.Scan()
	stringArray := strings.Split(sc3.Text(), separator)

	// 変換して配列に格納
	var list []int64 = make([]int64, length)
	for i, v := range stringArray {
		a, _ := strconv.ParseInt(v, 10, 64)
		list[i] = a
	}
	return list
}

// 1行スキャン用関数 intList変換 Reader仕様
func getIntListByReader(length int64) []int64 {

	// 1行スキャン
	line, _ := reader.ReadString('\n')
	// 改行を取り除く
	line = strings.TrimSpace(line)
	// 文字列分割
	ary := strings.Split(line, " ")
	// 変換して格納
	var list []int64 = make([]int64, length)
	for i, v := range ary {
		a, _ := strconv.ParseInt(v, 10, 64)
		list[i] = a
	}
	return list
}

// 1行スキャン int変換
func getInt() int64 {
	sc.Scan()
	a, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return a
}

// 1行スキャン, sc2利用
func getIntBySc2() int64 {
	sc2.Scan()
	a, _ := strconv.ParseInt(sc2.Text(), 10, 64)
	return a
}

// 1行スキャン, sc3利用
func getIntBySc3() int64 {
	sc3.Scan()
	a, _ := strconv.ParseInt(sc3.Text(), 10, 64)
	return a
}

func getIntByReader() int64 {
	// 1行スキャン
	line, _ := reader.ReadString('\n')
	// 改行を取り除く
	line = strings.TrimSpace(line)
	// 数値変換
	a, _ := strconv.ParseInt(line, 10, 64)
	return a
}

// 最小値を返す
func minVal(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

// 最大値を返す
func maxVal(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// 線形探索問題1
// https://paiza.jp/works/mondai/sequence_search_problems/sequence_search_problems_search_value_boss/edit?language_uid=go
func linearSearch1() {
	n := getInt()
	list := getIntList(" ", n)
	k := getInt()

	var resList []int
	for i, v := range list {
		if v == k {
			// index番号を見るので+1
			resList = append(resList, i+1)
		}
	}
	for _, v := range resList {
		fmt.Println(v)
	}
}

// 線形探索問題2
// https://paiza.jp/works/mondai/sequence_search_problems/sequence_search_problems_minmax_boss/edit?language_uid=go
func linearSearch2() {
	/// 分割読み込み
	// setSanner2()
	// n := getIntBySc2()
	// list := getIntListBySplitScan(n)

	/// 大規模バッファ利用
	// setSanner3()
	// n := getIntBySc3()
	// list := getIntListByLargeBufferScan(" ", n)

	/// Reader利用
	n := getIntByReader()
	list := getIntListByReader(n)

	// ソートして最小値と最大値を取得
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	fmt.Println(list[n-1], list[0])
}

// 線形探索問題3
// https://paiza.jp/works/mondai/sequence_search_problems/sequence_search_problems_search_condition_boss/edit?language_uid=go
func linearSearch3() {
	var n, score int
	fmt.Scan(&n)

	var name string
	var names []string = make([]string, n)
	var scores []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &name, &score)
		names[i] = name
		scores[i] = score
	}
	var lower, upper int
	fmt.Scanf("%d %d", &lower, &upper)

	for i := 0; i < n; i++ {
		if lower <= scores[i] && scores[i] <= upper {
			fmt.Println(names[i])
		}
	}
}

// 線形探索問題
// https://paiza.jp/works/mondai/sequence_search_problems/sequence_search_problems_kthmax_boss/edit?language_uid=go
func linearSearch4() {
	setSanner3()
	n := getIntBySc3()
	list := getIntListByLargeBufferScan(" ", n)
	// 降順にソート
	sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })
	k := getIntBySc3()
	// index番号に補正
	fmt.Println(list[k-1])
}
