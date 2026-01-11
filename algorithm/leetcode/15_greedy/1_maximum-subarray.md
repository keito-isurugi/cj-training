# Maximum Subarray (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、最大の合計を持つ**部分配列**を見つけ、その合計を返す。

**部分配列**とは、配列内の連続した空でない要素の列のこと。

### 例

```
Input: nums = [2,-3,4,-2,2,1,-1,4]
Output: 8
```
説明: 部分配列 [4,-2,2,1,-1,4] が最大の合計 8 を持つ。

```
Input: nums = [-1]
Output: -1
```

### 制約

- `1 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`

## ソースコード

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    curSum := 0

    for _, num := range nums {
        if curSum < 0 {
            curSum = 0
        }
        curSum += num
        if curSum > maxSum {
            maxSum = curSum
        }
    }
    return maxSum
}
```

## アルゴリズムなど解説

### 基本戦略（Kadane's Algorithm）

各位置で2つの選択肢がある：
- **現在の要素を既存の部分配列に追加する**
- **現在の要素から新しい部分配列を開始する**

これは**貪欲法**のアプローチで、**Kadane's Algorithm** として知られている。

### 核心となる洞察

```
現在の合計が負の場合、新しい部分配列を開始した方が良い
```

負の累積和を引きずると、全体の合計が小さくなるため。

### 動作の仕組み

1. **累積和のリセット**
   ```go
   if curSum < 0 {
       curSum = 0
   }
   ```
   - 現在の累積和が負なら、ここから新しく始める

2. **要素の追加と最大値の更新**
   ```go
   curSum += num
   if curSum > maxSum {
       maxSum = curSum
   }
   ```
   - 現在の要素を追加
   - 最大値を更新

### 具体例

```
nums = [2, -3, 4, -2, 2, 1, -1, 4]

i=0: num=2,  curSum=2,  maxSum=2
i=1: num=-3, curSum=-1, maxSum=2
i=2: num=4,  curSum=0→4, maxSum=4  (リセット)
i=3: num=-2, curSum=2,  maxSum=4
i=4: num=2,  curSum=4,  maxSum=4
i=5: num=1,  curSum=5,  maxSum=5
i=6: num=-1, curSum=4,  maxSum=5
i=7: num=4,  curSum=8,  maxSum=8

結果: 8 (部分配列: [4,-2,2,1,-1,4])
```

### なぜこれが最適解なのか

任意の時点で：
- **負の累積和**を持ち続けると、次の要素を加えても損をする
- 「負の荷物」を捨てて、新しくスタートする方が常に良い

これがKadane's Algorithmの本質。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回走査 |
| 空間計算量 | O(1) | 変数2つのみ使用 |

### 別解：DP配列を使用

```go
func maxSubArray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    maxSum := dp[0]

    for i := 1; i < n; i++ {
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        if dp[i] > maxSum {
            maxSum = dp[i]
        }
    }
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

`dp[i]` = インデックス `i` で終わる部分配列の最大合計

### 別解：分割統治法

```go
func maxSubArray(nums []int) int {
    return divideAndConquer(nums, 0, len(nums)-1)
}

func divideAndConquer(nums []int, left, right int) int {
    if left > right {
        return -1 << 31
    }
    if left == right {
        return nums[left]
    }

    mid := (left + right) / 2
    leftMax := divideAndConquer(nums, left, mid)
    rightMax := divideAndConquer(nums, mid+1, right)
    crossMax := maxCrossing(nums, left, mid, right)

    return max(leftMax, max(rightMax, crossMax))
}

func maxCrossing(nums []int, left, mid, right int) int {
    leftSum := -1 << 31
    sum := 0
    for i := mid; i >= left; i-- {
        sum += nums[i]
        if sum > leftSum {
            leftSum = sum
        }
    }

    rightSum := -1 << 31
    sum = 0
    for i := mid + 1; i <= right; i++ {
        sum += nums[i]
        if sum > rightSum {
            rightSum = sum
        }
    }

    return leftSum + rightSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**計算量**: 時間 O(n log n)、空間 O(log n)

### Greedy パターン

この問題は **Kadane's Algorithm** パターン：
- 局所最適解を積み重ねて大域最適解を得る
- 負の累積和は捨てる
- 1回の走査で解が求まる

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Maximum Subarray** | Kadane's Algorithm |
| Maximum Product Subarray | 変形版 Kadane's |
| Best Time to Buy and Sell Stock | 類似の貪欲法 |
| Maximum Sum Circular Subarray | Kadane's + 円環考慮 |
