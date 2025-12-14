# Top K Frequent Elements (Medium)

## 問題内容

整数配列 `nums` と整数 `k` が与えられたとき、配列内で出現頻度が高い上位 `k` 個の要素を返す。

出力の順序は任意。

### 例

```
Input: nums = [1,2,2,3,3,3], k = 2
Output: [2,3]

Input: nums = [7,7], k = 1
Output: [7]
```

### 制約

- `1 <= nums.length <= 10^4`
- `-1000 <= nums[i] <= 1000`
- `1 <= k <= 配列内のユニークな要素数`

## ソースコード

```go
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n] = 1 + count[n]
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
```

## アルゴリズムなど解説

### 基本戦略

Bucket Sort を使用して、頻度ごとに要素をグループ化し、高頻度から順に k 個取得する。

### 動作の仕組み

1. **頻度カウント**
   ```go
   count := make(map[int]int)
   for _, n := range nums {
       count[n] = 1 + count[n]
   }
   ```
   - 各要素の出現回数をHash Mapでカウント

2. **Bucket への振り分け**
   ```go
   freq := make([][]int, len(nums)+1)
   for n, c := range count {
       freq[c] = append(freq[c], n)
   }
   ```
   - 頻度をインデックスとして、要素をBucketに格納
   - `freq[i]` には i 回出現する要素が入る

3. **高頻度から取得**
   ```go
   for i := len(freq) - 1; i > 0; i-- {
       for _, n := range freq[i] {
           res = append(res, n)
           if len(res) == k {
               return res
           }
       }
   }
   ```
   - 最高頻度から順にBucketを見ていく
   - k 個集まったら終了

### 具体例

#### 例: `nums = [1,2,2,3,3,3]`, `k = 2`

```
Step 1: 頻度カウント
  count = {1:1, 2:2, 3:3}

Step 2: Bucket 振り分け
  freq[0] = []
  freq[1] = [1]    ← 1回出現
  freq[2] = [2]    ← 2回出現
  freq[3] = [3]    ← 3回出現
  freq[4] = []
  freq[5] = []
  freq[6] = []

Step 3: 高頻度から取得
  i=6: freq[6] = [] → スキップ
  ...
  i=3: freq[3] = [3] → res = [3]
  i=2: freq[2] = [2] → res = [3, 2] → len(res) == k → 終了

return [3, 2]
```

### なぜBucket Sortが適切か

```
問題の特性:
- 頻度の最大値は配列の長さ n
- つまり頻度は 1 〜 n の範囲に収まる

Bucket Sort の利点:
- 頻度をインデックスとして直接アクセス可能
- ソート不要で O(n) を達成

比較:
- ソート: O(n log n)
- Heap: O(n log k)
- Bucket Sort: O(n) ← 最も効率的
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を2回スキャン |
| 空間計算量 | O(n) | count と freq 配列 |
