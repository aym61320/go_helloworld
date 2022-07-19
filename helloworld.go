package main

import "fmt"

func main() {

	msg := "Hello,World!"
	printer(msg)

	fmt.Println(returnFavorites())

	fmt.Println(returnCompliment("Hairstyle"))

	arrayPrinter()

	SlicePrinter()
}

// Hello,World
func printer(msg string) {
	fmt.Println(msg)
}

// 複数return
func returnFavorites() (string, string) {
	return "EBI!", "Boston terrier!"
}

//変数return
func returnCompliment(word string) (msg string) {
	msg = "Good " + word
	return
}

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

//スライス
func SlicePrinter() {
	// make関数でスライスを作成する
	s1 := make([]int, 3) // [0 0 0]

	// 値を割り当てたスライスを作成する
	s2 := []int{1, 3, 5} // 配列の宣言と似ている
	fmt.Println(s1)
	fmt.Println(s2)

	// appendでスライスの末尾に要素を追加する
	s3 := append(s2, 8, 2, 10)
	fmt.Println(s3)

	// copyでスライスをコピーする（戻り値は要素数）
	t := make([]int, len(s3))
	s4 := copy(t, s3)
	fmt.Println(s4)
}
