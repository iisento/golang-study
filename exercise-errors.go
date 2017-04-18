package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt ...
// float64が平方根を取れるかのエラーを扱えるようにした型
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

// NewtonSqrtCorrection ...
// Newton法で平方根を計算する補正を行う関数
// 引数
// 第1引数：計算したい平方根の引数
// 第2引数：補正したい近似値
// 戻り値
// 第1戻り値：補正値
// 第2戻り値：補正値の一つ前の近似からの絶対誤差
func NewtonSqrtCorrection(sqrtArg, before float64) (after, abserr float64) {
	after = 0.5 * (before + sqrtArg/before)
	abserr = math.Abs(after - before)
	return after, abserr
}

// Sqrt ...
// Newton法で平方根を計算する関数
// 第1引数：計算したい平方根の引数
// 第2引数：Newton法で計算するための初期値
// 第3引数：一つ前の補正からの許容する絶対誤差の上限
// 第4引数：計算のループ回数の上限
// 第5引数：途中経過を出力するか．trueなら出力する．falseなら出力しない．
func Sqrt(sqrtArg, init, eps float64, loopMax int, output bool) (float64, error) {
	var before, after float64
	var abserr float64
	var loopCount int
	if sqrtArg < 0 {
		return 0.0, ErrNegativeSqrt(sqrtArg)
	}
	before = init
	for {
		loopCount++
		after, abserr = NewtonSqrtCorrection(sqrtArg, before)
		if output {
			fmt.Printf("loop[%02d] : sqrt(2) = %.20f : abserr = %.20f\n", loopCount, after, abserr)
		}
		if abserr < eps {
			break
		} else if loopCount >= loopMax {
			fmt.Printf("上限ループ回数 %d 以内に指定された絶対誤差 %1.1e 以内で収束しませんでした．\n", loopCount, eps)
			break
		}
		before = after
	}

	return after, nil
}

func main() {
	fmt.Println(Sqrt(-2.0, 1.0, 1.0e-15, 20, false))
	fmt.Println(Sqrt(+2.0, 1.0, 1.0e-15, 20, false))
}
