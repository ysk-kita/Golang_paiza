package main

import (
	"fmt"
	"strconv"
	"strings"
)

type GotArea struct {
	x, y int
}

type GotAreaWithCount struct {
	x, y, cnt int
}

// 設問1
func Call() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	var ary []string
	var mys []GotArea

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		ary = append(ary, str)

		// 自陣地情報があれば座標を保存
		if strings.Contains(str, "*") {
			x := GetSingleArea(str)
			my := GotArea{x, i}
			mys = append(mys, my)
		}
	}
	// 座標上書き
	resAry := RobArea(ary, mys, H, W)

	for i := 0; i < H; i++ {
		fmt.Println(resAry[i])
	}
}

// 設問2
func Call2() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	var ary []string
	var mys []GotArea

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		ary = append(ary, str)

		// 自陣地情報があれば座標を保存
		if strings.Contains(str, "*") {
			x := GetSingleArea(str)
			my := GotArea{x, i}
			mys = append(mys, my)
		}
	}
	// 座標上書き
	resAry := RobAreaButBlockWall(ary, mys, H, W)

	for i := 0; i < H; i++ {
		fmt.Println(resAry[i])
	}
}

// 設問3
func Call3() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	var ary []string
	var mys []GotArea

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		ary = append(ary, str)

		// 自陣地情報があれば座標を保存
		if strings.Contains(str, "*") {
			x := GetSingleArea(str)
			my := GotArea{x, i}
			mys = append(mys, my)
		}
	}
	// 座標上書き
	resAry := RobAreaWithQueue(ary, mys, H, W)

	for i := 0; i < H; i++ {
		fmt.Println(resAry[i])
	}
}

// 設問4 マス埋めを操作回数に
func Call4() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	var ary []string
	var mys []GotAreaWithCount

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		// 自陣地情報があれば座標を保存
		if strings.Contains(str, "*") {
			x := GetSingleArea(str)
			my := GotAreaWithCount{x, i, 0}
			mys = append(mys, my)

			// 座標を変換
			str = strings.Replace(str, "*", "0", -1)
		}
		ary = append(ary, str)
	}
	// 座標上書き
	resAry := RobAreaToDisplayCount(ary, mys, H, W)

	for i := 0; i < H; i++ {
		fmt.Println(resAry[i])
	}
}

// 設問5 マス埋めしつつ、指定距離ならマス埋め方法を変更
func Call5() {
	var H, W, N int
	fmt.Scanf("%d %d %d", &H, &W, &N)

	var ary []string
	var mys []GotAreaWithCount

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		// 自陣地情報があれば座標を保存
		if strings.Contains(str, "*") {
			x := GetSingleArea(str)
			my := GotAreaWithCount{x, i, 0}
			mys = append(mys, my)

			// 座標を変換
			// str = strings.Replace(str, "*", "0", -1)
		}
		ary = append(ary, str)
	}

	// アクション距離取得
	var distances []int
	for i := 0; i < N; i++ {
		var n int
		fmt.Scan(&n)
		distances = append(distances, n)
	}

	// 座標上書き
	resAry := RobAreaAndDistanceAction(ary, mys, H, W, distances)

	for i := 0; i < H; i++ {
		fmt.Println(resAry[i])
	}
}

// 設問6 A vs B
func Call6() {
	var H, W int
	fmt.Scanf("%d %d %d", &H, &W)

	var areaMap []string
	var areaA, areaB []GotAreaWithCount

	var first string
	fmt.Scan(&first)

	// 陣地の取得　兼、自陣地座標の取得
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		// A陣地情報があれば座標を保存
		if strings.Contains(str, "A") {
			x := GetSingleAreaPrefix(str, 'A')
			point := GotAreaWithCount{x, i, 0}
			areaA = append(areaA, point)
		}

		// B陣地情報があれば座標を保存
		if strings.Contains(str, "B") {
			x := GetSingleAreaPrefix(str, 'B')
			point := GotAreaWithCount{x, i, 0}
			areaB = append(areaB, point)
		}
		areaMap = append(areaMap, str)
	}
	// 座標上書き
	hoge := RobAreaBattle(areaMap, areaA, areaB, H, W, first)
	for _, v := range hoge {
		fmt.Println(v)
	}
}

// "*"のX座標を返す
func GetSingleArea(str string) int {
	var res int

	for i := 0; i < len(str); i++ {

		if str[i] == '*' {
			res = i
			break
		}
	}
	return res
}

// 与えられた文字のx座標を返す
func GetSingleAreaPrefix(str string, prefix rune) int {
	var res int

	for i := 0; i < len(str); i++ {

		if str[i] == byte(prefix) {
			res = i
			break
		}
	}
	return res
}

// 自陣地の上下左右を奪う
func RobArea(ary []string, mys []GotArea, H, W int) []string {
	// 事前に保存した自陣地文だけエリア強盗
	for i := 0; i < len(mys); i++ {
		// 基準点取得
		my := mys[i]

		// 上方向
		ary = RobAreaUpper(ary, my)

		// 下方向
		ary = RobAreaLower(ary, my, H)

		// 左方向
		ary = RobAreaLeft(ary, my)

		// 右方向
		ary = RobAreaRight(ary, my, W)
	}
	return ary
}

// 上のエリアを奪う
func RobAreaUpper(ary []string, my GotArea) []string {

	// y座標が0なら終了
	if my.y == 0 {
		return ary
	}
	// 1行上を取得
	str := ary[my.y-1]
	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y-1] = robbed
	return ary
}

// 下のエリアを奪う
func RobAreaLower(ary []string, my GotArea, H int) []string {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return ary
	}
	// 1行下を取得
	str := ary[my.y+1]
	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y+1] = robbed
	return ary
}

// 左のエリアを奪う
func RobAreaLeft(ary []string, my GotArea) []string {
	// x座標が0なら終了
	if my.x == 0 {
		return ary
	}

	// 同じ行の左マスを上書き
	str := ary[my.y]
	robbed := str[:my.x-1] + "*" + str[my.x:]
	ary[my.y] = robbed
	return ary
}

// 右のエリアを奪う
func RobAreaRight(ary []string, my GotArea, W int) []string {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return ary
	}

	// 同じ行の右マスを上書き
	str := ary[my.y]

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	robbed := ""
	if my.x+1 == W-1 {
		robbed = str[:my.x+1] + "*"
	} else {
		robbed = str[:my.x+1] + "*" + str[my.x+2:]
	}
	ary[my.y] = robbed
	return ary
}

// 自陣地の上下左右を奪う、ただし壁なら奪えない
func RobAreaButBlockWall(ary []string, mys []GotArea, H, W int) []string {
	// 事前に保存した自陣地文だけエリア強盗
	for i := 0; i < len(mys); i++ {
		// 基準点取得
		my := mys[i]

		// 上方向
		ary = RobAreaUpperButBlockWall(ary, my)

		// 下方向
		ary = RobAreaLowerButBlockWall(ary, my, H)

		// 左方向
		ary = RobAreaLeftButBlockWall(ary, my)

		// 右方向
		ary = RobAreaRightButBlockWall(ary, my, W)
	}
	return ary
}

// 上のエリアを奪う、ただし壁なら奪えない
func RobAreaUpperButBlockWall(ary []string, my GotArea) []string {

	// y座標が0なら終了
	if my.y == 0 {
		return ary
	}
	// 1行上を取得
	str := ary[my.y-1]

	// 1行上が壁なら終了
	if str[my.x] == '#' {
		return ary
	}
	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y-1] = robbed
	return ary
}

// 下のエリアを奪う、ただし壁なら奪えない
func RobAreaLowerButBlockWall(ary []string, my GotArea, H int) []string {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return ary
	}
	// 1行下を取得
	str := ary[my.y+1]

	// 1行下が壁なら終了
	if str[my.x] == '#' {
		return ary
	}

	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y+1] = robbed
	return ary
}

// 左のエリアを奪う、ただし壁なら奪えない
func RobAreaLeftButBlockWall(ary []string, my GotArea) []string {
	// x座標が0なら終了
	if my.x == 0 {
		return ary
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁なら終了
	if str[my.x-1] == '#' {
		return ary
	}

	// 同じ行の左マスを上書き
	robbed := str[:my.x-1] + "*" + str[my.x:]
	ary[my.y] = robbed
	return ary
}

// 右のエリアを奪う、ただし壁なら奪えない
func RobAreaRightButBlockWall(ary []string, my GotArea, W int) []string {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return ary
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁なら終了
	if str[my.x+1] == '#' {
		return ary
	}

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	robbed := ""
	if my.x+1 == W-1 {
		robbed = str[:my.x+1] + "*"
	} else {
		robbed = str[:my.x+1] + "*" + str[my.x+2:]
	}
	ary[my.y] = robbed
	return ary
}

// 自陣地の上下左右を奪う、ただし壁なら奪えない
// また追加したエリアを奪っていく
func RobAreaWithQueue(ary []string, mys []GotArea, H, W int) []string {
	// キュー定義
	var queue []GotArea

	// 基準点は1つ
	my := mys[0]

	// 行動不能になるまでループ
	for {

		// 上方向
		ary, queue = RobUpperAreaWithQueue(ary, my, queue)

		// 下方向
		ary, queue = RobLowerAreaWithQueue(ary, my, H, queue)

		// 左方向
		ary, queue = RobLeftAreaWithQueue(ary, my, queue)

		// 右方向
		ary, queue = RobRightAreaWithQueue(ary, my, W, queue)

		// キューを使い切ったら終了
		if len(queue) == 0 {
			break
		}
		my = queue[0]
		queue = queue[1:] // キュー先頭削除
	}
	return ary
}

// Queue情報を読み取り、次の行動予定のQueueは動けるか判定
func IsMobableQueueArea(ary []string, queue []GotArea, H, W int) bool {
	// 準備
	next := queue[0]

	// 上下左右チェック
	isUpper := IsRobUpperArea(ary, next)
	isLower := IsRobLowerArea(ary, next, H)
	isLeft := IsRobLeftArea(ary, next)
	isRight := IsRobRightArea(ary, next, W)

	// どこか1つでも奪えるなら継続
	return isUpper || isLower || isLeft || isRight
}

// 上のエリアを奪えるかチェック
func IsRobUpperArea(ary []string, my GotArea) bool {
	// y座標が0なら終了
	if my.y == 0 {
		return false
	}
	// 上エリア判定 '.'なら奪える
	str := ary[my.y-1]
	return str[my.x] == '.'
}

// 下のエリアを奪えるかチェック
func IsRobLowerArea(ary []string, my GotArea, H int) bool {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return false
	}
	// 下エリア判定 '.'なら奪える
	str := ary[my.y+1]
	return str[my.x] == '.'
}

// 左のエリアを奪えるかチェック
func IsRobLeftArea(ary []string, my GotArea) bool {
	// x座標が0なら終了
	if my.x == 0 {
		return false
	}

	// 左エリア判定 '.'なら奪える
	str := ary[my.y]
	return str[my.x-1] == '.'
}

// 右のエリアを奪えるかチェック
func IsRobRightArea(ary []string, my GotArea, W int) bool {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return false
	}

	// 右エリア判定 '.'なら奪える
	str := ary[my.y]
	return str[my.x+1] == '.'
}

// 上のエリアを奪う、ただし壁なら奪えない
// 上のエリアを次の処理判定に追加する
func RobUpperAreaWithQueue(ary []string, my GotArea, queue []GotArea) ([]string, []GotArea) {

	// y座標が0なら終了
	if my.y == 0 {
		return ary, queue
	}
	// 1行上を取得
	str := ary[my.y-1]

	// 1行上が壁 または、自陣地なら終了
	if str[my.x] == '#' || str[my.x] == '*' {
		return ary, queue
	}
	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y-1] = robbed

	// Queueに座標追加
	queue = append(queue, GotArea{my.x, my.y - 1})

	return ary, queue
}

// 下のエリアを奪う、ただし壁なら奪えない
func RobLowerAreaWithQueue(ary []string, my GotArea, H int, queue []GotArea) ([]string, []GotArea) {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return ary, queue
	}
	// 1行下を取得
	str := ary[my.y+1]

	// 1行下が壁なら終了
	if str[my.x] == '#' || str[my.x] == '*' {
		return ary, queue
	}

	// 上書き
	robbed := str[:my.x] + "*" + str[my.x+1:]
	ary[my.y+1] = robbed

	// Queueに座標追加
	queue = append(queue, GotArea{my.x, my.y + 1})

	return ary, queue
}

// 左のエリアを奪う、ただし壁なら奪えない
func RobLeftAreaWithQueue(ary []string, my GotArea, queue []GotArea) ([]string, []GotArea) {
	// x座標が0なら終了
	if my.x == 0 {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁なら終了
	if str[my.x-1] == '#' || str[my.x-1] == '*' {
		return ary, queue
	}

	// 同じ行の左マスを上書き
	robbed := str[:my.x-1] + "*" + str[my.x:]
	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotArea{my.x - 1, my.y})

	return ary, queue
}

// 右のエリアを奪う、ただし壁なら奪えない
func RobRightAreaWithQueue(ary []string, my GotArea, W int, queue []GotArea) ([]string, []GotArea) {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁なら終了
	if str[my.x+1] == '#' || str[my.x+1] == '*' {
		return ary, queue
	}

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	robbed := ""
	if my.x+1 == W-1 {
		robbed = str[:my.x+1] + "*"
	} else {
		robbed = str[:my.x+1] + "*" + str[my.x+2:]
	}
	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotArea{my.x + 1, my.y})
	return ary, queue
}

// ---
// エリアを操作回数で上書きする
func RobAreaToDisplayCount(ary []string, mys []GotAreaWithCount, H, W int) []string {
	resAry := ary

	// キュー定義
	var queue []GotAreaWithCount

	// 基準点は1つ
	my := mys[0]

	// 行動不能になるまでループ
	for {

		// 上方向
		ary, queue = RobUpperAreaToDisplayCount(ary, my, queue)

		// 下方向
		ary, queue = RobLowerAreaToDisplayCount(ary, my, H, queue)

		// 左方向
		ary, queue = RobLeftAreaToDisplayCount(ary, my, queue)

		// 右方向
		ary, queue = RobRightAreaToDisplayCount(ary, my, W, queue)

		// キューを使い切ったら終了
		if len(queue) == 0 {
			break
		}
		my = queue[0]
		queue = queue[1:] // キュー先頭削除
	}
	return resAry
}

// 上のエリアを奪う、ただし壁なら奪えない
// 上のエリアを次の処理判定に追加する
func RobUpperAreaToDisplayCount(ary []string, my GotAreaWithCount, queue []GotAreaWithCount) ([]string, []GotAreaWithCount) {

	// y座標が0なら終了
	if my.y == 0 {
		return ary, queue
	}
	// 1行上を取得
	str := ary[my.y-1]

	// 1行上が壁 または、自陣地なら終了
	if str[my.x] != '.' {
		return ary, queue
	}
	// 現在の座標から1マス動いた状態を保存
	robbed := str[:my.x] + strconv.Itoa(my.cnt+1) + str[my.x+1:]
	ary[my.y-1] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x, my.y - 1, my.cnt + 1})

	return ary, queue
}

// 下のエリアを奪う、ただし壁なら奪えない
func RobLowerAreaToDisplayCount(ary []string, my GotAreaWithCount, H int, queue []GotAreaWithCount) ([]string, []GotAreaWithCount) {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return ary, queue
	}
	// 1行下を取得
	str := ary[my.y+1]

	// 1行下が壁または自陣地なら終了
	if str[my.x] != '.' {
		return ary, queue
	}

	// 現在の座標から1マス動いた状態を保存
	robbed := str[:my.x] + strconv.Itoa(my.cnt+1) + str[my.x+1:]
	ary[my.y+1] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x, my.y + 1, my.cnt + 1})

	return ary, queue
}

// 左のエリアを奪う、ただし壁なら奪えない
func RobLeftAreaToDisplayCount(ary []string, my GotAreaWithCount, queue []GotAreaWithCount) ([]string, []GotAreaWithCount) {
	// x座標が0なら終了
	if my.x == 0 {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁または自陣地なら終了
	if str[my.x-1] != '.' {
		return ary, queue
	}

	// 同じ行の左マスを上書き
	// 現在の座標から1マス動いたカウントを保存
	robbed := str[:my.x-1] + strconv.Itoa(my.cnt+1) + str[my.x:]
	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x - 1, my.y, my.cnt + 1})

	return ary, queue
}

// 右のエリアを奪う、ただし壁なら奪えない
func RobRightAreaToDisplayCount(ary []string, my GotAreaWithCount, W int, queue []GotAreaWithCount) ([]string, []GotAreaWithCount) {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁または自陣地なら終了
	if str[my.x+1] != '.' {
		return ary, queue
	}

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	// 現在の座標から1マス動いたカウントを保存
	robbed := ""
	if my.x+1 == W-1 {
		robbed = str[:my.x+1] + strconv.Itoa(my.cnt+1)
	} else {
		robbed = str[:my.x+1] + strconv.Itoa(my.cnt+1) + str[my.x+2:]
	}
	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x + 1, my.y, my.cnt + 1})
	return ary, queue
}

// 自陣地の上下左右を奪う、ただし壁なら奪えない
// また追加したエリアを奪っていく
func RobAreaAndDistanceAction(ary []string, mys []GotAreaWithCount, H, W int, distances []int) []string {
	resAry := ary

	// キュー定義
	var queue []GotAreaWithCount

	// 基準点は1つ
	my := mys[0]

	// 開始位置がイベント距離かチェック
	if DistanceCheck(distances, 0) {
		str := ary[my.y]
		robbed := str[:my.x] + "?" + str[my.x+1:]
		ary[my.y] = robbed
	}

	// 行動不能になるまでループ
	for {

		// 上方向
		ary, queue = RobUpperAreaAndDistanceAction(ary, my, queue, distances)

		// 下方向
		ary, queue = RobLowerAreaAndDistanceAction(ary, my, H, queue, distances)

		// 左方向
		ary, queue = RobLeftAreaAndDistanceAction(ary, my, queue, distances)

		// 右方向
		ary, queue = RobRightAreaAndDistanceAction(ary, my, W, queue, distances)

		// キューを使い切ったら終了
		if len(queue) == 0 {
			break
		}
		my = queue[0]
		queue = queue[1:] // キュー先頭削除
	}
	return resAry
}

// 上のエリアを奪う、ただし壁なら奪えない
// 上のエリアを次の処理判定に追加する
func RobUpperAreaAndDistanceAction(ary []string, my GotAreaWithCount, queue []GotAreaWithCount, distances []int) ([]string, []GotAreaWithCount) {

	// y座標が0なら終了
	if my.y == 0 {
		return ary, queue
	}
	// 1行上を取得
	str := ary[my.y-1]

	// 1行上が壁 または、自陣地なら終了
	if str[my.x] != '.' {
		return ary, queue
	}
	// 現在の座標から1マス動いた状態を保存
	// ただし、動いた位置がアクション距離ならマス目の状態を'?'に変更
	robbed := ""
	if DistanceCheck(distances, my.cnt+1) {
		robbed = str[:my.x] + "?" + str[my.x+1:]
	} else {
		robbed = str[:my.x] + "*" + str[my.x+1:]
	}
	ary[my.y-1] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x, my.y - 1, my.cnt + 1})

	return ary, queue
}

// 下のエリアを奪う、ただし壁なら奪えない
func RobLowerAreaAndDistanceAction(ary []string, my GotAreaWithCount, H int, queue []GotAreaWithCount, distances []int) ([]string, []GotAreaWithCount) {
	// 奪うy座標がエリアの高さを超えてれば終了
	if my.y+1 == H {
		return ary, queue
	}
	// 1行下を取得
	str := ary[my.y+1]

	// 1行下が壁または自陣地なら終了
	if str[my.x] != '.' {
		return ary, queue
	}

	// 現在の座標から1マス動いた状態を保存
	// ただし、動いた位置がアクション距離ならマス目の状態を'?'に変更
	robbed := ""
	if DistanceCheck(distances, my.cnt+1) {
		robbed = str[:my.x] + "?" + str[my.x+1:]
	} else {
		robbed = str[:my.x] + "*" + str[my.x+1:]
	}

	ary[my.y+1] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x, my.y + 1, my.cnt + 1})

	return ary, queue
}

// 左のエリアを奪う、ただし壁なら奪えない
func RobLeftAreaAndDistanceAction(ary []string, my GotAreaWithCount, queue []GotAreaWithCount, distances []int) ([]string, []GotAreaWithCount) {
	// x座標が0なら終了
	if my.x == 0 {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁または自陣地なら終了
	if str[my.x-1] != '.' {
		return ary, queue
	}

	// 同じ行の左マスを上書き
	// 現在の座標から1マス動いたカウントを保存
	// ただし、動いた位置がアクション距離ならマス目の状態を'?'に変更
	robbed := ""
	if DistanceCheck(distances, my.cnt+1) {
		robbed = str[:my.x-1] + "?" + str[my.x:]
	} else {
		robbed = str[:my.x-1] + "*" + str[my.x:]
	}

	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x - 1, my.y, my.cnt + 1})

	return ary, queue
}

// 右のエリアを奪う、ただし壁なら奪えない
func RobRightAreaAndDistanceAction(ary []string, my GotAreaWithCount, W int, queue []GotAreaWithCount, distances []int) ([]string, []GotAreaWithCount) {
	// 奪うx座標が横幅越えなら終了
	if my.x+1 == W {
		return ary, queue
	}

	// 同じ行を取得
	str := ary[my.y]

	// 左が壁または自陣地なら終了
	if str[my.x+1] != '.' {
		return ary, queue
	}

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	// 現在の座標から1マス動いたカウントを保存
	// ただし、動いた位置がアクション距離ならマス目の状態を'?'に変更
	robbed := ""
	if DistanceCheck(distances, my.cnt+1) {
		if my.x+1 == W-1 {
			robbed = str[:my.x+1] + "?"
		} else {
			robbed = str[:my.x+1] + "?" + str[my.x+2:]
		}
	} else {
		if my.x+1 == W-1 {
			robbed = str[:my.x+1] + "*"
		} else {
			robbed = str[:my.x+1] + "*" + str[my.x+2:]
		}
	}
	ary[my.y] = robbed

	// Queueに座標追加
	queue = append(queue, GotAreaWithCount{my.x + 1, my.y, my.cnt + 1})
	return ary, queue
}

// 判定チェック
func DistanceCheck(distances []int, x int) bool {
	var res bool = false
	for _, v := range distances {
		if v == x {
			res = true
			break
		}
	}

	return res
}

// 自陣地の上下左右を奪う、ただし壁なら奪えない
// また追加したエリアを奪っていく
func RobAreaBattle(areaMap []string, areaA, areaB []GotAreaWithCount, H, W int, first string) []string {

	// キュー定義
	var queueA, queueB []GotAreaWithCount

	// 手番を設定
	turn := first

	// 初手のエリア設定
	pointA, pointB := areaA[0], areaB[0]

	// 移動不可フラグ
	var turnSkipA, turnSkipB bool = false, false

	// 使い切るまで奪取を行うエリアカウント
	nowAreaCountA, nowAreaCountB := 0, 0

	// 行動不能になるまでループ
	for {

		/*
			課題
			　キューとして先に確保したエリアが"."のままなので、奪われてしまっている
			対策
			　相手のエリア状態をいれたSliceにあるかチェックを追加
			→　そもそもロジックがダメだった。今のロジックは1手ずつだが、問題文は1手目で広がったマス全てを連続処理してエリアを広げるようにする必要がある
			→　どうすればよいか？
			　→　エリアカウント有の方を使う。そして、エリアカウントnを持っているエリアが全て広げる。なので、エリアa,bを渡して個数を確保しなくても良い
		*/

		// 手番A
		if turn == "A" && !turnSkipA {

			// キューに同じエリアカウントがあるならそれを使い切るまで奪取を続ける
			for {
				areaMap, queueA = RobAreaByPlayer(areaMap, pointA, queueA, H, W, "A")

				// キューが1つ以上あるとき、次のキューのエリアカウントを確認
				if len(queueA) > 0 {
					// Aのカウントが同じならもう1度手番を行う
					if nowAreaCountA == queueA[0].cnt {
						pointA = queueA[0]
						queueA = queueA[1:] // キュー先頭削除
					} else {
						// カウントがずれたら手番終了
						break
					}
				} else {
					break
				}
			}

			// キューを使い切ったら終了
			if len(queueA) == 0 {
				turnSkipA = true
			} else {
				pointA = queueA[0]
				queueA = queueA[1:] // キュー先頭削除
				// 次の連続手番を行うエリアカウントにずらす
				nowAreaCountA++
			}
		}

		// 手番B
		if turn == "B" && !turnSkipB {
			// キューに同じエリアカウントがあるならそれを使い切るまで奪取を続ける
			for {
				areaMap, queueB = RobAreaByPlayer(areaMap, pointB, queueB, H, W, "B")

				// キューが1つ以上あるとき、次のキューのエリアカウントを確認
				if len(queueB) > 0 {
					// Bのカウントが同じならもう1度手番を行う
					if nowAreaCountB == queueB[0].cnt {
						pointB = queueB[0]
						queueB = queueB[1:] // キュー先頭削除
					} else {
						// カウントがずれたら手番終了
						break
					}
				} else {
					break
				}
			}

			// キューを使い切ったら終了
			if len(queueB) == 0 {
				turnSkipB = true
			} else {
				pointB = queueB[0]
				queueB = queueB[1:] // キュー先頭削除

				// 次の連続手番を行うエリアカウントにずらす
				nowAreaCountB++
			}
		}
		// どちらもターンスキップ状態になったら終了
		if turnSkipA && turnSkipB {
			break
		}

		// 手番を入れ替え
		if turn == "A" || turnSkipA {
			turn = "B"
		} else if turn == "B" || turnSkipB {
			turn = "A"
		}

	}

	// 各ユーザのエリア量表示
	cntA, cntB := 0, 0
	for _, v := range areaMap {
		cntA += strings.Count(v, "A")
		cntB += strings.Count(v, "B")
	}

	fmt.Println(cntA, cntB)

	// 勝利判定
	if cntA > cntB {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	return areaMap
}

// 与えられた任意のユーザの陣地周辺をそのユーザの陣地にする
func RobAreaByPlayer(areaMap []string, pointX GotAreaWithCount, queueX []GotAreaWithCount, H, W int, userTag string) ([]string, []GotAreaWithCount) {
	// 上方向
	areaMap, queueX = RobUpperAreaByPlayer(areaMap, pointX, queueX, userTag)

	// 下方向
	areaMap, queueX = RobLowerAreaByPlayer(areaMap, pointX, H, queueX, userTag)

	// 左方向
	areaMap, queueX = RobLeftAreaByPlayer(areaMap, pointX, queueX, userTag)

	// 右方向
	areaMap, queueX = RobRightAreaByPlayer(areaMap, pointX, W, queueX, userTag)

	return areaMap, queueX
}

// 上のエリアを奪う、ただし壁なら奪えない
// 上のエリアを次の処理判定に追加する
func RobUpperAreaByPlayer(areaMap []string, pointX GotAreaWithCount, queueX []GotAreaWithCount, userTag string) ([]string, []GotAreaWithCount) {

	// y座標が0なら終了
	if pointX.y == 0 {
		return areaMap, queueX
	}
	// 1行上を取得
	str := areaMap[pointX.y-1]

	// 1行上が壁 または、自陣地なら終了
	if str[pointX.x] != '.' {
		return areaMap, queueX
	}
	// 上書き
	robbed := str[:pointX.x] + userTag + str[pointX.x+1:]
	areaMap[pointX.y-1] = robbed

	// Queueに座標追加
	queueX = append(queueX, GotAreaWithCount{pointX.x, pointX.y - 1, pointX.cnt + 1})

	return areaMap, queueX
}

// 下のエリアを奪う、ただし壁なら奪えない
func RobLowerAreaByPlayer(areaMap []string, pointX GotAreaWithCount, H int, queueX []GotAreaWithCount, userTag string) ([]string, []GotAreaWithCount) {
	// 奪うy座標がエリアの高さを超えてれば終了
	if pointX.y+1 == H {
		return areaMap, queueX
	}
	// 1行下を取得
	str := areaMap[pointX.y+1]

	// 1行下が壁なら終了
	if str[pointX.x] != '.' {
		return areaMap, queueX
	}

	// 上書き
	robbed := str[:pointX.x] + userTag + str[pointX.x+1:]
	areaMap[pointX.y+1] = robbed

	// Queueに座標追加
	queueX = append(queueX, GotAreaWithCount{pointX.x, pointX.y + 1, pointX.cnt + 1})

	return areaMap, queueX
}

// 左のエリアを奪う、ただし壁なら奪えない
func RobLeftAreaByPlayer(areaMap []string, pointX GotAreaWithCount, queueX []GotAreaWithCount, userTag string) ([]string, []GotAreaWithCount) {
	// x座標が0なら終了
	if pointX.x == 0 {
		return areaMap, queueX
	}

	// 同じ行を取得
	str := areaMap[pointX.y]

	// 左が壁なら終了
	if str[pointX.x-1] != '.' {
		return areaMap, queueX
	}

	// 同じ行の左マスを上書き
	robbed := str[:pointX.x-1] + userTag + str[pointX.x:]
	areaMap[pointX.y] = robbed

	// Queueに座標追加
	queueX = append(queueX, GotAreaWithCount{pointX.x - 1, pointX.y, pointX.cnt + 1})

	return areaMap, queueX
}

// 右のエリアを奪う、ただし壁なら奪えない
func RobRightAreaByPlayer(areaMap []string, pointX GotAreaWithCount, W int, queueX []GotAreaWithCount, userTag string) ([]string, []GotAreaWithCount) {
	// 奪うx座標が横幅越えなら終了
	if pointX.x+1 == W {
		return areaMap, queueX
	}

	// 同じ行を取得
	str := areaMap[pointX.y]

	// 左が壁なら終了
	if str[pointX.x+1] != '.' {
		return areaMap, queueX
	}

	// 右端のマスを更新する場合とそうでない場合で処理を変える(x+1が更新したいマス)
	robbed := ""
	if pointX.x+1 == W-1 {
		robbed = str[:pointX.x+1] + userTag
	} else {
		robbed = str[:pointX.x+1] + userTag + str[pointX.x+2:]
	}
	areaMap[pointX.y] = robbed

	// Queueに座標追加
	queueX = append(queueX, GotAreaWithCount{pointX.x + 1, pointX.y, pointX.cnt + 1})

	return areaMap, queueX
}
