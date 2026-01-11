# Product of Array Except Self (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、`output[i]` が `nums[i]` を除く全要素の積となる配列 `output` を返す。

各積は32ビット整数に収まることが保証されている。

**Follow-up:** 除算を使わずに O(n) で解けるか？

### 例

```
Input: nums = [1,2,4,6]
Output: [48,24,12,8]

Input: nums = [-1,0,1,2,3]
Output: [0,-6,0,0,0]
```

### 制約

- `2 <= nums.length <= 1000`
- `-20 <= nums[i] <= 20`

## ソースコード

```go
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
```

## アルゴリズムなど解説

### 基本戦略

Prefix（左からの累積積）と Postfix（右からの累積積）を組み合わせて、各位置の「自分以外の積」を計算する。

### 動作の仕組み

1. **Prefix の計算**
   ```go
   for i := 1; i < len(nums); i++ {
       res[i] = res[i-1] * nums[i-1]
   }
   ```
   - `res[i]` に「i より左の全要素の積」を格納

2. **Postfix の計算と結合**
   ```go
   postfix := 1
   for i := len(nums) - 1; i >= 0; i-- {
       res[i] *= postfix
       postfix *= nums[i]
   }
   ```
   - 右から左へ走査し、Postfix（右の全要素の積）を掛ける
   - `res[i] = prefix[i] * postfix[i]` となる

### 具体例

#### 例: `nums = [1, 2, 4, 6]`

```
Step 1: Prefix の計算
  i=0: res[0] = 1        (左に要素なし)
  i=1: res[1] = 1 * 1 = 1
  i=2: res[2] = 1 * 2 = 2
  i=3: res[3] = 2 * 4 = 8

  res = [1, 1, 2, 8]  ← 各位置より左の積

Step 2: Postfix の計算と結合
  postfix = 1
  i=3: res[3] = 8 * 1 = 8,   postfix = 1 * 6 = 6
  i=2: res[2] = 2 * 6 = 12,  postfix = 6 * 4 = 24
  i=1: res[1] = 1 * 24 = 24, postfix = 24 * 2 = 48
  i=0: res[0] = 1 * 48 = 48, postfix = 48 * 1 = 48

  res = [48, 24, 12, 8]
```

#### 検証

```
nums = [1, 2, 4, 6]
res[0] = 2 * 4 * 6 = 48 ✓
res[1] = 1 * 4 * 6 = 24 ✓
res[2] = 1 * 2 * 6 = 12 ✓
res[3] = 1 * 2 * 4 = 8  ✓
```

### なぜこの方法が効率的か

```
アイデア:
- res[i] = (左の全要素の積) × (右の全要素の積)
- 左の積: prefix[i]
- 右の積: postfix[i]

効率性:
- 除算を使わない（0の問題を回避）
- 2回のパスで完了
- 追加の配列なしで実装可能（O(1) extra space）

比較:
- 除算を使う方法: 0が含まれると問題
- Brute Force: O(n²)
- Prefix + Postfix: O(n)
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を2回スキャン |
| 空間計算量 | O(1) | 出力配列を除く追加スペース |
