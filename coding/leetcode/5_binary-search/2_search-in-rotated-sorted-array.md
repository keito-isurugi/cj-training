# Search in Rotated Sorted Array (Medium)

## 問題内容

回転されたソート済み配列 `nums` と整数 `target` が与えられたとき、`target` のインデックスを返す。存在しない場合は `-1` を返す。

O(log n) で解くことが求められる。

### 例

```
Input: nums = [3,4,5,6,1,2], target = 1
Output: 4

Input: nums = [3,5,6,0,1,2], target = 4
Output: -1
```

### 制約

- `1 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-1000 <= target <= 1000`
- すべての要素はユニーク

## ソースコード

```go
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1

    for l <= r {
        mid := (l + r) / 2
        if target == nums[mid] {
            return mid
        }

        if nums[l] <= nums[mid] {
            if target > nums[mid] || target < nums[l] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else {
            if target < nums[mid] || target > nums[r] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}
```

## アルゴリズムなど解説

### 基本戦略

Binary Searchを使うが、回転配列なので「どちら半分がソート済みか」を判定してから探索方向を決める。

### 回転配列の特性

```
[3, 4, 5, 6, 1, 2]
 \_______/  \___/
 ソート済み   ソート済み

どちらか一方は必ずソート済み
→ ソート済みの方にtargetがあるか判定できる
```

### 動作の仕組み

1. **midを見つけたら返す**
   ```go
   if target == nums[mid] {
       return mid
   }
   ```

2. **左半分がソート済みか判定**
   ```go
   if nums[l] <= nums[mid] {
       // 左半分 [l, mid] はソート済み
   } else {
       // 右半分 [mid, r] はソート済み
   }
   ```

3. **targetがどちらにあるか判定**
   ```go
   // 左半分がソート済みの場合
   if target > nums[mid] || target < nums[l] {
       l = mid + 1  // targetは右半分にある
   } else {
       r = mid - 1  // targetは左半分にある
   }
   ```

### なぜこの判定が正しいか

```
左半分 [l, mid] がソート済みの場合:

[3, 4, 5, 6, 1, 2]   target = 5
 l        m     r

左半分: [3, 4, 5, 6] はソート済み
→ nums[l]=3 <= target=5 <= nums[mid]=6
→ targetは左半分にある → r = mid - 1

[3, 4, 5, 6, 1, 2]   target = 1
 l        m     r

左半分: [3, 4, 5, 6] はソート済み
→ target=1 < nums[l]=3
→ targetは左半分にない → l = mid + 1
```

### 具体例: `nums = [3,4,5,6,1,2], target = 1`

```
初期: l=0, r=5

ステップ1:
[3, 4, 5, 6, 1, 2]
 l        m     r

mid = 2, nums[mid] = 5
target(1) != 5

nums[l]=3 <= nums[mid]=5 → 左半分はソート済み
target=1 < nums[l]=3 → targetは左半分にない
→ l = mid + 1 = 3

ステップ2:
[3, 4, 5, 6, 1, 2]
          l  m  r

mid = 4, nums[mid] = 1
target(1) == 1 → return 4 ✓
```

### 具体例: `nums = [4,5,6,7,0,1,2], target = 0`

```
初期: l=0, r=6

ステップ1:
[4, 5, 6, 7, 0, 1, 2]
 l        m        r

mid = 3, nums[mid] = 7
target(0) != 7

nums[l]=4 <= nums[mid]=7 → 左半分はソート済み
target=0 < nums[l]=4 → targetは左半分にない
→ l = mid + 1 = 4

ステップ2:
[4, 5, 6, 7, 0, 1, 2]
             l  m  r

mid = 5, nums[mid] = 1
target(0) != 1

nums[l]=0 <= nums[mid]=1 → 左半分はソート済み
nums[l]=0 <= target=0 <= nums[mid]=1 → targetは左半分
→ r = mid - 1 = 4

ステップ3:
[4, 5, 6, 7, 0, 1, 2]
             l
             r
             m

mid = 4, nums[mid] = 0
target(0) == 0 → return 4 ✓
```

### Find Minimum との違い

| 問題 | 目的 | ループ条件 | 返り値 |
|------|------|-----------|--------|
| Find Minimum | 最小値を探す | `l < r` | `nums[l]` |
| Search | targetを探す | `l <= r` | `mid` or `-1` |

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(log n) | Binary Searchで毎回半分に絞る |
| 空間計算量 | O(1) | ポインタのみ使用 |
