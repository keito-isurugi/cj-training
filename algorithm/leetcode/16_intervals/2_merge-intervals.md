# Merge Intervals (Medium)

## 問題内容

区間の配列 `intervals` が与えられる。`intervals[i] = [start_i, end_i]` は `i` 番目の区間を表す。

すべての重複する区間をマージし、入力のすべての区間をカバーする重複しない区間の配列を返す。

結果は任意の順序で返してよい。

### 例

```
Input: intervals = [[1,3],[1,5],[6,7]]
Output: [[1,5],[6,7]]
```

```
Input: intervals = [[1,2],[2,3]]
Output: [[1,3]]
```

### 制約

- `1 <= intervals.length <= 1000`
- `intervals[i].length == 2`
- `0 <= start <= end <= 1000`

## ソースコード

```go
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    output := [][]int{intervals[0]}

    for _, interval := range intervals[1:] {
        start, end := interval[0], interval[1]
        lastEnd := output[len(output)-1][1]

        if start <= lastEnd {
            output[len(output)-1][1] = max(lastEnd, end)
        } else {
            output = append(output, []int{start, end})
        }
    }
    return output
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

1. **区間を開始時刻でソート**
2. **順番に走査してマージ**
   - 現在の区間が前の区間と重複 → マージ
   - 重複しない → 新しい区間として追加

### 重複の判定（ソート後）

```
前の区間: [..., lastEnd]
現在の区間: [start, end]

重複する条件: start <= lastEnd
```

### 動作の仕組み

1. **ソート**
   ```go
   sort.Slice(intervals, func(i, j int) bool {
       return intervals[i][0] < intervals[j][0]
   })
   ```

2. **最初の区間で初期化**
   ```go
   output := [][]int{intervals[0]}
   ```

3. **残りの区間を処理**
   ```go
   for _, interval := range intervals[1:] {
       start, end := interval[0], interval[1]
       lastEnd := output[len(output)-1][1]

       if start <= lastEnd {
           // マージ：終了時刻を更新
           output[len(output)-1][1] = max(lastEnd, end)
       } else {
           // 新しい区間を追加
           output = append(output, []int{start, end})
       }
   }
   ```

### 具体例

```
intervals = [[1,3],[1,5],[6,7]]

ソート後: [[1,3],[1,5],[6,7]]

output = [[1,3]]

interval = [1,5]:
  start=1 <= lastEnd=3 → マージ
  output = [[1,5]]

interval = [6,7]:
  start=6 > lastEnd=5 → 新規追加
  output = [[1,5],[6,7]]

結果: [[1,5],[6,7]]
```

### なぜソートが必要か

ソートすることで：
- **隣接する区間のみ比較**すればよくなる
- 区間Aの後に来る区間Bは必ず `A.start <= B.start`
- 区間Aと重複しない区間Bは、それ以降の区間とも区間Aは重複しない

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n log n) | ソートが支配的 |
| 空間計算量 | O(n) | 出力配列のサイズ |

### 別解：Sweep Line Algorithm

```go
func merge(intervals [][]int) [][]int {
    mp := make(map[int]int)
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        mp[start]++
        mp[end]--
    }

    res := [][]int{}
    interval := []int{}
    have := 0
    keys := make([]int, 0, len(mp))
    for key := range mp {
        keys = append(keys, key)
    }
    sort.Ints(keys)

    for _, i := range keys {
        if len(interval) == 0 {
            interval = append(interval, i)
        }
        have += mp[i]
        if have == 0 {
            interval = append(interval, i)
            res = append(res, append([]int{}, interval...))
            interval = []int{}
        }
    }
    return res
}
```

**Sweep Line の考え方**:
- 区間の開始で +1、終了で -1
- カウントが 0 → 0+ になる点が区間の開始
- カウントが 1+ → 0 になる点が区間の終了

### Interval パターン

この問題は **区間マージ** の基本パターン：
- ソートして隣接比較
- 重複判定は `start <= lastEnd`
- マージは終了時刻の max を取る

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Insert Interval | 1つの区間を挿入 |
| **Merge Intervals** | 全区間マージ |
| Non-overlapping Intervals | 削除最小化 |
| Meeting Rooms II | 同時開催数 |
