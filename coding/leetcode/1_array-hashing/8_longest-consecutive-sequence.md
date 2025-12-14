# Longest Consecutive Sequence (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、形成できる最長の連続した要素のシーケンスの長さを返す。

**連続シーケンス**とは、各要素が前の要素より1大きいシーケンスのこと。要素は元の配列で連続している必要はない。

O(n) 時間で解くアルゴリズムを書く必要がある。

### 例

```
Input: nums = [2,20,4,10,3,4,5]
Output: 4
Explanation: 最長の連続シーケンスは [2, 3, 4, 5]

Input: nums = [0,3,2,5,4,6,1,1]
Output: 7
```

### 制約

- `0 <= nums.length <= 1000`
- `-10^9 <= nums[i] <= 10^9`

## ソースコード

```go
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, n := range nums {
		numSet[n] = true
	}
	longest := 0

	for n := range numSet {
		if !numSet[n-1] {
			length := 1
			for numSet[n+length] {
				length++
			}
			longest = max(length, longest)
		}
	}
	return longest
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

Hash Setを使い、各シーケンスの開始点を見つけて、そこから連続する要素を数える。

### 動作の仕組み

1. **Hash Setの作成**
   ```go
   numSet := make(map[int]bool)
   for _, n := range nums {
       numSet[n] = true
   }
   ```
   - O(1) で要素の存在確認ができるようにする

2. **シーケンスの開始点を見つける**
   ```go
   if !numSet[n-1] {
   ```
   - `n-1` が存在しない = `n` はシーケンスの開始点
   - これにより、同じシーケンスを複数回カウントすることを防ぐ

3. **シーケンスの長さを計算**
   ```go
   length := 1
   for numSet[n+length] {
       length++
   }
   ```
   - 開始点から連続する要素を数える

### 具体例

#### 例: `nums = [2, 20, 4, 10, 3, 4, 5]`

```
numSet = {2, 20, 4, 10, 3, 5}  ← 重複は自動的に除去

各要素をチェック:
  n=2:  numSet[1] = false → 開始点！
        length=1, numSet[3]=true → length=2
        length=2, numSet[4]=true → length=3
        length=3, numSet[5]=true → length=4
        length=4, numSet[6]=false → 終了
        longest = 4

  n=20: numSet[19] = false → 開始点！
        length=1, numSet[21]=false → 終了
        longest = max(4, 1) = 4

  n=4:  numSet[3] = true → 開始点ではない、スキップ

  n=10: numSet[9] = false → 開始点！
        length=1, numSet[11]=false → 終了
        longest = max(4, 1) = 4

  n=3:  numSet[2] = true → 開始点ではない、スキップ

  n=5:  numSet[4] = true → 開始点ではない、スキップ

return 4
```

### なぜ開始点の判定が重要か

```
問題点:
- 単純に各要素から数え始めると O(n²)
- 例: [1,2,3,4,5] で全要素から数えると重複作業が発生

解決策:
- n-1 が存在しない要素のみを開始点とする
- 各シーケンスは1回だけカウントされる
- トータルで O(n) を達成

例: [1, 2, 3, 4, 5]
  1: 開始点（0がない）→ シーケンス全体をカウント
  2: スキップ（1がある）
  3: スキップ（2がある）
  4: スキップ（3がある）
  5: スキップ（4がある）
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各要素は最大2回アクセスされる |
| 空間計算量 | O(n) | Hash Set に全要素を格納 |
