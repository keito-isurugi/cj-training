# Contains Duplicate (Easy)

## 問題内容

整数配列 `nums` が与えられたとき、配列内に重複する値が存在すれば `true` を、存在しなければ `false` を返す。

### 例

```
Input: nums = [1, 2, 3, 3]
Output: true

Input: nums = [1, 2, 3, 4]
Output: false
```

## ソースコード

```go
func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    for _, num := range nums {
        seen[num] = struct{}{}
    }
    return len(seen) < len(nums)
}
```

## アルゴリズムなど解説

### 基本戦略

HashMapを使って、配列内のユニークな要素数と元の配列の長さを比較する。

### 動作の仕組み

1. **空のマップを作成**
   ```go
   seen := make(map[int]struct{})
   ```
   - `map[int]struct{}` を使用（値に空構造体を使うことでメモリ効率が良い）

2. **全要素をマップに追加**
   ```go
   for _, num := range nums {
       seen[num] = struct{}{}
   }
   ```
   - マップのキーは重複しないため、同じ値は1つにまとまる

3. **長さを比較**
   ```go
   return len(seen) < len(nums)
   ```
   - ユニーク数 < 元の長さ → 重複あり
   - ユニーク数 == 元の長さ → 重複なし

### 具体例

```
nums = [1, 2, 3, 3]

ステップ1: seen = {1: {}}
ステップ2: seen = {1: {}, 2: {}}
ステップ3: seen = {1: {}, 2: {}, 3: {}}
ステップ4: seen = {1: {}, 2: {}, 3: {}}  ← 3は既に存在するので変化なし

len(seen) = 3
len(nums) = 4

3 < 4 → true（重複あり）
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回スキャン |
| 空間計算量 | O(n) | 最大でn個の要素をマップに格納 |

### 別解：早期リターン版

```go
func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    for _, num := range nums {
        if _, exists := seen[num]; exists {
            return true  // 重複を見つけたら即座にtrue
        }
        seen[num] = struct{}{}
    }
    return false
}
```

- 重複が見つかった時点で早期リターンできる
- 平均的なケースではより効率的
