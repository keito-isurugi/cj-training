# Maximum Product Subarray (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、**最大の積**を持つ**部分配列**を見つけ、その積を返す。

**部分配列**とは、配列内の連続した空でない要素のシーケンス。

出力は**32ビット**整数に収まると仮定してよい。

### 例

```
Input: nums = [1,2,-3,4]
Output: 4
```

```
Input: nums = [-2,-1]
Output: 2
```

### 制約

- `1 <= nums.length <= 1000`
- `-10 <= nums[i] <= 10`

## ソースコード

```go
func maxProduct(nums []int) int {
    res := nums[0]
    curMin, curMax := 1, 1

    for _, num := range nums {
        tmp := curMax * num
        curMax = max(max(num*curMax, num*curMin), num)
        curMin = min(min(tmp, num*curMin), num)
        res = max(res, curMax)
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
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

この問題が厄介な理由：
- **負の数** → 符号が反転（非常に小さい負の値が非常に大きい正の値になる可能性）
- **ゼロ** → 積をリセット

そのため、各インデックスで**2つの値**を追跡：
- `curMax`: このインデックスで終わる最大積
- `curMin`: このインデックスで終わる最小積

`curMin` が重要な理由：
- 次に負の数が来たら、`curMin × 負` が新しい最大になる可能性

### 核心となる洞察

```
curMax = max(num * curMax, num * curMin, num)
curMin = min(num * curMax, num * curMin, num)
```

各要素で3つの選択肢：
1. 前の最大と掛ける
2. 前の最小と掛ける（負×負=正の可能性）
3. この要素から新しく始める

### 動作の仕組み

1. **初期化**
   ```go
   res := nums[0]
   curMin, curMax := 1, 1
   ```
   - `res`: グローバルな最大値
   - `curMin, curMax`: 現在位置までの最小/最大積

2. **各要素での更新**
   ```go
   tmp := curMax * num
   curMax = max(max(num*curMax, num*curMin), num)
   curMin = min(min(tmp, num*curMin), num)
   res = max(res, curMax)
   ```
   - `tmp`: `curMax` が更新される前に保存
   - 3つの選択肢の最大/最小を計算
   - グローバル最大を更新

### 具体例

```
nums = [2, 3, -2, 4]

初期: res=2, curMin=1, curMax=1

num=2:
  curMax = max(2*1, 2*1, 2) = 2
  curMin = min(2*1, 2*1, 2) = 2
  res = max(2, 2) = 2

num=3:
  curMax = max(3*2, 3*2, 3) = 6
  curMin = min(6, 6, 3) = 3
  res = max(2, 6) = 6

num=-2:
  tmp = 6 * (-2) = -12
  curMax = max(-12, 3*(-2), -2) = max(-12, -6, -2) = -2
  curMin = min(-12, -6, -2) = -12
  res = max(6, -2) = 6

num=4:
  tmp = -2 * 4 = -8
  curMax = max(-8, -12*4, 4) = max(-8, -48, 4) = 4
  curMin = min(-8, -48, 4) = -48
  res = max(6, 4) = 6

結果: 6
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 1回のパス |
| 空間計算量 | O(1) | 定数個の変数 |

### 別解：ブルートフォース

```go
func maxProduct(nums []int) int {
    res := nums[0]

    for i := 0; i < len(nums); i++ {
        cur := nums[i]
        res = max(res, cur)
        for j := i + 1; j < len(nums); j++ {
            cur *= nums[j]
            res = max(res, cur)
        }
    }

    return res
}
```
- 時間計算量: O(n²)
- すべての部分配列を試す

### なぜ最小値も追跡するのか

```
例: [-2, 3, -4]

最小を追跡しない場合:
  i=0: max = -2
  i=1: max = max(-6, 3) = 3
  i=2: max = max(-12, -4) = -4
  結果: 3（誤り）

最小も追跡する場合:
  i=0: max=-2, min=-2
  i=1: max=3, min=-6
  i=2: max=max(-12, 24, -4)=24, min=...
  結果: 24（正解）
```

負×負=正になるため、最小値が最大値に変わる可能性がある！

### Kadane's Algorithm との違い

| 観点 | Maximum Subarray (Sum) | Maximum Product Subarray |
|------|------------------------|-------------------------|
| 演算 | 加算 | 乗算 |
| 追跡 | 最大のみ | 最大と最小 |
| 負の影響 | 常にマイナス | 符号反転の可能性 |
| ゼロの影響 | 単に足される | リセットになる |

### DPパターン

この問題は **修正版 Kadane's Algorithm** パターン：
- 元の Kadane は最大のみ追跡
- 積の問題では最大と最小の両方が必要
- 各要素で「新しく始める」選択肢を含める

### 関連問題との比較

| 問題 | 演算 | 追跡するもの |
|------|------|-------------|
| Maximum Subarray | 和 | 最大のみ |
| **Maximum Product Subarray** | 積 | 最大と最小 |
| Best Time to Buy and Sell Stock | 差 | 最小価格 |
