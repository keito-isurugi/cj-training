# House Robber (Medium)

## 問題内容

`nums[i]` が `i` 番目の家にあるお金の量を表す整数配列 `nums` が与えられる。家は一直線に並んでおり、`i` 番目の家は `(i-1)` 番目と `(i+1)` 番目の家の隣にある。

家からお金を盗む計画を立てているが、セキュリティシステムが**隣接する2つの家**が両方とも侵入された場合に自動的に警察に通報するため、隣接する家を連続して盗むことはできない。

警察に通報されずに盗める**最大金額**を返す。

### 例

```
Input: nums = [1,1,3,3]
Output: 4
```
説明: `nums[0] + nums[2] = 1 + 3 = 4`

```
Input: nums = [2,9,8,3,6]
Output: 16
```
説明: `nums[0] + nums[2] + nums[4] = 2 + 8 + 6 = 16`

### 制約

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

## ソースコード

```go
func rob(nums []int) int {
    rob1, rob2 := 0, 0
    for _, num := range nums {
        temp := max(num+rob1, rob2)
        rob1 = rob2
        rob2 = temp
    }
    return rob2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## アルゴリズムなど解説

### 基本戦略

各家で**2つの選択肢**がある：
- **スキップ**する → 次の家に進む
- **盗む** → そのお金を取り、次の次の家に進む

目標は**最大金額**を得る選択をすること。

### 核心となる洞察

家 `i` での最大利益は：
```
dp[i] = max(dp[i-1], nums[i] + dp[i-2])
```
- `dp[i-1]`: この家をスキップ（前の家までの最大）
- `nums[i] + dp[i-2]`: この家を盗む（2つ前までの最大 + この家の金額）

### 動作の仕組み

1. **空間最適化されたDP**
   ```go
   rob1, rob2 := 0, 0
   for _, num := range nums {
       temp := max(num+rob1, rob2)
       rob1 = rob2
       rob2 = temp
   }
   ```
   - `rob1`: 2つ前の家までの最大利益
   - `rob2`: 1つ前の家までの最大利益
   - 各家で「盗む」か「スキップ」の最大を選択

2. **変数のシフト**
   ```go
   rob1 = rob2
   rob2 = temp
   ```
   - 次の家に進む準備
   - 古い状態を捨て、新しい状態を保持

### 具体例

```
nums = [2, 9, 8, 3, 6]

初期: rob1=0, rob2=0

家0 (2): temp = max(2+0, 0) = 2  → rob1=0, rob2=2
家1 (9): temp = max(9+0, 2) = 9  → rob1=2, rob2=9
家2 (8): temp = max(8+2, 9) = 10 → rob1=9, rob2=10
家3 (3): temp = max(3+9, 10) = 12 → rob1=10, rob2=12
家4 (6): temp = max(6+10, 12) = 16 → rob1=12, rob2=16

結果: 16 (家0 + 家2 + 家4 = 2 + 8 + 6)
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各家を1回処理 |
| 空間計算量 | O(1) | 2つの変数のみ使用 |

### 別解：DP配列

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], nums[i]+dp[i-2])
    }

    return dp[n-1]
}
```

### 別解：再帰 + メモ化

```go
func rob(nums []int) int {
    n := len(nums)
    memo := make([]int, n+1)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            return 0
        }
        if memo[i] != -1 {
            return memo[i]
        }
        memo[i] = max(dfs(i+1), nums[i]+dfs(i+2))
        return memo[i]
    }

    return dfs(0)
}
```

### DPパターン

この問題は **「取る/取らない」型DP** パターン：
- 各要素で二択（取る or スキップ）
- 隣接要素の制約あり
- 状態遷移: `dp[i] = max(dp[i-1], nums[i] + dp[i-2])`

### 関連問題との比較

| 問題 | 違い |
|------|------|
| **House Robber** | 一直線の家 |
| House Robber II | 円形に並んだ家 |
| Climbing Stairs | 選択の最大化ではなく数え上げ |
| Delete and Earn | 値をキーとして同様のロジック |
