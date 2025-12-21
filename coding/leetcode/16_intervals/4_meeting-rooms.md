# Meeting Rooms (Easy)

## 問題内容

開始時刻と終了時刻からなるミーティングの時間区間オブジェクトの配列 `[[start_1, end_1], [start_2, end_2], ...]`（`start_i < end_i`）が与えられる。

1人の人がすべてのミーティングに出席できるかどうかを判定する。

**注**: `(0, 8), (8, 10)` は時刻 8 では競合とみなさない。

### 例

```
Input: intervals = [(0,30),(5,10),(15,20)]
Output: false
```
説明:
- `(0,30)` と `(5,10)` が競合
- `(0,30)` と `(15,20)` が競合

```
Input: intervals = [(5,8),(9,15)]
Output: true
```

### 制約

- `0 <= intervals.length <= 500`
- `0 <= intervals[i].start < intervals[i].end <= 1,000,000`

## ソースコード

```go
/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    for i := 1; i < len(intervals); i++ {
        i1 := intervals[i-1]
        i2 := intervals[i]
        if i1.end > i2.start {
            return false
        }
    }
    return true
}
```

## アルゴリズムなど解説

### 基本戦略

1. **開始時刻でソート**
2. **隣接する区間のペアを確認**
   - 前の区間の終了時刻 > 次の区間の開始時刻 → 重複

### 重複の判定

```
区間1: [start1, end1]
区間2: [start2, end2]  (start2 >= start1 でソート済み)

重複する条件: end1 > start2
```

### なぜソートして隣接比較で十分か

ソート後は `start_i <= start_{i+1}` が保証される。

- 区間 A と 区間 B が重複しない（A.end <= B.start）場合
- 区間 A の後に来るすべての区間 C（C.start >= B.start）とも重複しない

つまり、重複があれば必ず隣接するペアで検出される。

### 動作の仕組み

```go
for i := 1; i < len(intervals); i++ {
    i1 := intervals[i-1]
    i2 := intervals[i]
    if i1.end > i2.start {
        return false
    }
}
return true
```

### 具体例

```
intervals = [(0,30),(5,10),(15,20)]

ソート後: [(0,30),(5,10),(15,20)]

i=1: (0,30) と (5,10)
  30 > 5 → 重複あり → false

結果: false
```

```
intervals = [(5,8),(9,15)]

ソート後: [(5,8),(9,15)]

i=1: (5,8) と (9,15)
  8 <= 9 → 重複なし

結果: true
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n log n) | ソートが支配的 |
| 空間計算量 | O(1) | ソート用の空間を除く |

### 別解：Brute Force

```go
func canAttendMeetings(intervals []Interval) bool {
    n := len(intervals)
    for i := 0; i < n; i++ {
        A := intervals[i]
        for j := i + 1; j < n; j++ {
            B := intervals[j]
            if min(A.end, B.end) > max(A.start, B.start) {
                return false
            }
        }
    }
    return true
}
```

**計算量**: 時間 O(n²)、空間 O(1)

重複条件：`min(A.end, B.end) > max(A.start, B.start)`

### 重複判定の一般式

任意の2つの区間 A=[s1, e1] と B=[s2, e2] が重複する条件：
```
max(s1, s2) < min(e1, e2)
```

### Interval パターン

この問題は **区間重複判定** の基本：
- ソートして線形走査
- 隣接ペアのみチェック
- 1つでも重複があれば false

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Meeting Rooms** | 重複判定 |
| Meeting Rooms II | 必要な部屋数 |
| Non-overlapping Intervals | 削除最小化 |
| Merge Intervals | 全区間マージ |
