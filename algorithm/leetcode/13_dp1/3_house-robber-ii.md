# House Robber II (Medium)

## 問題内容

`nums[i]` が `i` 番目の家にあるお金の量を表す整数配列 `nums` が与えられる。家は**円形**に並んでおり、最初の家と最後の家が隣接している。

家からお金を盗む計画を立てているが、セキュリティシステムが**隣接する2つの家**が両方とも侵入された場合に自動的に警察に通報するため、隣接する家を連続して盗むことはできない。

警察に通報されずに盗める**最大金額**を返す。

### 例

```
Input: nums = [3,4,3]
Output: 4
```
説明: `nums[0] + nums[2] = 6` は不可能（最初と最後が隣接）。最大は `nums[1] = 4`。

```
Input: nums = [2,9,8,3,6]
Output: 15
```
説明: `nums[0] + nums[2] + nums[4] = 16` は不可能（nums[0]とnums[4]が隣接）。最大は `nums[1] + nums[4] = 15`。

### 制約

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

## ソースコード

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
    rob1, rob2 := 0, 0
    for i := start; i <= end; i++ {
        temp := max(nums[i]+rob1, rob2)
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

**House Robber** の拡張版で、家が**円形**に並んでいる。
最初の家と最後の家が隣接しているため、**両方を同時に盗むことはできない**。

解決策：問題を**2つの線形ケース**に分割：
1. 最初の家から最後から2番目の家まで（最後の家を除外）
2. 2番目の家から最後の家まで（最初の家を除外）

この2つの結果の**最大値**が答え。

### 核心となる洞察

```
結果 = max(rob(nums[0:n-1]), rob(nums[1:n]))
```

どちらかのケースでは必ず「最初と最後が同時に含まれない」ことが保証される。

### 動作の仕組み

1. **特殊ケース処理**
   ```go
   if n == 1 {
       return nums[0]
   }
   ```
   - 家が1つだけの場合、その家を盗む

2. **2つの範囲で計算**
   ```go
   return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
   ```
   - `robRange(0, n-2)`: 最初の家を含み、最後を除外
   - `robRange(1, n-1)`: 最初の家を除外し、最後を含む

3. **各範囲で House Robber I と同じロジック**
   ```go
   func robRange(nums []int, start, end int) int {
       rob1, rob2 := 0, 0
       for i := start; i <= end; i++ {
           temp := max(nums[i]+rob1, rob2)
           rob1 = rob2
           rob2 = temp
       }
       return rob2
   }
   ```

### 具体例

```
nums = [2, 9, 8, 3, 6]

ケース1: nums[0:4] = [2, 9, 8, 3]
  初期: rob1=0, rob2=0
  家0 (2): temp = max(2+0, 0) = 2  → rob1=0, rob2=2
  家1 (9): temp = max(9+0, 2) = 9  → rob1=2, rob2=9
  家2 (8): temp = max(8+2, 9) = 10 → rob1=9, rob2=10
  家3 (3): temp = max(3+9, 10) = 12 → rob1=10, rob2=12
  結果: 12

ケース2: nums[1:5] = [9, 8, 3, 6]
  初期: rob1=0, rob2=0
  家1 (9): temp = max(9+0, 0) = 9  → rob1=0, rob2=9
  家2 (8): temp = max(8+0, 9) = 9  → rob1=9, rob2=9
  家3 (3): temp = max(3+9, 9) = 12 → rob1=9, rob2=12
  家4 (6): temp = max(6+9, 12) = 15 → rob1=12, rob2=15
  結果: 15

最終結果: max(12, 15) = 15
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を2回走査（各O(n)） |
| 空間計算量 | O(1) | 定数個の変数のみ使用 |

### 別解：DP配列

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    return max(helper(nums[1:]), helper(nums[:n-1]))
}

func helper(nums []int) int {
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

### なぜ2つのケースに分割するのか

```
円形配置: [A, B, C, D, E, A]（AとEが隣接）

最適解は必ず以下のいずれか：
1. Aを盗まない → [B, C, D, E] の最適解
2. Eを盗まない → [A, B, C, D] の最適解

両方の場合を試して最大を取れば、
円形制約を満たしつつ最大利益を得られる
```

### DPパターン

この問題は **円形DP** パターン：
- 線形DPの拡張
- 最初と最後の要素の依存関係を解消
- 問題を2つの線形サブ問題に分割

### 関連問題との比較

| 問題 | 配置 | 追加制約 |
|------|------|---------|
| House Robber | 線形 | なし |
| **House Robber II** | 円形 | 最初と最後が隣接 |
| House Robber III | 木構造 | 親子が隣接 |
