// 入力
// まず、pacage main
package main
// fmtをインポート
import fmt

// N
// K
var N int
var S string
fmt.Scan(&N)
fmt.Scan(&S)

// N K
var N, K int
fmt.Scan(&N, &K)
// A1 A2 ... AN
A := make([]int, N) // まず整数Nこの配列を作る
for i := 0; i < N; i++ { // 読み取ってA[i]に格納していく
	fmt.Scan(&A[i])
}




