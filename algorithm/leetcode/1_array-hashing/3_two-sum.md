# Two Sum (Easy)

## 問題内容

整数配列 `nums` と整数 `target` が与えられたとき、和が `target` になる2つの要素のインデックス `i` と `j` を返す（`i != j`）。

各入力には必ず1つの解が存在することが保証されている。回答は小さいインデックスを先に返す。

### 例

```
Input: nums = [3,4,5,6], target = 7
Output: [0,1]
Explanation: nums[0] + nums[1] == 7

Input: nums = [4,5,6], target = 10
Output: [0,2]

Input: nums = [5,5], target = 10
Output: [0,1]
```

### 制約

- `2 <= nums.length <= 1000`
- `-10,000,000 <= nums[i] <= 10,000,000`
- `-10,000,000 <= target <= 10,000,000`
- 有効な解は1つのみ存在

## ソースコード

```go
func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, n := range nums {
			diff := target - n
			if j, found := prevMap[diff]; found {
					return []int{j, i}
			}
			prevMap[n] = i
	}
	return []int{}
}
```

## アルゴリズムなど解説

### 基本戦略

Hash Mapを使って、各要素の値とインデックスを記録しながら、補数（target - 現在の値）が既に存在するかを確認する。

### 動作の仕組み

1. **Hash Mapの初期化**
   ```go
   prevMap := make(map[int]int)
   ```
   - 値 → インデックス のマッピングを格納

2. **配列を走査**
   ```go
   for i, n := range nums {
       diff := target - n
       if j, found := prevMap[diff]; found {
           return []int{j, i}
       }
       prevMap[n] = i
   }
   ```
   - 各要素について、補数（`target - n`）がMapに存在するか確認
   - 存在すれば、その2つのインデックスを返す
   - 存在しなければ、現在の要素をMapに追加

### 具体例

#### 例: `nums = [3,4,5,6]`, `target = 7`

```
i=0, n=3:
  diff = 7 - 3 = 4
  prevMap = {} → 4は存在しない
  prevMap[3] = 0 → {3:0}

i=1, n=4:
  diff = 7 - 4 = 3
  prevMap = {3:0} → 3が存在！
  return [0, 1]
```

### なぜHash Mapが適切か

```
問題の性質:
- nums[i] + nums[j] = target
- つまり nums[j] = target - nums[i]

アプローチ:
- 各要素について、その補数を探す
- Hash Mapを使えばO(1)で検索可能

比較:
- Brute Force: O(n²) - 全ペアを試す
- ソート + Two Pointers: O(n log n)
- Hash Map: O(n) - 最も効率的
```

### One-Pass（1回のループで解決）

```
重要な点:
- 要素をMapに追加する前に補数を検索
- これにより同じ要素を2回使うことを防ぐ
- 例: target=6で[3,3]の場合も正しく動作
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回スキャン |
| 空間計算量 | O(n) | Hash Mapに最大n個の要素を格納 |
