package slice

import "fmt"

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
