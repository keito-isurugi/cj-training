# Insert Interval (Medium)

## 問題内容

重複しない区間の配列 `intervals` が与えられる。`intervals[i] = [start_i, end_i]` は `i` 番目の区間の開始時刻と終了時刻を表す。`intervals` は最初から `start_i` の昇順でソートされている。

別の区間 `newInterval = [start, end]` が与えられる。

`newInterval` を `intervals` に挿入し、`intervals` が引き続き `start_i` の昇順でソートされ、重複する区間がないようにする。必要に応じて重複する区間をマージする。

挿入後の `intervals` を返す。

### 例

```
Input: intervals = [[1,3],[4,6]], newInterval = [2,5]
Output: [[1,6]]
```

```
Input: intervals = [[1,2],[3,5],[9,10]], newInterval = [6,7]
Output: [[1,2],[3,5],[6,7],[9,10]]
```

### 制約

- `0 <= intervals.length <= 1000`
- `newInterval.length == intervals[i].length == 2`
- `0 <= start <= end <= 1000`

## ソースコード

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    i := 0
    var res [][]int

    // 新しい区間より前に終わる区間を追加
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    }

    // 重複する区間をマージ
    for i < n && newInterval[1] >= intervals[i][0] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    res = append(res, newInterval)

    // 残りの区間を追加
    for i < n {
        res = append(res, intervals[i])
        i++
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
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

区間を3つのフェーズで処理する：

1. **新しい区間より前に終わる区間**：そのまま結果に追加
2. **新しい区間と重複する区間**：マージして1つの区間にする
3. **新しい区間より後に始まる区間**：そのまま結果に追加

### 重複の判定

```
区間Aと区間Bが重複する条件:
A.end >= B.start かつ A.start <= B.end
```

ソート済み配列なので、`newInterval[1] >= intervals[i][0]` のみで重複を判定可能。

### 動作の仕組み

1. **フェーズ1: 前の区間を追加**
   ```go
   for i < n && intervals[i][1] < newInterval[0] {
       res = append(res, intervals[i])
       i++
   }
   ```
   - 現在の区間の終了時刻が新区間の開始時刻より前なら追加

2. **フェーズ2: 重複区間をマージ**
   ```go
   for i < n && newInterval[1] >= intervals[i][0] {
       newInterval[0] = min(newInterval[0], intervals[i][0])
       newInterval[1] = max(newInterval[1], intervals[i][1])
       i++
   }
   res = append(res, newInterval)
   ```
   - 開始時刻は最小値、終了時刻は最大値を取る

3. **フェーズ3: 残りを追加**
   ```go
   for i < n {
       res = append(res, intervals[i])
       i++
   }
   ```

### 具体例

```
intervals = [[1,3],[4,6]], newInterval = [2,5]

フェーズ1: intervals[0] = [1,3]
  - intervals[0][1]=3 >= newInterval[0]=2 → 重複あり、フェーズ2へ

フェーズ2:
  - [2,5]と[1,3]をマージ → [1,5]
  - [1,5]と[4,6]をマージ → [1,6]
  - 結果に[1,6]を追加

フェーズ3: 残りなし

結果: [[1,6]]
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回走査 |
| 空間計算量 | O(1) | 出力配列を除き定数空間 |

### Interval パターン

この問題は **区間マージ** パターン：
- 区間がソート済みであることを活用
- 重複区間を逐次マージ
- 3フェーズで処理を分割

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Insert Interval** | 区間挿入・マージ |
| Merge Intervals | 全区間マージ |
| Non-overlapping Intervals | 削除数を最小化 |
| Meeting Rooms | 重複判定 |
