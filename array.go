package main

import "fmt"

//配列
func arrayPrinter() {
	// 宣言と代入を分ける
	var a [5]int // a[0] - a[4]

	a[2] = 3
	a[4] = 10

	fmt.Println(a)

	// 宣言と代入を同時に行う
	b := [3]int{1, 3, 6}
	fmt.Println(b)

	// 配列の個数が未定の場合
	c := [...]int{2, 4, 7, 5, 5}
	fmt.Println(c)
}
