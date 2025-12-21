# Meeting Rooms II (Medium)

## 問題内容

開始時刻と終了時刻からなるミーティングの時間区間オブジェクトの配列 `[[start_1, end_1], [start_2, end_2], ...]`（`start_i < end_i`）が与えられる。

すべてのミーティングを競合なくスケジュールするために必要な最小日数（部屋数）を求める。

**注**: `(0, 8), (8, 10)` は時刻 8 では競合とみなさない。

### 例

```
Input: intervals = [(0,40),(5,10),(15,20)]
Output: 2
```
説明:
- day1: (0,40)
- day2: (5,10), (15,20)

```
Input: intervals = [(4,9)]
Output: 1
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

func minMeetingRooms(intervals []Interval) int {
    start := make([]int, len(intervals))
    end := make([]int, len(intervals))

    for i, interval := range intervals {
        start[i] = interval.start
        end[i] = interval.end
    }

    sort.Ints(start)
    sort.Ints(end)

    res, count := 0, 0
    s, e := 0, 0

    for s < len(intervals) {
        if start[s] < end[e] {
            s++
            count++
        } else {
            e++
            count--
        }
        if count > res {
            res = count
        }
    }

    return res
}
```

## アルゴリズムなど解説

### 基本戦略（Two Pointers）

**問題の本質**: 同時に行われるミーティングの最大数を求める

開始時刻と終了時刻を**別々に**ソートし、2つのポインタで処理：
- 開始時刻が先 → ミーティング開始（count++）
- 終了時刻が先 → ミーティング終了（count--）

同時刻の場合、終了を先に処理（部屋を解放してから新しいミーティングを開始）。

### 動作の仕組み

```go
for s < len(intervals) {
    if start[s] < end[e] {
        // 新しいミーティング開始
        s++
        count++
    } else {
        // ミーティング終了
        e++
        count--
    }
    if count > res {
        res = count
    }
}
```

### 具体例

```
intervals = [(0,40),(5,10),(15,20)]

start = [0, 5, 15]
end   = [10, 20, 40]

s=0, e=0: start[0]=0 < end[0]=10 → count=1, res=1
s=1, e=0: start[1]=5 < end[0]=10 → count=2, res=2
s=2, e=0: start[2]=15 >= end[0]=10 → count=1
s=2, e=1: start[2]=15 < end[1]=20 → count=2
s=3: 終了

結果: 2
```

### なぜ別々にソートしてよいか

重要なのは「ある時点で何個のミーティングが同時に行われているか」。

- 開始イベント：部屋が1つ必要になる
- 終了イベント：部屋が1つ解放される

イベントの発生順序だけが重要で、どのミーティングがどの部屋を使うかは関係ない。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n log n) | ソートが支配的 |
| 空間計算量 | O(n) | 開始・終了時刻の配列 |

### 別解1：Sweep Line Algorithm

```go
func minMeetingRooms(intervals []Interval) int {
    mp := make(map[int]int)
    for _, i := range intervals {
        mp[i.start]++
        mp[i.end]--
    }

    keys := make([]int, 0, len(mp))
    for k := range mp {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    prev := 0
    res := 0
    for _, k := range keys {
        prev += mp[k]
        if prev > res {
            res = prev
        }
    }
    return res
}
```

**Sweep Line の考え方**:
- 各時刻での変化量を記録
- 開始時刻で +1、終了時刻で -1
- 累積和の最大値が答え

### 別解2：Min Heap

```go
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

func minMeetingRooms(intervals []Interval) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    h := &MinHeap{}
    heap.Init(h)

    for _, interval := range intervals {
        if h.Len() > 0 && (*h)[0] <= interval.start {
            heap.Pop(h)
        }
        heap.Push(h, interval.end)
    }

    return h.Len()
}
```

**Min Heap の考え方**:
- 開始時刻でソート
- ヒープには使用中の部屋の終了時刻を格納
- 新しいミーティング開始時、最も早く終わる部屋が空いていれば再利用
- ヒープサイズが必要な部屋数

### 3つの解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|------------|------------|------|
| Two Pointers | O(n log n) | O(n) | シンプル |
| Sweep Line | O(n log n) | O(n) | 汎用的 |
| Min Heap | O(n log n) | O(n) | 部屋の割り当ても追跡可能 |

### Interval パターン

この問題は **最大同時実行数** パターン：
- イベントの開始と終了を分離
- 同時実行数の変化を追跡
- 最大値が答え

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Meeting Rooms | 重複判定 |
| **Meeting Rooms II** | 最大同時実行数 |
| Merge Intervals | 全区間マージ |
| Non-overlapping Intervals | 削除最小化 |
