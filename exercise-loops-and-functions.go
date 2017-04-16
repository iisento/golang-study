package main

import (
	"fmt"
	"math"
)

// Newton法で平方根を計算する補正を行う関数
// 引数
// 第1引数：計算したい平方根の引数
// 第2引数：補正したい近似値
// 戻り値
// 第1戻り値：補正値
// 第2戻り値：補正値の一つ前の近似からの絶対誤差
func NewtonSqrtCorrection(sqrt_arg, before float64) (after, abserr float64) {
	after = 0.5 * (before + sqrt_arg/before)
	abserr = math.Abs(after - before)
	return after, abserr
}

// Newton法で平方根を計算する関数
// 第1引数：計算したい平方根の引数
// 第2引数：Newton法で計算するための初期値
// 第3引数：一つ前の補正からの許容する絶対誤差の上限
// 第4引数：計算のループ回数の上限
// 第5引数：途中経過を出力するか．trueなら出力する．falseなら出力しない．
func Sqrt(sqrt_arg, init, eps float64, loop_max int, output bool) float64 {
	var before, after float64
	var abserr float64
	var loop_count int = 0
	before = init
	for {
		loop_count++
		after, abserr = NewtonSqrtCorrection(sqrt_arg, before)
		if output {
			fmt.Printf("loop[%02d] : sqrt(2) = %.20f : abserr = %.20f\n", loop_count, after, abserr)
		}
		if abserr < eps {
			break
		} else if loop_count >= loop_max {
			fmt.Printf("上限ループ回数 %d 以内に指定された絶対誤差 %1.1e 以内で収束しませんでした．\n", loop_count, eps)
			break
		}
		before = after
	}
	return after
}

func main() {
	fmt.Printf("Newton法：%.20f\n", Sqrt(2.0, 1.0, 1.0e-15, 20, true))
	fmt.Printf("math PKG：%.20f\n", math.Sqrt(2.0))
}
