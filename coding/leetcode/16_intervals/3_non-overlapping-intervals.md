# Non-overlapping Intervals (Medium)

## 問題内容

区間の配列 `intervals` が与えられる。`intervals[i] = [start_i, end_i]` は `i` 番目の区間を表す。

残りの区間が重複しないようにするために削除する必要がある区間の最小数を返す。

**注**: 共通点を持つ区間は重複とみなす。例えば `[1, 3]` と `[2, 4]` は重複するが、`[1, 2]` と `[2, 3]` は重複しない。

### 例

```
Input: intervals = [[1,2],[2,4],[1,4]]
Output: 1
```
説明: `[1,4]` を削除すると、残りの区間は重複しない。

```
Input: intervals = [[1,2],[2,4]]
Output: 0
```

### 制約

- `1 <= intervals.length <= 1000`
- `intervals[i].length == 2`
- `-50000 <= start_i < end_i <= 50000`

## ソースコード

```go
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    res := 0
    prevEnd := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start >= prevEnd {
            // 重複なし
            prevEnd = end
        } else {
            // 重複あり：終了が遅い方を削除
            res++
            prevEnd = min(prevEnd, end)
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## アルゴリズムなど解説

### 基本戦略（Greedy）

**発想の転換**: 削除数を最小化 = 残す区間を最大化

重複する区間のペアがある場合、**終了時刻が早い方を残す**のが最適：
- 終了が早い方を残すと、後続の区間と重複しにくい
- より多くの区間を残せる可能性が高まる

### 貪欲法の正当性

```
区間A: [1, 4]
区間B: [1, 2]
区間C: [3, 5]

Aを残す場合: Cと重複 → 残せるのは1つ
Bを残す場合: Cと重複しない → 残せるのは2つ
```

終了時刻が早い区間を優先的に残すことで、後の区間を残す余地が増える。

### 動作の仕組み

1. **開始時刻でソート**

2. **順番に走査**
   ```go
   for i := 1; i < len(intervals); i++ {
       start, end := intervals[i][0], intervals[i][1]
       if start >= prevEnd {
           // 重複なし：現在の区間を採用
           prevEnd = end
       } else {
           // 重複あり：終了が遅い方を削除
           res++
           prevEnd = min(prevEnd, end)
       }
   }
   ```

3. **重複時の処理**
   - `prevEnd = min(prevEnd, end)` で終了が早い方を残す
   - 削除カウントをインクリメント

### 具体例

```
intervals = [[1,2],[2,4],[1,4]]

ソート後: [[1,2],[1,4],[2,4]]

prevEnd = 2, res = 0

[1,4]: start=1 < prevEnd=2 → 重複
  prevEnd = min(2, 4) = 2, res = 1

[2,4]: start=2 >= prevEnd=2 → 重複なし
  prevEnd = 4

結果: 1
```

### 別のソート方法：終了時刻でソート

```go
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    res := 0
    prevEnd := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] >= prevEnd {
            prevEnd = intervals[i][1]
        } else {
            res++
        }
    }
    return res
}
```

終了時刻でソートすると、常に先の区間を残せばよいのでコードがシンプルになる。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n log n) | ソートが支配的 |
| 空間計算量 | O(1) | ソート用の空間を除く |

### 別解：DP（参考）

```go
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    n := len(intervals)
    dp := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] {
                dp[i] = max(dp[i], 1+dp[j])
            }
        }
    }

    maxNonOverlapping := 0
    for _, v := range dp {
        if v > maxNonOverlapping {
            maxNonOverlapping = v
        }
    }
    return n - maxNonOverlapping
}
```

**計算量**: 時間 O(n²)、空間 O(n)

### Greedy パターン

この問題は **Activity Selection Problem** の変形：
- 終了時刻が早い区間を優先
- 重複したら終了が遅い方を削除
- 削除数 = 全体 - 残せる最大数

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Insert Interval | 区間挿入 |
| Merge Intervals | 全区間マージ |
| **Non-overlapping Intervals** | 削除最小化（Greedy） |
| Meeting Rooms | 重複判定 |
