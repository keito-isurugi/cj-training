# Coin Change (Medium)

## 問題内容

異なる金額のコインを表す整数配列 `coins`（例: 1ドル、5ドルなど）と、目標金額を表す整数 `amount` が与えられる。

**正確な**目標金額を作るのに必要な**最小のコイン数**を返す。金額を作ることが不可能な場合は `-1` を返す。

各コインを**無制限に使用できる**と仮定してよい。

### 例

```
Input: coins = [1,5,10], amount = 12
Output: 3
```
説明: 12 = 10 + 1 + 1。すべての種類のコインを使う必要はない。

```
Input: coins = [2], amount = 3
Output: -1
```
説明: 金額 3 はコイン 2 では作れない。

```
Input: coins = [1], amount = 0
Output: 0
```
説明: 0枚のコインを選ぶのが金額0を作る有効な方法。

### 制約

- `1 <= coins.length <= 10`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10000`

## ソースコード

```go
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0

    for a := 1; a <= amount; a++ {
        for _, c := range coins {
            if c <= a {
                dp[a] = min(dp[a], dp[a-c]+1)
            }
        }
    }

    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## アルゴリズムなど解説

### 基本戦略

与えられた `amount` に対して、**すべてのコイン**を試す：
- 1枚選ぶ
- 残りの金額の問題を解く
- すべての選択肢の中から**最小コイン数**を選ぶ

これは**Unbounded Knapsack**（無制限ナップサック）パターン。

### 核心となる洞察

```
dp[a] = min(dp[a], dp[a - coin] + 1)（すべてのコインに対して）
```

金額 `a` を作る最小コイン数は：
- `a - coin` を作る最小コイン数 + 1

### 動作の仕組み

1. **DP配列の初期化**
   ```go
   dp := make([]int, amount+1)
   for i := range dp {
       dp[i] = amount + 1  // 不可能を示す大きな値
   }
   dp[0] = 0  // 金額0は0枚で作れる
   ```

2. **ボトムアップで計算**
   ```go
   for a := 1; a <= amount; a++ {
       for _, c := range coins {
           if c <= a {
               dp[a] = min(dp[a], dp[a-c]+1)
           }
       }
   }
   ```
   - 各金額 `a` に対して
   - すべてのコイン `c` を試す
   - `c` を使った場合のコイン数と比較

3. **結果の判定**
   ```go
   if dp[amount] > amount {
       return -1  // 不可能
   }
   return dp[amount]
   ```

### 具体例

```
coins = [1, 2, 5], amount = 11

dp[0] = 0

a=1: coin=1 → dp[1] = min(12, dp[0]+1) = 1
a=2: coin=1 → dp[2] = min(12, dp[1]+1) = 2
     coin=2 → dp[2] = min(2, dp[0]+1) = 1
a=3: coin=1 → dp[3] = min(12, dp[2]+1) = 2
     coin=2 → dp[3] = min(2, dp[1]+1) = 2
a=4: coin=1,2 → dp[4] = 2
a=5: coin=5 → dp[5] = min(3, dp[0]+1) = 1
a=6: coin=1,5 → dp[6] = 2
...
a=11: dp[11] = 3 (5+5+1)

結果: 3
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n × amount) | n = コイン数 |
| 空間計算量 | O(amount) | DP配列 |

### 別解：再帰 + メモ化

```go
func coinChange(coins []int, amount int) int {
    memo := make(map[int]int)

    var dfs func(amt int) int
    dfs = func(amt int) int {
        if amt == 0 {
            return 0
        }
        if val, ok := memo[amt]; ok {
            return val
        }

        res := amount + 1
        for _, coin := range coins {
            if amt-coin >= 0 {
                res = min(res, 1+dfs(amt-coin))
            }
        }
        memo[amt] = res
        return res
    }

    result := dfs(amount)
    if result > amount {
        return -1
    }
    return result
}
```

### 別解：BFS

```go
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    visited := make([]bool, amount+1)
    queue := []int{0}
    visited[0] = true
    steps := 0

    for len(queue) > 0 {
        steps++
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]
            for _, coin := range coins {
                next := curr + coin
                if next == amount {
                    return steps
                }
                if next < amount && !visited[next] {
                    visited[next] = true
                    queue = append(queue, next)
                }
            }
        }
    }

    return -1
}
```

**BFSのアイデア**：
- 各レベルが「使用したコイン数」を表す
- 最短経路問題として解く

### なぜ `amount + 1` で初期化するか

```
最大でも amount 枚のコイン（すべて1円コイン）で作れる
→ amount + 1 は「不可能」を表す十分に大きな値
→ MAX_INT を使うとオーバーフローの危険
```

### DPパターン

この問題は **Unbounded Knapsack** パターン：
- 各アイテム（コイン）を無制限に使用可能
- 最小化/最大化を求める
- 状態遷移: `dp[i] = min/max(dp[i], dp[i-w] + v)`

### 関連問題との比較

| 問題 | 目的 | 制限 |
|------|------|------|
| **Coin Change** | 最小コイン数 | 無制限 |
| Coin Change II | 組み合わせ数 | 無制限 |
| 0/1 Knapsack | 最大価値 | 各1個 |
| Climbing Stairs | 方法の数 | 1または2 |
