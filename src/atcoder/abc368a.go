package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	// カードの山を格納するスライス
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// 山の下から K 枚のカードを取り出し、順序を保ったまま山の一番上に乗せる
	result := append(A[N-K:], A[:N-K]...)

	// 結果を出力
	for i, card := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(card)
	}
	fmt.Println()
}
