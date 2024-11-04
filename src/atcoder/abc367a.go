package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	// 高橋君が起きている時間帯に愛を叫べるかを判定
	// B時に就寝し、C時に起床
	// A時に起きているか？
	// 寝ている時間：B ~ C
	// B<Cの時はそのまま. B~CにＡが入っていたらNG
	if B < C {
		if B < A && A < C {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	} else {
		// B>Cのときは日にちをまたいでいる. <Cと>Bがアウト
		if C < A && A < B {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
