package main

import "fmt"

// x,yで受け取った座標から、今の向きと左右どちらかを受け取り
// その方向に対して移動した座標を返す
func MukiIdou(x, y int, leftOrRight, muki string) (int, int) {

	var i, j int

	switch muki {
	case "N":
		i, j = KitaMukiIdou(x, y, leftOrRight)
	case "S":
		i, j = MinamiMukiIdou(x, y, leftOrRight)
	case "E":
		i, j = HigashiMukiIdou(x, y, leftOrRight)
	case "W":
		i, j = NishiMukiIdou(x, y, leftOrRight)
	}
	return i, j
}

/* 北を向いてる状態で、与えられた座標から移動した座標を返す */
func KitaMukiIdou(x, y int, leftOrRight string) (int, int) {
	if leftOrRight == "R" {
		return x + 1, y
	} else {
		return x - 1, y
	}
}

/* 南を向いてる状態で、与えられた座標から移動した座標を返す */
func MinamiMukiIdou(x, y int, leftOrRight string) (int, int) {
	if leftOrRight == "R" {
		return x - 1, y
	} else {
		return x + 1, y
	}
}

/* 東向いてる状態で、与えられた座標から移動した座標を返す */
func HigashiMukiIdou(x, y int, leftOrRight string) (int, int) {
	if leftOrRight == "R" {
		return x, y + 1
	} else {
		return x, y - 1
	}
}

/* 西向いてる状態で、与えられた座標から移動した座標を返す */
func NishiMukiIdou(x, y int, leftOrRight string) (int, int) {
	if leftOrRight == "R" {
		return x, y - 1
	} else {
		return x, y + 1
	}
}

///-----

/* スタート位置から時計回りにhosu歩移動した時の座標を返す */
func UzumakiIdouTokeiMawari(x, y, hosu int) (int, int) {

	// 与えられた歩数が時計回りに何週するかを求める
	// ここでの1週は原点である
	// ただし総歩数に原点は含まれないため歩数は+1した状態で始める
	// (算出ロジックでは原点を含んだ値を出すため)
	syu, sum := 0, 0
	for sum < hosu+1 {
		sum += GaisyuMasuCalc(syu + 1)
		syu++
	}
	// 算術のタイミングで原点から数えて1週目, 2週目としたいので -1して調整
	syu -= 1
	// 	fmt.Println("何週するか(原点含む): ", syu)

	/// n+1 週目で消費される歩数を算出する
	nokori := GaisyuNokoriHosu(syu, hosu)
	// 	fmt.Println("残り歩数: ", nokori)

	// (n, n)の位置から次のように移動する
	// 通常の数学的なマスで一旦考えるとこう。
	// 1. 下方向に 2*nマス ( n, -n)
	// 2. 左方向に 2*nマス (-n, -n)
	// 3. 上方向に 2*nマス (-n,  n)
	// 4. 右方向に 2*nマス ( n,  n) ここで開始位置に戻る
	// 問題文で考えるとこう。
	// 開始位置: (n, -n)の位置から次のように移動する
	// 1. 下方向に 2*nマス ゴール:( n,  n)
	// 2. 左方向に 2*nマス ゴール:(-n,  n)
	// 3. 上方向に 2*nマス ゴール:(-n, -n)
	// 4. 右方向に 2*nマス ゴール:( n, -n) ここで開始位置に戻る
	resX, resY := GoalMasu(x, y, syu, nokori)
	return resX, resY
}

// 残り歩数から移動先のマスを計算する
func GoalMasu(x, y, n, nokori int) (int, int) {
	startX, startY, move, direction := StartMasu(x, y, n, nokori)

	// fmt.Println("開始位置: ", startX, startY)
	// fmt.Println("何歩動く？:", move)
	// fmt.Println("移動方向は？:", direction)

	switch direction {

	case "North": // 上に move マス移動
		startY -= move

	case "South": // 下に move マス移動
		startY += move

	case "West": // 左に move マス移動
		startX -= move

	case "East": // 右に move マス移動
		startX += move
	}

	return startX, startY

}

// スタート位置を4隅のどこから始めるかを求める。
// またその位置から、どの方向に何歩歩くかも返す
// 残り歩数から 2*nマスを引いて行き、マイナスにならないぎりぎりの位置をスタート位置として出す
func StartMasu(x, y, n, nokori int) (int, int, int, string) {
	resX, resY := x, y
	// 1回目、(残り歩数 - 2n) < 0なら (n, -n)の位置からスタート、かつ下へ
	move1 := nokori - 2*n
	if move1 < 0 {
		resX += n
		resY -= n
		return resX, resY, nokori, "South"
	}
	// 2回目 (残り歩数 - 2n*2 ) < 0なら (n, n)の位置からスタート、かつ左へ
	move2 := nokori - 2*n*2
	if move2 < 0 {
		resX += n
		resY += n
		return resX, resY, move1, "West"
	}
	// 3回目 (残り歩数 - 2n*3 ) < 0なら (-n, n)の位置からスタート、かつ上へ
	move3 := nokori - 2*n*3
	if move3 < 0 {
		resX -= n
		resY += n
		return resX, resY, move2, "North"
	}
	// 4回目 (残り歩数 - 2n*4 ) < 0なら (-n, -n)の位置からスタート、かつ右へ
	move4 := nokori - 2*n*4
	if move4 < 0 {
		resX -= n
		resY -= n
		return resX, resY, move3, "East"
	}

	// 4回目のチェックも抜けた場合、move4 == 0なので、(n, -n)からスタートかつ移動無しを返す
	resX += n
	resY -= n
	return resX, resY, move4, "None"
}

// 原点を1週目とした時、最後の週で移動する残りの歩数を求める。
// (syu - 1)でマイナスにならない省略できる周回分の歩数を求める
func GaisyuNokoriHosu(syu, hosu int) int {
	sum := 0
	hosu += 1 // 原点の分のマスを追加
	for i := 1; i <= syu; i++ {
		sum += GaisyuMasuCalc(i)
	}

	return hosu - sum
}

// (x,y)を原点とし、 n週目で使うマスの数を計算する 2週は8, 3週は 16、4週は24...
// 原点は1週目として含む
// a_1 = 1, a_2 = 8, a_3 = 16
// a_n = {1+ 2*(n-1)}^2 - sum(a_n-1..a_1)
func GaisyuMasuCalc(n int) int {
	if n == 1 {
		return 1
	}
	return (n - 1) * 8
}

// -----------[new]--------

// x,yで受け取った座標から、今の向きと左右どちらかを受け取り
// その方向に対して移動した座標と、移動した後向いている向きを返す
func MukiIdouAndMukiUpdate(x, y int, leftOrRight, muki string) (int, int, string) {
	// 移動後の座標を求める
	resX, resY := MukiIdou(x, y, leftOrRight, muki)

	// 新しい向きを求める
	newMuki := GetNewMuki(leftOrRight, muki)

	return resX, resY, newMuki
}

// 今向いている方角から左右に移動した時、向いている方角を求める
func GetNewMuki(leftOrRight, muki string) string {
	newMuki := ""

	switch muki {
	case "N": // 北向き
		if leftOrRight == "R" {
			newMuki = "E"
		} else {
			newMuki = "W"
		}
	case "S": // 南向き
		if leftOrRight == "R" {
			newMuki = "W"
		} else {
			newMuki = "E"
		}
	case "W": // 西向き
		if leftOrRight == "R" {
			newMuki = "N"
		} else {
			newMuki = "S"
		}
	case "E": // 東向き
		if leftOrRight == "R" {
			newMuki = "S"
		} else {
			newMuki = "N"
		}
	}

	return newMuki
}

// -----12/9
// 指定した方角に移動できるか判定する
func IsMovableToDirection(ary []string, x, y, H, W int, muki string) bool {
	// 移動先のマス用変数
	var afterX, afterY int = x, y

	// 向きで変わる移動先の座標を設定
	switch muki {
	case "E": // 東
		afterX += 1
	case "W": //西
		afterX -= 1
	case "N": //北
		afterY -= 1
	case "S": //南
		afterY += 1
	}

	// 壁なら移動できないので結果を反転
	return IsWall(ary, x, y, H, W)
}

// 与えられた座標が壁か判定
func IsWall(ary []string, x, y, H, W int) bool {
	// 座標が最上段、左端を超えていれば壁判定
	if x < 0 || y < 0 {
		return true
	}

	// 座標が最下段、右端を超えていれば壁判定
	// 縦: H, 横: W で定義されているが座標はindex番号で0,0開始のため-1した値で判定
	if x > W-1 || y > H-1 {
		return true
	}

	// 範囲内に収まっていれば壁判定の'#'か判定する
	return ary[y][x] == '#'
}

/// ----

// x,yで受け取った座標から、今の向きと左右どちらかを受け取り
// その方向に対して移動した座標が壁とぶつかるかどうかを判定する
func IsMovableToSpin(ary []string, x, y, H, W int, leftOrRight, muki string) bool {

	var afterX, afterY int

	switch muki {
	case "N":
		afterX, afterY = KitaMukiIdou(x, y, leftOrRight)
	case "S":
		afterX, afterY = MinamiMukiIdou(x, y, leftOrRight)
	case "E":
		afterX, afterY = HigashiMukiIdou(x, y, leftOrRight)
	case "W":
		afterX, afterY = NishiMukiIdou(x, y, leftOrRight)
	}

	return !IsWall(ary, afterX, afterY, H, W)
}

// 回転しながら動く時、壁に判定があれば”Stop”を返す
func IsMovableToSpin2(ary []string, x, y, H, W int, leftOrRight, muki string) (int, int, string) {

	afterX, afterY, newMuki := MukiIdouAndMukiUpdate(x, y, leftOrRight, muki)
	fmt.Println("direction:", leftOrRight, muki)
	fmt.Println("Turned Pos:", x, y, newMuki)
	fmt.Println("After  pos:", afterX, afterY, newMuki)

	if !IsWall(ary, afterX, afterY, H, W) {
		return afterX, afterY, newMuki
	} else {
		return -1, -1, "Stop"
	}
}

// ---

// 壁に当たるまでまっすぐ進む。
// 突進を終えた後の座標と向きを返す
func TyototsuMoushin(ary []string, x, y, H, W int, muki, leftOrRight string) (int, int, string) {
	var resX, resY int
	var resMuki string

	switch muki {
	case "N": // 北向き
		if leftOrRight == "R" {
			resX, resY = RushEast(ary, x, y, H, W)
			resMuki = "E"
		} else {
			resX, resY = RushWest(ary, x, y, H, W)
			resMuki = "W"
		}
	case "S": // 南向き
		if leftOrRight == "R" {
			resX, resY = RushWest(ary, x, y, H, W)
			resMuki = "W"
		} else {
			resX, resY = RushEast(ary, x, y, H, W)
			resMuki = "E"
		}
	case "W": // 西向き
		if leftOrRight == "R" {
			resX, resY = RushNorth(ary, x, y, H, W)
			resMuki = "N"
		} else {
			resX, resY = RushSouth(ary, x, y, H, W)
			resMuki = "S"
		}
	case "E": // 東向き
		if leftOrRight == "R" {
			resX, resY = RushSouth(ary, x, y, H, W)
			resMuki = "S"
		} else {
			resX, resY = RushNorth(ary, x, y, H, W)
			resMuki = "N"
		}
	}

	return resX, resY, resMuki
}

// 東方向に突撃。壁にぶつかるまでループ。
func RushEast(ary []string, x, y, H, W int) (int, int) {
	nowX, nowY := x, y
	// 右端 or 壁のどちらかを満たすまで x座標を + 1
	for nowX+1 != W || !IsWall(ary, x, y, H, W) {
		nowX += 1
	}

	return nowX, nowY
}

// 西方向に突撃。壁にぶつかるまでループ。
func RushWest(ary []string, x, y, H, W int) (int, int) {
	nowX, nowY := x, y
	// 左端 or 壁のどちらかを満たすまで x座標を - 1
	for nowX-1 != 0 || !IsWall(ary, x, y, H, W) {
		nowX -= 1
	}
	return nowX, nowY
}

// 北方向に突撃。壁にぶつかるまでループ。
func RushNorth(ary []string, x, y, H, W int) (int, int) {
	nowX, nowY := x, y
	// 上端 or 壁のどちらかを満たすまで y座標を - 1
	for nowY-1 != 0 || !IsWall(ary, x, y, H, W) {
		nowX += 1
	}
	return nowX, nowY
}

// 南方向に突撃。壁にぶつかるまでループ。
func RushSouth(ary []string, x, y, H, W int) (int, int) {
	nowX, nowY := x, y
	// 下端 or 壁のどちらかを満たすまで y座標を + 1
	for nowY+1 != H || !IsWall(ary, x, y, H, W) {
		nowX += 1
	}
	return nowX, nowY
}

// 回転した後、歩数分直進する
// 壁にぶつかったかどうかと移動後の座標を返す
func TyototsuMoushinWithHosu(ary []string, x, y, H, W, hosu int, muki, leftOrRight string) (int, int, string, bool) {
	var resX, resY int
	var resMuki string
	var isCrash bool
	switch muki {
	case "N": // 北向き
		if leftOrRight == "R" {
			resX, resY, isCrash = RushEastWithHosu(ary, x, y, H, W, hosu)
			resMuki = "E"
		} else {
			resX, resY, isCrash = RushWestWithHosu(ary, x, y, H, W, hosu)
			resMuki = "W"
		}
	case "S": // 南向き
		if leftOrRight == "R" {
			resX, resY, isCrash = RushWestWithHosu(ary, x, y, H, W, hosu)
			resMuki = "W"
		} else {
			resX, resY, isCrash = RushEastWithHosu(ary, x, y, H, W, hosu)
			resMuki = "E"
		}
	case "W": // 西向き
		if leftOrRight == "R" {
			resX, resY, isCrash = RushNorthWithHosu(ary, x, y, H, W, hosu)
			resMuki = "N"
		} else {
			resX, resY, isCrash = RushSouthWithHosu(ary, x, y, H, W, hosu)
			resMuki = "S"
		}
	case "E": // 東向き
		if leftOrRight == "R" {
			resX, resY, isCrash = RushSouthWithHosu(ary, x, y, H, W, hosu)
			resMuki = "S"
		} else {
			resX, resY, isCrash = RushNorthWithHosu(ary, x, y, H, W, hosu)
			resMuki = "N"
		}
	}

	return resX, resY, resMuki, isCrash
}

// 東方向に突撃。壁にぶつかるまで指定歩数分
// ただし壁にぶつかった場合は以後の処理を打ち切り
func RushEastWithHosu(ary []string, x, y, H, W, hosu int) (int, int, bool) {
	nowX, nowY := x, y
	isCrash := false
	// 歩数分歩く or 壁のどちらかを満たすまで x座標を + 1
	for i := 0; i < hosu; i++ {
		nowX += 1

		// 壁にぶつかったら終了
		if IsWall(ary, nowX, nowY, H, W) {
			// 壁にぶつかったので１マス戻る
			nowX -= 1
			isCrash = true
			break
		}
	}
	return nowX, nowY, isCrash
}

// 西方向に突撃。壁にぶつかるまでループ。
func RushWestWithHosu(ary []string, x, y, H, W, hosu int) (int, int, bool) {
	nowX, nowY := x, y
	isCrash := false
	// 歩数分歩く or 壁のどちらかを満たすまで x座標を - 1
	for i := 0; i < hosu; i++ {
		nowX -= 1

		// 壁にぶつかったら終了
		if IsWall(ary, nowX, nowY, H, W) {
			// 壁にぶつかったので１マス戻る
			nowX += 1
			isCrash = true
			break
		}
	}
	return nowX, nowY, isCrash
}

// 北方向に突撃。壁にぶつかるまでループ。
func RushNorthWithHosu(ary []string, x, y, H, W, hosu int) (int, int, bool) {
	nowX, nowY := x, y
	isCrash := false
	// 歩数分歩く or 壁のどちらかを満たすまで y座標を - 1
	for i := 0; i < hosu; i++ {
		nowY -= 1

		// 壁にぶつかったら終了
		if IsWall(ary, nowX, nowY, H, W) {
			// 壁にぶつかったので１マス戻る
			nowY += 1
			isCrash = true
			break
		}
	}
	return nowX, nowY, isCrash
}

// 南方向に突撃。壁にぶつかるまでループ。
func RushSouthWithHosu(ary []string, x, y, H, W, hosu int) (int, int, bool) {
	nowX, nowY := x, y

	isCrash := false
	// 歩数分歩く or 壁のどちらかを満たすまで y座標を + 1
	for i := 0; i < hosu; i++ {
		nowY += 1

		// 壁にぶつかったら終了
		if IsWall(ary, nowX, nowY, H, W) {
			// 壁にぶつかったので１マス戻る
			nowY -= 1
			isCrash = true
			break
		}
	}
	return nowX, nowY, isCrash
}

// x,yで受け取った座標から、今の向きと左右どちらかを受け取り
// その方向に対して移動した座標と、移動した後向いている向きを返す
func MukiIdouAndMukiUpdate2(x, y, hosu int, leftOrRight, muki string) (int, int, string) {
	// 移動後の座標を求める
	nowX, nowY := x, y

	for i := 0; i < hosu; i++ {
		nowX, nowY = MukiIdou(x, y, leftOrRight, muki)
	}

	resX, resY := nowX, nowY

	// 新しい向きを求める
	newMuki := GetNewMuki(leftOrRight, muki)

	return resX, resY, newMuki
}

// 回転しながら動く時、壁に判定があれば”Stop”を返す
// また旋回後、与えられた歩数分だけ進
func IsMovableToSpin3(ary, ary2 []string, x, y, H, W, hosu int, leftOrRight, muki string) (int, int, string, []string) {

	//fmt.Println("-- IsMovableToSpin3 start --")

	nowAry := ary2
	nowX, nowY := x, y
	//fmt.Println("現在の座標:", nowX, nowY)
	// 開始位置を踏んだ判定に
	str := nowAry[nowY]
	str2 := str[0:nowX] + "*" + str[nowX+1:]
	nowAry[nowY] = str2

	// 次の向きを求める
	newMuki := GetNewMuki(leftOrRight, muki)

	// 与えられた歩数分
	for i := 0; i < hosu; i++ {
		//fmt.Println(string(i+1) + "歩目")
		nowX, nowY = MukiIdou(nowX, nowY, leftOrRight, muki)

		//fmt.Println("現在の座標:", nowX, nowY)
		if !IsWall(ary, nowX, nowY, H, W) {
			// 壁じゃなければ踏破マスを更新
			str = nowAry[nowY]
			str2 = str[0:nowX] + "*" + str[nowX+1:]
			nowAry[nowY] = str2
		} else {
			nowX, nowY, newMuki = -1, -1, "Stop"
			break
		}
	}

	//fmt.Println("-- IsMovableToSpin3 end --")
	return nowX, nowY, newMuki, nowAry
}

// 蛇のように這いまわる課題要構造体
type SnakeAction struct {
	time  int
	angle string
}

// 蛇のように這いまわる課題
func SnakeSpinOnTime() {
	// H W sy sx N
	var H, W, sy, sx, N int
	fmt.Scanf("%d %d %d %d %d", &H, &W, &sy, &sx, &N)

	// エリアをスキャンで受取
	var ary []string
	for i := 0; i < H; i++ {
		var str string
		fmt.Scan(&str)
		ary = append(ary, str)
	}

	// 方向転換のタイミングを取得
	var actions []SnakeAction
	for i := 0; i < N; i++ {
		var time int
		var angle string
		fmt.Scanf("%d %s", &time, &angle)

		act := SnakeAction{time, angle}
		actions = append(actions, act)
	}

	// 100秒間稼働
	actNo := 0
	actionTime := actions[0].time
	actionAngle := actions[0].angle
	nowX, nowY, nowMuki := sx, sy, "N"

	for i := 0; i < 100; i++ {

		// 方向転換のタイミングになったら、Spinする
		if actionTime == i {

			nowX, nowY, nowMuki = IsMovableToSpin2(ary, nowX, nowY, H, W, actionAngle, nowMuki)
			// 移動後次の方向転換のタイミングを取得

			if actNo+1 < N { // 入力アクションを超えなければ次の要素取得
				actNo++
				actionTime = actions[actNo].time
				actionAngle = actions[actNo].angle
			}

		} else {
			// そうでなければ1歩直進移動する
			isCrash := false
			nowX, nowY, isCrash = TyototsuMoushinWithHosu2(ary, nowX, nowY, H, W, 1, nowMuki)

			// クラッシュしていれば失敗フラグを付与
			if isCrash {
				nowX, nowY, nowMuki = -1, -1, "Stop"
			}
		}

		// クラッシュしてればループ終了
		if nowMuki == "Stop" {
			fmt.Println("Stop")
			break
		}

		// 移動後の座標出力
		fmt.Println(nowY, nowX)

	}

}

// 歩数分直進する
// 壁にぶつかったかどうかと移動後の座標を返す
func TyototsuMoushinWithHosu2(ary []string, x, y, H, W, hosu int, muki string) (int, int, bool) {
	var resX, resY int
	var isCrash bool
	switch muki {
	case "N": // 北向き
		resX, resY, isCrash = RushNorthWithHosu(ary, x, y, H, W, hosu)

	case "S": // 南向き
		resX, resY, isCrash = RushSouthWithHosu(ary, x, y, H, W, hosu)

	case "W": // 西向き
		resX, resY, isCrash = RushWestWithHosu(ary, x, y, H, W, hosu)

	case "E": // 東向き
		resX, resY, isCrash = RushEastWithHosu(ary, x, y, H, W, hosu)
	}

	return resX, resY, isCrash
}

func SnakeSpinOnTime2() {
	// H W sy sx N
	var H, W, sy, sx, N int = 29, 40, 20, 1, 3
	//var H, W, sy, sx, N int = 15, 42, 11, 32, 4
	//fmt.Scanf("%d %d %d %d %d", &H, &W, &sy, &sx, &N)

	// エリアをスキャンで受取
	var ary []string = []string{
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#........######################.........",
		"#........#....................#.........",
		"#........#....................#.........",
		"#........#....................#.........",
		"#........#....................#.........",
		"#........#....................#.........",
		"#........#....................#.........",
		"#........######################.........",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",
		"#.......................................",

		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "..........................................",
		// "...............................##.........",
		// "................................##........",
		// "..................................##......",
	}
	// for i := 0; i < H; i++ {
	// 	var str string
	// 	fmt.Scan(&str)
	// 	ary = append(ary, str)
	// }

	// 方向転換のタイミングを取得
	var actions []SnakeAction
	// for i := 0; i < N; i++ {
	// 	var time int
	// 	var angle string
	// 	fmt.Scanf("%d %s", &time, &angle)

	// 	act := SnakeAction{time, angle}
	// 	actions = append(actions, act)
	// }
	act := SnakeAction{20, "R"}
	actions = append(actions, act)
	act = SnakeAction{58, "R"}
	actions = append(actions, act)
	act = SnakeAction{72, "R"}
	actions = append(actions, act)

	// act := SnakeAction{10, "L"}
	// actions = append(actions, act)
	// act = SnakeAction{40, "L"}
	// actions = append(actions, act)
	// act = SnakeAction{52, "L"}
	// actions = append(actions, act)
	// act = SnakeAction{80, "L"}
	// actions = append(actions, act)

	// 100秒間稼働
	actNo := 0
	actionTime := actions[0].time
	actionAngle := actions[0].angle
	nowX, nowY, nowMuki := sx, sy, "N"

	fmt.Println("start  Pos:", nowX, nowY)

	ary2 := ary
	for i := 0; i < 100; i++ {

		// 現在位置をマーク
		str := ary2[nowY]
		str2 := str[0:nowX] + "*" + str[nowX+1:]
		ary2[nowY] = str2

		// 方向転換のタイミングになったら、Spinする
		if actionTime == i {
			fmt.Println("turend :", i)
			// ターン位置をマーク
			str := ary2[nowY]
			str2 := str[0:nowX] + "%" + str[nowX+1:]
			ary2[nowY] = str2

			nowX, nowY, nowMuki = IsMovableToSpin2(ary, nowX, nowY, H, W, actionAngle, nowMuki)
			// 移動後次の方向転換のタイミングを取得

			if actNo+1 < N { // 入力アクションを超えなければ次の要素取得
				actNo++
				actionTime = actions[actNo].time
				actionAngle = actions[actNo].angle

				fmt.Println("次回:", actionTime, ",今:", i)
			}

		} else {
			// そうでなければ1歩直進移動する
			isCrash := false
			nowX, nowY, isCrash = TyototsuMoushinWithHosu2(ary, nowX, nowY, H, W, 1, nowMuki)

			// クラッシュしていれば失敗フラグを付与
			if isCrash {
				nowX, nowY, nowMuki = -1, -1, "Stop"
			}
		}

		// クラッシュしてればループ終了
		if nowMuki == "Stop" {
			break
		}

		// 既に通った道でもループ終了
		if ary2[nowY][nowX] == '*' {
			break
		}
	}

	// 移動ルートを出力
	for _, v := range ary2 {
		fmt.Println(v)
	}
}
