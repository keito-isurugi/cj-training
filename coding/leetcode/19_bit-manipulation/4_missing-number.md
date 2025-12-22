# Missing Number (Easy)

## 問題内容

`n` 個の整数を含む配列 `nums` が与えられる。配列には `[0, n]` の範囲の数が重複なく含まれている。

範囲内で配列に含まれていない1つの数を返す。

**フォローアップ**: `O(1)` の空間計算量と `O(n)` の時間計算量で実装できるか？

### 例

```
Input: nums = [1,2,3]

Output: 0
```

説明: 3つの数があるので範囲は[0,3]。0が配列にないので答えは0。

```
Input: nums = [0,2]

Output: 1
```

### 制約

- `1 <= nums.length <= 1000`

## ソースコード

```go
func missingNumber(nums []int) int {
    res := len(nums)
    for i := 0; i < len(nums); i++ {
        res += i - nums[i]
    }
    return res
}
```

## アルゴリズムなど解説

### 基本戦略（数学的アプローチ）

**発想**: 期待される合計 - 実際の合計 = 欠けている数

0から `n` までの合計: `n(n+1)/2`
配列の合計: `sum(nums)`
欠けている数: 期待 - 実際

### 動作の仕組み

```go
func missingNumber(nums []int) int {
    res := len(nums)  // nから開始
    for i := 0; i < len(nums); i++ {
        res += i - nums[i]  // 期待値iを足し、実際の値を引く
    }
    return res
}
```

これは `res = n + (0 + 1 + ... + (n-1)) - sum(nums)` と等価。

### 具体例

```
nums = [3, 0, 1]  (n = 3)

期待される合計: 0 + 1 + 2 + 3 = 6
実際の合計: 3 + 0 + 1 = 4
欠けている数: 6 - 4 = 2 ✓

アルゴリズムの動作:
res = 3
i=0: res += 0 - 3 = 0
i=1: res += 1 - 0 = 1
i=2: res += 2 - 1 = 2

結果: 2
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回走査 |
| 空間計算量 | O(1) | 定数空間 |

### 別解：XOR

```go
func missingNumber(nums []int) int {
    n := len(nums)
    xorr := n
    for i := 0; i < n; i++ {
        xorr ^= i ^ nums[i]
    }
    return xorr
}
```

**XORの性質**:
- `a ^ a = 0`（同じ数をXORすると0）
- `a ^ 0 = a`
- 交換法則・結合法則が成立

0からnまでと配列の全要素をXORすると、ペアになる数は消え、欠けている数だけ残る。

### XOR解の具体例

```
nums = [3, 0, 1]  (n = 3)

xorr = 3
xorr ^= 0 ^ 3 = 3 ^ 0 ^ 3 = 0
xorr ^= 1 ^ 0 = 0 ^ 1 ^ 0 = 1
xorr ^= 2 ^ 1 = 1 ^ 2 ^ 1 = 2

結果: 2 ✓
```

### 別解：HashSet

```go
func missingNumber(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }
    n := len(nums)
    for i := 0; i <= n; i++ {
        if _, exists := numSet[i]; !exists {
            return i
        }
    }
    return -1
}
```

シンプルだが空間計算量が O(n)。

### 3つの解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|------------|------------|------|
| 数学 | O(n) | O(1) | オーバーフローに注意 |
| XOR | O(n) | O(1) | オーバーフローなし |
| HashSet | O(n) | O(n) | 最もシンプル |

### なぜ数学的アプローチが安全か

```go
res += i - nums[i]
```

この形式だと、足し算と引き算が交互に行われるため、中間値が大きくなりにくい。

### Bit Manipulation パターン

XOR解は以下のパターンを利用：
- ペアの消去（`a ^ a = 0`）
- 欠けている要素の検出
- 配列と期待値の比較

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Single Number | XORで1つだけ残る |
| **Missing Number** | XORで欠けている数を検出 |
| Find the Duplicate | フロイドのサイクル検出 |
| First Missing Positive | インデックスマッピング |
