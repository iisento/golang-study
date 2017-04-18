package main

import "golang.org/x/tour/pic"

// Pic ...
func Pic(dx, dy int) (image [][]uint8) {
	image = make([][]uint8, dy) // スライスを要素に持てる長さ dy のスライスを生成
	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx) // 要素のスライスを長さ dx で生成して代入する
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			image[x][y] = uint8(x * y) // 画像を入れる
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
