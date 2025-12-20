# Climbing Stairs (Easy)

## 問題内容

階段の頂上に到達するためのステップ数を表す整数 `n` が与えられる。1回に `1` または `2` ステップ登ることができる。

頂上に到達するための**異なる方法の数**を返す。

### 例

```
Input: n = 2
Output: 2
```
説明:
1. `1 + 1 = 2`
2. `2 = 2`

```
Input: n = 3
Output: 3
```
説明:
1. `1 + 1 + 1 = 3`
2. `1 + 2 = 3`
3. `2 + 1 = 3`

### 制約

- `1 <= n <= 30`

## ソースコード

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    one, two := 1, 1
    for i := 0; i < n-1; i++ {
        one, two = one+two, one
    }
    return one
}
```

## アルゴリズムなど解説

### 基本戦略

各ステップで、**2つの選択肢**がある：
- **1段** 登る
- **2段** 登る

ステップ `i` に到達する方法の数は：
- ステップ `i - 1` からの方法数
- ステップ `i - 2` からの方法数
- の**合計**

これは**フィボナッチ数列**と同じパターン！

### 核心となる洞察

```
dp[i] = dp[i-1] + dp[i-2]
```

この漸化式は、各ステップへの到達方法が前の2つのステップの方法数の和であることを示す。

### 動作の仕組み

1. **基本ケース**
   ```go
   if n <= 2 {
       return n
   }
   ```
   - `n = 1`: 1通り（1段登る）
   - `n = 2`: 2通り（1+1 または 2）

2. **空間最適化されたDP**
   ```go
   one, two := 1, 1
   for i := 0; i < n-1; i++ {
       one, two = one+two, one
   }
   ```
   - `one`: 現在のステップへの方法数
   - `two`: 1つ前のステップへの方法数
   - 2つの変数だけで計算可能

### 具体例

```
n = 5

ステップ 1: 1通り
ステップ 2: 2通り
ステップ 3: 1 + 2 = 3通り
ステップ 4: 2 + 3 = 5通り
ステップ 5: 3 + 5 = 8通り

結果: 8
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | n回のループ |
| 空間計算量 | O(1) | 2つの変数のみ使用 |

### 別解：再帰 + メモ化

```go
func climbStairs(n int) int {
    cache := make([]int, n+1)
    for i := range cache {
        cache[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        if cache[i] != -1 {
            return cache[i]
        }
        cache[i] = dfs(i+1) + dfs(i+2)
        return cache[i]
    }
    return dfs(0)
}
```

**メモ化のアイデア**：
- 同じステップへの計算結果をキャッシュ
- 重複計算を回避

### 別解：DP配列

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

### DPパターン

この問題は **フィボナッチ型DP** パターン：
- 状態遷移: `dp[i] = dp[i-1] + dp[i-2]`
- 2つの前の状態のみに依存
- 空間最適化で O(1) 可能

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Climbing Stairs** | フィボナッチ型DP |
| House Robber | 「取る/取らない」型DP |
| Decode Ways | 文字列のフィボナッチ型DP |
| Fibonacci Number | 純粋なフィボナッチ |
