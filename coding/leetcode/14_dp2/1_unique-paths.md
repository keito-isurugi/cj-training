# Unique Paths (Medium)

## 問題内容

`m x n` のグリッドがあり、任意の時点で**下**または**右**にのみ移動できる。

整数 `m` と `n` が与えられたとき、グリッドの左上隅（`grid[0][0]`）から右下隅（`grid[m - 1][n - 1]`）までの**ユニークなパスの数**を返す。

出力は **32ビット整数**に収まると仮定してよい。

### 例

```
Input: m = 3, n = 6
Output: 21
```

```
Input: m = 3, n = 3
Output: 6
```

### 制約

- `1 <= m, n <= 100`

## ソースコード

```go
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }

    for i := m - 2; i >= 0; i-- {
        for j := n - 2; j >= 0; j-- {
            dp[j] += dp[j+1]
        }
    }

    return dp[0]
}
```

## アルゴリズムなど解説

### 基本戦略

各セルから目的地に到達するパスの数は：
- **下のセル**からのパス数
- **右のセル**からのパス数
- の**合計**

これは典型的な **2次元DP** 問題。

### 核心となる洞察

```
dp[i][j] = dp[i+1][j] + dp[i][j+1]
```

この漸化式は、各セルへのパス数が下と右のセルのパス数の和であることを示す。

### 動作の仕組み

1. **初期化**
   ```go
   dp := make([]int, n)
   for i := range dp {
       dp[i] = 1
   }
   ```
   - 最下行は右にしか移動できないため、すべて1通り

2. **空間最適化されたDP**
   ```go
   for i := m - 2; i >= 0; i-- {
       for j := n - 2; j >= 0; j-- {
           dp[j] += dp[j+1]
       }
   }
   ```
   - 各行は下の行のみに依存するため、1次元配列で十分
   - 右から左に更新することで、正しい値を参照

### 具体例

```
m = 3, n = 3

初期状態（最下行）:
dp = [1, 1, 1]

i = 1 の処理後:
dp = [3, 2, 1]
  j=1: dp[1] = 1 + 1 = 2
  j=0: dp[0] = 1 + 2 = 3

i = 0 の処理後:
dp = [6, 3, 1]
  j=1: dp[1] = 2 + 1 = 3
  j=0: dp[0] = 3 + 3 = 6

結果: 6
```

### グリッドの視覚化

```
3x3 グリッドの各セルからのパス数:

+---+---+---+
| 6 | 3 | 1 |
+---+---+---+
| 3 | 2 | 1 |
+---+---+---+
| 1 | 1 | 1 |
+---+---+---+

左上(0,0)から右下(2,2)へのパス数 = 6
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n) | 全セルを1回ずつ処理 |
| 空間計算量 | O(n) | 1行分の配列のみ使用 |

### 別解：2次元DP配列

```go
func uniquePaths(m int, n int) int {
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    dp[m-1][n-1] = 1

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            dp[i][j] += dp[i+1][j] + dp[i][j+1]
        }
    }

    return dp[0][0]
}
```

### 別解：再帰 + メモ化

```go
func uniquePaths(m int, n int) int {
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == m-1 && j == n-1 {
            return 1
        }
        if i >= m || j >= n {
            return 0
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }

        memo[i][j] = dfs(i, j+1) + dfs(i+1, j)
        return memo[i][j]
    }

    return dfs(0, 0)
}
```

### 別解：数学的アプローチ

右に `n-1` 回、下に `m-1` 回移動する必要がある。
合計 `m+n-2` 回の移動から `n-1` 回（または `m-1` 回）を選ぶ組み合わせ。

```go
func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {
        return 1
    }
    if m < n {
        m, n = n, m
    }

    res := 1
    j := 1
    for i := m; i < m+n-1; i++ {
        res *= i
        res /= j
        j++
    }

    return res
}
```

**計算量**：時間 O(min(m, n))、空間 O(1)

### DPパターン

この問題は **グリッドDP** パターン：
- 状態遷移: `dp[i][j] = dp[i+1][j] + dp[i][j+1]`
- 下と右のセルに依存
- 空間最適化で O(n) 可能

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Unique Paths** | グリッドDP |
| Unique Paths II | 障害物付きグリッドDP |
| Minimum Path Sum | グリッドDP（最小コスト） |
| Climbing Stairs | フィボナッチ型DP |
