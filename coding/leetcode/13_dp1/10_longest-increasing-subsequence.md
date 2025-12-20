# Longest Increasing Subsequence (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、**最長の厳密に増加する部分列**の**長さ**を返す。

**部分列**とは、与えられた配列から要素を削除して（または削除せずに）、残りの要素の相対的な順序を変えずに得られる配列。

- 例えば、`"cat"` は `"crabt"` の部分列。

### 例

```
Input: nums = [9,1,4,2,3,3,7]
Output: 4
```
説明: 最長の増加部分列は [1,2,3,7] で、長さは 4。

```
Input: nums = [0,3,1,3,2,3]
Output: 4
```

### 制約

- `1 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`

## ソースコード

```go
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }

    maxLen := 1
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if nums[i] < nums[j] {
                dp[i] = max(dp[i], 1+dp[j])
            }
        }
        maxLen = max(maxLen, dp[i])
    }

    return maxLen
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

各インデックス `i` で：
> "`nums[i]` から始まる最長増加部分列の長さは？"

`nums[i] < nums[j]`（j > i）であるすべての `j` について：
- `nums[j]` を部分列に含めることができる
- `1 + dp[j]` が可能な長さ

すべての選択肢の**最大**を取る。

### 核心となる洞察

```
dp[i] = max(1 + dp[j])（すべての j > i かつ nums[i] < nums[j] について）
```

または等価な前方向の定式化：
```
dp[i] = max(1 + dp[j])（すべての j < i かつ nums[j] < nums[i] について）
```

### 動作の仕組み

1. **DP配列の初期化**
   ```go
   dp := make([]int, n)
   for i := range dp {
       dp[i] = 1  // 各要素自体が長さ1のLIS
   }
   ```

2. **後ろから前へ計算**
   ```go
   for i := n - 1; i >= 0; i-- {
       for j := i + 1; j < n; j++ {
           if nums[i] < nums[j] {
               dp[i] = max(dp[i], 1+dp[j])
           }
       }
       maxLen = max(maxLen, dp[i])
   }
   ```

3. **グローバル最大を返す**
   ```go
   return maxLen
   ```

### 具体例

```
nums = [10, 9, 2, 5, 3, 7, 101, 18]

後ろから計算:
dp[7] = 1 (18)
dp[6] = 1 (101)
dp[5] = 2 (7 → 18 or 101)
dp[4] = 3 (3 → 7 → 18)
dp[3] = 3 (5 → 7 → 18)
dp[2] = 4 (2 → 3 → 7 → 18)
dp[1] = 1 (9: 後続に9より大きく使えるものがない)
dp[0] = 1 (10: 同様)

最大: dp[2] = 4
LIS: [2, 3, 7, 18] or [2, 5, 7, 101] など
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n²) | 2重ループ |
| 空間計算量 | O(n) | DP配列 |

### 別解：再帰 + メモ化

```go
func lengthOfLIS(nums []int) int {
    n := len(nums)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if memo[i] != -1 {
            return memo[i]
        }

        LIS := 1
        for j := i + 1; j < n; j++ {
            if nums[i] < nums[j] {
                LIS = max(LIS, 1+dfs(j))
            }
        }

        memo[i] = LIS
        return LIS
    }

    maxLen := 1
    for i := 0; i < n; i++ {
        maxLen = max(maxLen, dfs(i))
    }
    return maxLen
}
```

### 別解：二分探索 O(n log n)

```go
func lengthOfLIS(nums []int) int {
    tails := []int{}

    for _, num := range nums {
        pos := sort.SearchInts(tails, num)
        if pos == len(tails) {
            tails = append(tails, num)
        } else {
            tails[pos] = num
        }
    }

    return len(tails)
}
```

**二分探索のアイデア**：
- `tails[i]` = 長さ `i+1` のLISの最小の終端要素
- 新しい要素が来たら、適切な位置を二分探索で見つける
- `tails` の長さが答え

### 二分探索の動作例

```
nums = [10, 9, 2, 5, 3, 7, 101, 18]

tails = []
num=10: tails = [10]
num=9:  tails = [9]  (10を9で置換)
num=2:  tails = [2]  (9を2で置換)
num=5:  tails = [2, 5]
num=3:  tails = [2, 3]  (5を3で置換)
num=7:  tails = [2, 3, 7]
num=101: tails = [2, 3, 7, 101]
num=18: tails = [2, 3, 7, 18]  (101を18で置換)

長さ: 4
```

### DPパターン

この問題は **LIS (Longest Increasing Subsequence)** パターン：
- 典型的な1次元DP
- O(n²) のナイーブ解
- O(n log n) の二分探索解

### 関連問題との比較

| 問題 | 制約 | 目的 |
|------|------|------|
| **Longest Increasing Subsequence** | 厳密に増加 | 長さ |
| Number of Longest Increasing Subsequence | 厳密に増加 | 個数 |
| Longest Continuous Increasing Subsequence | 連続 | 長さ |
| Russian Doll Envelopes | 2次元LIS | 長さ |
| Increasing Triplet Subsequence | 長さ3存在 | bool |
