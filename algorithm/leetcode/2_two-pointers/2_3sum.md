# 3Sum (Medium)

## 問題内容

整数配列 `nums` が与えられたとき、`nums[i] + nums[j] + nums[k] == 0` を満たす全てのトリプレット `[nums[i], nums[j], nums[k]]` を返す。ただし、インデックス `i`、`j`、`k` は全て異なる必要がある。

出力に重複するトリプレットが含まれてはならない。出力とトリプレットの順序は任意で良い。

### 例

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

説明:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0
異なるトリプレットは [-1,0,1] と [-1,-1,2] の2つ。
```

```
Input: nums = [0,1,1]
Output: []

説明: 合計が0になるトリプレットは存在しない。
```

```
Input: nums = [0,0,0]
Output: [[0,0,0]]

説明: 唯一のトリプレットの合計は0。
```

### 制約

- `3 <= nums.length <= 1000`
- `-10^5 <= nums[i] <= 10^5`

## ソースコード

```go
func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        a := nums[i]
        if a > 0 {
            break
        }
        if i > 0 && a == nums[i-1] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            threeSum := a + nums[l] + nums[r]
            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                res = append(res, []int{a, nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            }
        }
    }

    return res
}
```

## アルゴリズムなど解説

### 基本戦略

1. 配列をソートする
2. 1つの要素を固定し、残りの2つの要素をTwo Pointersで探索する
3. 重複を避けるための処理を行う

### なぜソートするのか

- Two Pointersテクニックを使用するため（ソートされていないと正しく動作しない）
- 重複のスキップが容易になる
- 早期終了の条件を設定できる（最小値が正なら解なし）

### 動作の仕組み

1. **配列をソート**
   ```go
   sort.Ints(nums)
   ```
   - 例: `[-1,0,1,2,-1,-4]` → `[-4,-1,-1,0,1,2]`

2. **最初の要素を固定してループ**
   ```go
   for i := 0; i < len(nums); i++ {
       a := nums[i]
   ```

3. **早期終了条件**
   ```go
   if a > 0 {
       break
   }
   ```
   - ソート済みなので、最初の要素が正なら残りも全て正
   - 3つの正の数の合計は0にならない

4. **重複をスキップ**
   ```go
   if i > 0 && a == nums[i-1] {
       continue
   }
   ```
   - 同じ値で既に処理済みならスキップ

5. **Two Pointersで残り2つを探索**
   ```go
   l, r := i+1, len(nums)-1
   for l < r {
       threeSum := a + nums[l] + nums[r]
   ```
   - `l`: 固定要素の右隣からスタート
   - `r`: 配列の右端からスタート

6. **合計に応じてポインタを移動**
   ```go
   if threeSum > 0 {
       r--  // 合計が大きすぎる → 右を小さく
   } else if threeSum < 0 {
       l++  // 合計が小さすぎる → 左を大きく
   } else {
       // 見つかった！
   }
   ```

7. **解が見つかった後の重複スキップ**
   ```go
   for l < r && nums[l] == nums[l-1] {
       l++
   }
   ```

### 具体例

```
nums = [-4,-1,-1,0,1,2] (ソート済み)

i=0, a=-4
  l=1, r=5: -4 + (-1) + 2 = -3 < 0 → l++
  l=2, r=5: -4 + (-1) + 2 = -3 < 0 → l++
  l=3, r=5: -4 + 0 + 2 = -2 < 0 → l++
  l=4, r=5: -4 + 1 + 2 = -1 < 0 → l++
  l=5, r=5: 終了

i=1, a=-1
  l=2, r=5: -1 + (-1) + 2 = 0 → 解発見! [-1,-1,2]
  l=3, r=4: -1 + 0 + 1 = 0 → 解発見! [-1,0,1]
  l=4, r=4: 終了

i=2, a=-1 → 前と同じ値なのでスキップ

i=3, a=0 → 以降は解なし

結果: [[-1,-1,2], [-1,0,1]]
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n²) | 外側ループ O(n) × Two Pointers O(n) |
| 空間計算量 | O(1) または O(n) | ソートアルゴリズムに依存（結果を除く） |

### なぜTwo Pointersが有効か

- ソート済み配列では、合計が大きすぎれば右を小さく、小さすぎれば左を大きくすることで効率的に探索できる
- 全ペアを調べる O(n²) の代わりに、各固定要素に対して O(n) で探索可能

### 別解：HashMapを使用する方法

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }

    var res [][]int
    for i := 0; i < len(nums); i++ {
        count[nums[i]]--
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        for j := i + 1; j < len(nums); j++ {
            count[nums[j]]--
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            target := -(nums[i] + nums[j])
            if count[target] > 0 {
                res = append(res, []int{nums[i], nums[j], target})
            }
        }

        for j := i + 1; j < len(nums); j++ {
            count[nums[j]]++
        }
    }

    return res
}
```

- 時間計算量: O(n²)
- 空間計算量: O(n)

Two Pointersの方が空間効率が良い。
