package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Oserocall_osero() {
	debugcall_osero6()
}

func debugcall_osero2() {
	fmt.Println("debug start")
	var H, W, Y, X int = 3, 3, 0, 0

	var field []string = []string{
		"..*",
		"...",
		"***",
	}
	field = ExecReversiVerticalAndHorizontal(H, W, Y, X, field)
	fmt.Println("-- finished --")
	for _, v := range field {
		fmt.Println(v)
	}
}

// 課題1
func call_osero1() {
	var H, W, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &Y, &X)

	for _, v := range getFieldLevel1(H, W, Y, X) {
		fmt.Println(v)
	}
}

// 指定した (x, y)座標の上下左右を端まで塗りつぶしたフィールドを生成
func getFieldLevel1(H, W, Y, X int) []string {
	var res []string

	for i := 0; i < H; i++ {
		var str string

		// 行生成
		for j := 0; j < W; j++ {
			// 縦, 横が指定座標なら"!"を指定
			if j == X && i == Y {
				str += "!"
			} else if j == X { // 指定座標と同じ列なら"*""
				str += "*"
			} else if i == Y { // 指定座標と同じ行なら"*""
				str += "*"
			} else { // それ以外は"."
				str += "."
			}
		}
		// 追加
		res = append(res, str)
	}
	return res
}

// 課題2
func call_osero2() {
	var H, W, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &Y, &X)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		field = append(field, str)
	}
	field = ExecReversiVerticalAndHorizontal(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

// 指定した上下左右マスの位置まで"*"に変更する
func ExecReversiVerticalAndHorizontal(H, W, Y, X int, field []string) []string {

	// 上下左右の更新終了マス取得
	upper := GetNearestUpperField(Y, X, field)
	lower := GetNearestLowerField(H, Y, X, field)
	left := GetNearestLeftField(Y, X, field)
	right := GetNearestRightField(W, Y, X, field)

	// 上下左右をひっくり返す
	field = ReverseUpperSquare(Y, X, upper, field)
	field = ReverseLowerSquare(H, Y, X, lower, field)
	field = ReverseLeftSquare(Y, X, left, field)
	field = ReverseRightSquare(Y, X, right, field)
	return field
}

// 指したマスをひっくり返す
func ReverseSetPosition(Y, X int, field []string) []string {
	str := field[Y]
	str = str[:X] + "*" + str[X+1:]
	field[Y] = str
	return field
}

// 上側をひっくり返す
func ReverseUpperSquare(Y, X, upper int, field []string) []string {
	// 更新無し
	if upper < 0 {
		return field
	}

	// マス更新
	for i := Y - 1; i >= upper; i-- {
		str := field[i]
		str = str[:X] + "*" + str[X+1:]
		field[i] = str
	}
	return field
}

// 下側をひっくり返す
func ReverseLowerSquare(H, Y, X, lower int, field []string) []string {
	// 更新無し
	if lower < 0 {
		return field
	}

	// マス更新
	for i := Y; i <= lower; i++ {
		str := field[i]
		str = str[:X] + "*" + str[X+1:]
		field[i] = str
	}
	return field
}

// 左側をひっくり返す
func ReverseLeftSquare(Y, X, left int, field []string) []string {
	// 更新無し
	if left < 0 {
		return field
	}

	// マス更新
	for i := X - 1; i >= left; i-- {
		str := field[Y]
		str = str[:i] + "*" + str[i+1:]
		field[Y] = str
	}
	return field
}

// 右側をひっくり返す
func ReverseRightSquare(Y, X, right int, field []string) []string {
	// 更新無し
	if right < 0 {
		return field
	}

	// マス更新
	for i := X + 1; i <= right; i++ {
		str := field[Y]
		str = str[:i] + "*" + str[i+1:]
		field[Y] = str
	}
	return field
}

// 上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段の場合は -1 を返す
func GetNearestUpperField(Y, X int, field []string) int {
	// 最上段範囲
	if Y == 0 {
		return -1
	}

	// 指定座標から 最上段までチェック
	res := -1
	for i := Y - 1; i >= 0; i-- {

		// 列の上マスが"*"なら終了
		str := field[i]
		if str[X] == '*' {
			res = i
			break
		}
	}
	return res
}

// 下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段の場合は -1 を返す
func GetNearestLowerField(H, Y, X int, field []string) int {
	// 最下段範囲
	if Y == H-1 {
		return -1
	}

	// 指定座標から 最下段までチェック
	res := -1
	for i := Y + 1; i < H; i++ {
		// 列の上マスが"*"なら終了
		str := field[i]
		if str[X] == '*' {
			res = i
			break
		}
	}
	return res
}

// 左方向で一番近い "*"マスの座標を返す
// 存在しない または 左端の場合は -1 を返す
func GetNearestLeftField(Y, X int, field []string) int {
	// 左端判定
	if X == 0 {
		return -1
	}

	// 指定座標から 左端までチェック
	res := -1
	for i := X - 1; i >= 0; i-- {
		// 行の左マスが"*"なら終了
		str := field[Y]
		if str[i] == '*' {
			res = i
			break
		}
	}
	return res
}

// 右方向で一番近い "*"マスの座標を返す
// 存在しない または 右端の場合は -1 を返す
func GetNearestRightField(W, Y, X int, field []string) int {
	// 左端判定
	if X == W-1 {
		return -1
	}

	// 指定座標から 右端までチェック
	res := -1
	for i := X + 1; i < W; i++ {
		// 行の右マスが"*"なら終了
		str := field[Y]
		if str[i] == '*' {
			res = i
			break
		}
	}
	return res
}

// 課題3
func call_osero3() {
	var H, W, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &Y, &X)
	// フィールド作成
	var field []string
	for i := 0; i < H; i++ {
		var str string
		for j := 0; j < W; j++ {
			str += "."
		}
		field = append(field, str)
	}
	// ひっくり返した奴を出力
	for _, v := range getFieldLevel3(H, W, Y, X, field) {
		fmt.Println(v)
	}
}

// 指定した (x, y)座標の斜め方向を端まで塗りつぶす
func getFieldLevel3(H, W, Y, X int, field []string) []string {

	// 指定座標を"!"に変更
	point := field[Y]
	point = point[:X] + "!" + point[X+1:]
	field[Y] = point

	// 指定座標の斜め方向を"*"に変更
	field = ReverseDiagonalUpperLeftSquare(H, W, Y, X, 0, 0, field)
	field = ReverseDiagonalUpperRightSquare(H, W, Y, X, 0, 0, field)
	field = ReverseDiagonalLowerLeftSquare(H, W, Y, X, 0, 0, field)
	field = ReverseDiagonalLowerRightSquare(H, W, Y, X, 0, 0, field)
	return field
}

// 外周判定を行う。上下左右1つでもtrueになればそれは場外
func isOutside(H, W, Y, X int) bool {
	topOut := Y < 0
	bottomOut := Y > H-1
	leftOut := X < 0
	rightOut := X > W-1
	return topOut || bottomOut || leftOut || rightOut
}

// 左上をひっくり返す
func ReverseDiagonalUpperLeftSquare(H, W, Y, X, stopY, stopX int, field []string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y-1, X-1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + "*" + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX--
		nowY--
	}
	return field
}

// 右上をひっくり返す
func ReverseDiagonalUpperRightSquare(H, W, Y, X, stopY, stopX int, field []string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y-1, X+1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + "*" + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX++
		nowY--
	}
	return field
}

// 左下をひっくり返す
func ReverseDiagonalLowerLeftSquare(H, W, Y, X, stopY, stopX int, field []string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y+1, X-1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + "*" + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX--
		nowY++
	}
	return field
}

// 右下をひっくり返す
func ReverseDiagonalLowerRightSquare(H, W, Y, X, stopY, stopX int, field []string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y+1, X+1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + "*" + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX++
		nowY++
	}
	return field
}

// 課題4
func call_osero4() {
	var H, W, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &Y, &X)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		field = append(field, str)
	}
	// 指定座標を"!"に変更
	point := field[Y]
	point = point[:X] + "*" + point[X+1:]
	field[Y] = point

	field = ExecReversiDiagonal(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

// オセロを行う(斜め方向)
func ExecReversiDiagonal(H, W, Y, X int, field []string) []string {
	// 斜めのの更新終了マス取得
	ulY, ulX := GetNearestDiagonalUpperLeftField(H, W, Y, X, field)
	urY, urX := GetNearestDiagonalUpperRightField(H, W, Y, X, field)
	llY, llX := GetNearestDiagonalLowerLeftField(H, W, Y, X, field)
	lrY, lrX := GetNearestDiagonalLowerRightField(H, W, Y, X, field)

	// マス更新実行
	field = ReverseDiagonalUpperLeftSquare(H, W, Y, X, ulY, ulX, field)
	field = ReverseDiagonalUpperRightSquare(H, W, Y, X, urY, urX, field)
	field = ReverseDiagonalLowerLeftSquare(H, W, Y, X, llY, llX, field)
	field = ReverseDiagonalLowerRightSquare(H, W, Y, X, lrY, lrX, field)
	return field
}

// 左上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段/左端の場合は -1 を返す
func GetNearestDiagonalUpperLeftField(H, W, Y, X int, field []string) (int, int) {
	// 左端か最上段なら終了
	if Y == 0 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y-1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX--
		nowY--
	}
	return nowY, nowX
}

// 右上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段・右端の場合は -1 を返す
func GetNearestDiagonalUpperRightField(H, W, Y, X int, field []string) (int, int) {
	// 右端か最上段なら終了
	if Y == 0 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y-1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX++
		nowY--
	}
	return nowY, nowX
}

// 左下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段/左端の場合は -1 を返す
func GetNearestDiagonalLowerLeftField(H, W, Y, X int, field []string) (int, int) {
	// 左端か最上段なら終了
	if Y == H-1 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y+1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX--
		nowY++
	}
	return nowY, nowX
}

// 右下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段・右端の場合は -1 を返す
func GetNearestDiagonalLowerRightField(H, W, Y, X int, field []string) (int, int) {
	// 右端か最上段なら終了
	if Y == H-1 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y+1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX++
		nowY++
	}
	return nowY, nowX
}

// 課題5
func call_osero5() {
	var H, W, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &Y, &X)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		field = append(field, str)
	}

	field = ExecReversi(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

// オセロロジック実行
func ExecReversi(H, W, Y, X int, field []string) []string {
	// 指したマスをひっくり返す
	field = ReverseSetPosition(Y, X, field)
	// 縦横をひっくり返す
	field = ExecReversiVerticalAndHorizontal(H, W, Y, X, field)
	// 斜めをひっくり返す
	field = ExecReversiDiagonal(H, W, Y, X, field)
	return field
}

// 課題5
func debugcall_osero5() {
	var H, W, Y, X = 5, 5, 2, 2

	var field []string = []string{
		"*.*.*",
		".....",
		"*...*",
		".....",
		"*.*.*",
	}
	field = ExecReversi(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

// 課題6
func call_osero6() {
	var H, W, Y, X int
	fmt.Scanf("%d %d", &H, &W)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)

		// 指した座標を取得
		if strings.Contains(str, "!") {
			Y = i

			for j := 0; j < W; j++ {
				if str[j] == '!' {
					X = j
					break
				}

			}
		}

		field = append(field, str)
	}

	field = ExecReversiWithWallBlock(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

func debugcall_osero6() {
	var H, W, Y, X = 5, 5, 2, 2

	var field []string = []string{
		"*.*.*",
		".....",
		"*#!.*",
		"...#*",
		"*.***",
	}
	field = ExecReversiWithWallBlock(H, W, Y, X, field)
	for _, v := range field {
		fmt.Println(v)
	}
}

// 上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段の場合は -1 を返す
func GetNearestUpperFieldWithWallBlock(Y, X int, field []string) int {
	// 最上段範囲
	if Y == 0 {
		return -1
	}

	// 指定座標から 最上段までチェック
	res := -1
	for i := Y - 1; i >= 0; i-- {

		// 列の上マスが"*"なら終了
		str := field[i]
		if str[X] == '*' {
			res = i
			break
		}

		// 壁にぶつかれば処理終了
		if str[X] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段の場合は -1 を返す
func GetNearestLowerFieldWithWallBlock(H, Y, X int, field []string) int {
	// 最下段範囲
	if Y == H-1 {
		return -1
	}

	// 指定座標から 最下段までチェック
	res := -1
	for i := Y + 1; i < H; i++ {
		// 列の上マスが"*"なら終了
		str := field[i]
		if str[X] == '*' {
			res = i
			break
		}

		// 壁にぶつかれば処理終了
		if str[X] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 左方向で一番近い "*"マスの座標を返す
// 存在しない または 左端の場合は -1 を返す
func GetNearestLeftFieldWithWallBlock(Y, X int, field []string) int {
	// 左端判定
	if X == 0 {
		return -1
	}

	// 指定座標から 左端までチェック
	res := -1
	for i := X - 1; i >= 0; i-- {

		// 行の左マスが"*"なら終了
		str := field[Y]
		if str[i] == '*' {
			res = i
			break
		}

		// 壁にぶつかれば処理終了
		if str[i] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 右方向で一番近い "*"マスの座標を返す
// 存在しない または 右端の場合は -1 を返す
func GetNearestRightFieldWithWallBlock(W, Y, X int, field []string) int {
	// 左端判定
	if X == W-1 {
		return -1
	}

	// 指定座標から 右端までチェック
	res := -1
	for i := X + 1; i < W; i++ {
		// 行の右マスが"*"なら終了
		str := field[Y]
		if str[i] == '*' {
			res = i
			break
		}

		// 壁にぶつかれば処理終了
		if str[i] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 左上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段/左端の場合は -1 を返す
func GetNearestDiagonalUpperLeftFieldWithWallBlock(H, W, Y, X int, field []string) (int, int) {
	// 左端か最上段なら終了
	if Y == 0 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y-1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク、壁にぶつかっても同様
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}
		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX--
		nowY--
	}
	return nowY, nowX
}

// 右上方向で一番近い "*"マスの座標を返す
// 存在しない または 最上段・右端の場合は -1 を返す
func GetNearestDiagonalUpperRightFieldWithWallBlock(H, W, Y, X int, field []string) (int, int) {
	// 右端か最上段なら終了
	if Y == 0 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y-1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク、壁にぶつかっても同様
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX++
		nowY--
	}
	return nowY, nowX
}

// 左下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段/左端の場合は -1 を返す
func GetNearestDiagonalLowerLeftFieldWithWallBlock(H, W, Y, X int, field []string) (int, int) {
	// 左端か最上段なら終了
	if Y == H-1 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y+1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク、壁にぶつかっても同様
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX--
		nowY++
	}
	return nowY, nowX
}

// 右下方向で一番近い "*"マスの座標を返す
// 存在しない または 最下段・右端の場合は -1 を返す
func GetNearestDiagonalLowerRightFieldWithWallBlock(H, W, Y, X int, field []string) (int, int) {
	// 右端か最上段なら終了
	if Y == H-1 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y+1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク、壁にぶつかっても同様
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// "*"が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == '*' {
			break
		}

		// 次の座標へ
		nowX++
		nowY++
	}
	return nowY, nowX
}

// オセロロジック実行(壁判定あり)
func ExecReversiWithWallBlock(H, W, Y, X int, field []string) []string {
	// 指したマスをひっくり返す
	field = ReverseSetPosition(Y, X, field)
	// 縦横をひっくり返す
	field = ExecReversiVerticalAndHorizontalWithWallBlock(H, W, Y, X, field)
	// 斜めをひっくり返す
	field = ExecReversiDiagonalWithWallBlock(H, W, Y, X, field)
	return field
}

// 指定した上下左右マスの位置まで"*"に変更する
func ExecReversiVerticalAndHorizontalWithWallBlock(H, W, Y, X int, field []string) []string {

	// 上下左右の更新終了マス取得
	upper := GetNearestUpperFieldWithWallBlock(Y, X, field)
	lower := GetNearestLowerFieldWithWallBlock(H, Y, X, field)
	left := GetNearestLeftFieldWithWallBlock(Y, X, field)
	right := GetNearestRightFieldWithWallBlock(W, Y, X, field)

	// 上下左右をひっくり返す
	field = ReverseUpperSquare(Y, X, upper, field)
	field = ReverseLowerSquare(H, Y, X, lower, field)
	field = ReverseLeftSquare(Y, X, left, field)
	field = ReverseRightSquare(Y, X, right, field)
	return field
}

// オセロを行う(斜め方向)
func ExecReversiDiagonalWithWallBlock(H, W, Y, X int, field []string) []string {
	// 斜めのの更新終了マス取得
	ulY, ulX := GetNearestDiagonalUpperLeftFieldWithWallBlock(H, W, Y, X, field)
	urY, urX := GetNearestDiagonalUpperRightFieldWithWallBlock(H, W, Y, X, field)
	llY, llX := GetNearestDiagonalLowerLeftFieldWithWallBlock(H, W, Y, X, field)
	lrY, lrX := GetNearestDiagonalLowerRightFieldWithWallBlock(H, W, Y, X, field)

	// マス更新実行
	field = ReverseDiagonalUpperLeftSquare(H, W, Y, X, ulY, ulX, field)
	field = ReverseDiagonalUpperRightSquare(H, W, Y, X, urY, urX, field)
	field = ReverseDiagonalLowerLeftSquare(H, W, Y, X, llY, llX, field)
	field = ReverseDiagonalLowerRightSquare(H, W, Y, X, lrY, lrX, field)
	return field
}

// 課題7
func call_osero7() {
	var H, W, N, Y, X int
	fmt.Scanf("%d %d %d", &H, &W, &N)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)
		field = append(field, str)
	}

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &Y, &X)
		field = ExecReversiWithWallBlock(H, W, Y, X, field)
	}

	for _, v := range field {
		fmt.Println(v)
	}
}

// 課題8
func call_osero8() {
	var H, W, N, Y, X int
	fmt.Scanf("%d %d %d", &H, &W, &N)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)
		field = append(field, str)
	}

	for i := 0; i < N*2; i++ {
		fmt.Scanf("%d %d", &Y, &X)
		// 偶数ターンは先手"A"
		if i%2 == 0 {
			field = ExecReversiBattleMode(H, W, Y, X, field, "A")
		} else {
			field = ExecReversiBattleMode(H, W, Y, X, field, "B")
		}

	}

	for _, v := range field {
		fmt.Println(v)
	}
}

// 課題9
func call_osero9() {
	var H, W, N, n, Y, X int
	fmt.Scanf("%d %d %d %d", &H, &W, &N, &n)

	var field []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)
		field = append(field, str)
	}

	for i := 1; i <= n; i++ {
		var p int
		fmt.Scanf("%d %d %d", &p, &Y, &X)
		turn := strconv.Itoa(p)
		field = ExecReversiBattleMode(H, W, Y, X, field, turn)
	}

	for _, v := range field {
		fmt.Println(v)
	}
}

// オセロロジック実行(壁判定あり)
func ExecReversiBattleMode(H, W, Y, X int, field []string, turn string) []string {
	// 指したマスをひっくり返す
	field = ReverseSetPositionBattleMode(Y, X, field, turn)
	// 縦横をひっくり返す
	field = ExecReversiVerticalAndHorizontalBattleMode(H, W, Y, X, field, turn)
	// 斜めをひっくり返す
	field = ExecReversiDiagonalBattleMode(H, W, Y, X, field, turn)
	return field
}

// 指したマスをひっくり返す
func ReverseSetPositionBattleMode(Y, X int, field []string, turn string) []string {
	str := field[Y]
	str = str[:X] + turn + str[X+1:]
	field[Y] = str
	return field
}

// 上側をひっくり返す
func ReverseUpperSquareBattleMode(Y, X, upper int, field []string, turn string) []string {
	// 更新無し
	if upper < 0 {
		return field
	}

	// マス更新
	for i := Y - 1; i >= upper; i-- {
		str := field[i]
		str = str[:X] + turn + str[X+1:]
		field[i] = str
	}
	return field
}

// 下側をひっくり返す
func ReverseLowerSquareBattleMode(H, Y, X, lower int, field []string, turn string) []string {
	// 更新無し
	if lower < 0 {
		return field
	}

	// マス更新
	for i := Y; i <= lower; i++ {
		str := field[i]
		str = str[:X] + turn + str[X+1:]
		field[i] = str
	}
	return field
}

// 左側をひっくり返す
func ReverseLeftSquareBattleMode(Y, X, left int, field []string, turn string) []string {
	// 更新無し
	if left < 0 {
		return field
	}

	// マス更新
	for i := X - 1; i >= left; i-- {
		str := field[Y]
		str = str[:i] + turn + str[i+1:]
		field[Y] = str
	}
	return field
}

// 右側をひっくり返す
func ReverseRightSquareBattleMode(Y, X, right int, field []string, turn string) []string {
	// 更新無し
	if right < 0 {
		return field
	}

	// マス更新
	for i := X + 1; i <= right; i++ {
		str := field[Y]
		str = str[:i] + turn + str[i+1:]
		field[Y] = str
	}
	return field
}

// 指定した上下左右マスの位置まで"*"に変更する
func ExecReversiVerticalAndHorizontalBattleMode(H, W, Y, X int, field []string, turn string) []string {

	// 上下左右の更新終了マス取得
	upper := GetNearestUpperFieldBattleMode(Y, X, field, turn)
	lower := GetNearestLowerFieldBattleMode(H, Y, X, field, turn)
	left := GetNearestLeftFieldBattleMode(Y, X, field, turn)
	right := GetNearestRightFieldBattleMode(W, Y, X, field, turn)

	// 上下左右をひっくり返す
	field = ReverseUpperSquareBattleMode(Y, X, upper, field, turn)
	field = ReverseLowerSquareBattleMode(H, Y, X, lower, field, turn)
	field = ReverseLeftSquareBattleMode(Y, X, left, field, turn)
	field = ReverseRightSquareBattleMode(Y, X, right, field, turn)
	return field
}

// オセロを行う(斜め方向)
func ExecReversiDiagonalBattleMode(H, W, Y, X int, field []string, turn string) []string {
	// 斜めのの更新終了マス取得
	ulY, ulX := GetNearestDiagonalUpperLeftFieldBattleMode(H, W, Y, X, field, turn)
	urY, urX := GetNearestDiagonalUpperRightFieldBattleMode(H, W, Y, X, field, turn)
	llY, llX := GetNearestDiagonalLowerLeftFieldBattleMode(H, W, Y, X, field, turn)
	lrY, lrX := GetNearestDiagonalLowerRightFieldBattleMode(H, W, Y, X, field, turn)

	// マス更新実行
	field = ReverseDiagonalUpperLeftSquareBattleMode(H, W, Y, X, ulY, ulX, field, turn)
	field = ReverseDiagonalUpperRightSquareBattleMode(H, W, Y, X, urY, urX, field, turn)
	field = ReverseDiagonalLowerLeftSquareBattleMode(H, W, Y, X, llY, llX, field, turn)
	field = ReverseDiagonalLowerRightSquareBattleMode(H, W, Y, X, lrY, lrX, field, turn)
	return field
}

// 上方向で一番近い 自分のマスの座標を返す
// 存在しない または 最上段の場合は -1 を返す
func GetNearestUpperFieldBattleMode(Y, X int, field []string, turn string) int {
	// 最上段範囲
	if Y == 0 {
		return -1
	}

	// 指定座標から 最上段までチェック
	res := -1
	for i := Y - 1; i >= 0; i-- {

		// 列の上マスが自分のマスなら終了
		str := field[i]
		if str[X] == turn[0] {
			res = i
			break
		}

		// 壁か対戦相手の駒にぶつかれば処理終了
		if str[X] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 下方向で一番近い 自分のマスの座標を返す
// 存在しない または 最下段の場合は -1 を返す
func GetNearestLowerFieldBattleMode(H, Y, X int, field []string, turn string) int {
	// 最下段範囲
	if Y == H-1 {
		return -1
	}

	// 指定座標から 最下段までチェック
	res := -1
	for i := Y + 1; i < H; i++ {
		// 列の上マスが"*"なら終了
		str := field[i]
		if str[X] == turn[0] {
			res = i
			break
		}

		// 壁か対戦相手の駒にぶつかれば処理終了
		if str[X] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 左方向で一番近い 自分のマスの座標を返す
// 存在しない または 左端の場合は -1 を返す
func GetNearestLeftFieldBattleMode(Y, X int, field []string, turn string) int {
	// 左端判定
	if X == 0 {
		return -1
	}

	// 指定座標から 左端までチェック
	res := -1
	for i := X - 1; i >= 0; i-- {

		// 行の左マスが"*"なら終了
		str := field[Y]
		if str[i] == turn[0] {
			res = i
			break
		}

		// 壁か対戦相手の駒にぶつかれば処理終了
		if str[i] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 右方向で一番近い 自分のマスの座標を返す
// 存在しない または 右端の場合は -1 を返す
func GetNearestRightFieldBattleMode(W, Y, X int, field []string, turn string) int {
	// 左端判定
	if X == W-1 {
		return -1
	}

	// 指定座標から 右端までチェック
	res := -1
	for i := X + 1; i < W; i++ {
		// 行の右マスが"*"なら終了
		str := field[Y]
		if str[i] == turn[0] {
			res = i
			break
		}

		// 壁か対戦相手の駒にぶつかれば処理終了
		if str[i] == '#' {
			res = -1
			break
		}
	}
	return res
}

// 左上方向で一番近い  自分のママスの座標を返す
// 存在しない または 最上段/左端の場合は -1 を返す
func GetNearestDiagonalUpperLeftFieldBattleMode(H, W, Y, X int, field []string, turn string) (int, int) {
	// 左端か最上段なら終了
	if Y == 0 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y-1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}
		// 自駒が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == turn[0] {
			break
		}

		// 次の座標へ
		nowX--
		nowY--
	}
	return nowY, nowX
}

// 右上方向で一番近い 自分のマスの座標を返す
// 存在しない または 最上段・右端の場合は -1 を返す
func GetNearestDiagonalUpperRightFieldBattleMode(H, W, Y, X int, field []string, turn string) (int, int) {
	// 右端か最上段なら終了
	if Y == 0 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y-1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// 自駒が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == turn[0] {
			break
		}

		// 次の座標へ
		nowX++
		nowY--
	}
	return nowY, nowX
}

// 左下方向で一番近い 自分のマスの座標を返す
// 存在しない または 最下段/左端の場合は -1 を返す
func GetNearestDiagonalLowerLeftFieldBattleMode(H, W, Y, X int, field []string, turn string) (int, int) {
	// 左端か最上段なら終了
	if Y == H-1 || X == 0 {
		return -1, -1
	}
	nowY, nowX := Y+1, X-1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// 自駒が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == turn[0] {
			break
		}

		// 次の座標へ
		nowX--
		nowY++
	}
	return nowY, nowX
}

// 右下方向で一番近い 自分のマスの座標を返す
// 存在しない または 最下段・右端の場合は -1 を返す
func GetNearestDiagonalLowerRightFieldBattleMode(H, W, Y, X int, field []string, turn string) (int, int) {
	// 右端か最上段なら終了
	if Y == H-1 || X == W-1 {
		return -1, -1
	}
	nowY, nowX := Y+1, X+1
	for {
		// 外に出た場合は存在しないとしてマーク
		if isOutside(H, W, nowY, nowX) || field[nowY][nowX] == '#' {
			nowY, nowX = -1, -1
			break
		}

		// 自駒が出ればそこの座標が終点としてマーク
		if field[nowY][nowX] == turn[0] {
			break
		}

		// 次の座標へ
		nowX++
		nowY++
	}
	return nowY, nowX
}

// 左上をひっくり返す
func ReverseDiagonalUpperLeftSquareBattleMode(H, W, Y, X, stopY, stopX int, field []string, turn string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y-1, X-1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + turn + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX--
		nowY--
	}
	return field
}

// 右上をひっくり返す
func ReverseDiagonalUpperRightSquareBattleMode(H, W, Y, X, stopY, stopX int, field []string, turn string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y-1, X+1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + turn + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX++
		nowY--
	}
	return field
}

// 左下をひっくり返す
func ReverseDiagonalLowerLeftSquareBattleMode(H, W, Y, X, stopY, stopX int, field []string, turn string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y+1, X-1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + turn + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX--
		nowY++
	}
	return field
}

// 右下をひっくり返す
func ReverseDiagonalLowerRightSquareBattleMode(H, W, Y, X, stopY, stopX int, field []string, turn string) []string {
	// 更新無し
	if stopY < 0 || stopX < 0 {
		return field
	}

	// 外に出る or 指定マスまではひっくり返す
	nowY, nowX := Y+1, X+1
	for {
		if isOutside(H, W, nowY, nowX) || (nowY == stopY && nowX == stopX) {
			break
		}
		// マスを更新
		point := field[nowY]
		point = point[:nowX] + turn + point[nowX+1:]
		field[nowY] = point
		// 次の座標へ
		nowX++
		nowY++
	}
	return field
}
