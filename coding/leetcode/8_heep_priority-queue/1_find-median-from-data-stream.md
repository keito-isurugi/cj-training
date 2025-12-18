# Find Median from Data Stream (Hard)

## 問題内容

中央値（メディアン）は、ソート済みリストの中央の値。偶数長のリストでは中央値がないため、中央の2つの値の平均が中央値となる。

例:
- `arr = [1,2,3]` の中央値は `2`
- `arr = [1,2]` の中央値は `(1 + 2) / 2 = 1.5`

`MedianFinder` クラスを実装する:
- `MedianFinder()`: オブジェクトを初期化
- `void addNum(int num)`: データストリームから整数を追加
- `double findMedian()`: これまでの全要素の中央値を返す

### 例

```
Input:
["MedianFinder", "addNum", "1", "findMedian", "addNum", "3", "findMedian", "addNum", "2", "findMedian"]

Output:
[null, null, 1.0, null, 2.0, null, 2.0]

Explanation:
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.findMedian(); // return 1.0
medianFinder.addNum(3);    // arr = [1, 3]
medianFinder.findMedian(); // return 2.0
medianFinder.addNum(2);    // arr = [1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

### 制約

- `-100,000 <= num <= 100,000`
- `findMedian` は少なくとも1つの整数が追加された後にのみ呼ばれる

## ソースコード

```go
type MedianFinder struct {
    small *priorityqueue.Queue // maxHeap
    large *priorityqueue.Queue // minHeap
}

func Constructor() MedianFinder {
    small := priorityqueue.NewWith(func(a, b interface{}) int {
        return b.(int) - a.(int)  // maxHeap
    })
    large := priorityqueue.NewWith(func(a, b interface{}) int {
        return a.(int) - b.(int)  // minHeap
    })
    return MedianFinder{small: small, large: large}
}

func (this *MedianFinder) AddNum(num int) {
    if this.large.Size() > 0 {
        largeTop, _ := this.large.Peek()
        if num > largeTop.(int) {
            this.large.Enqueue(num)
        } else {
            this.small.Enqueue(num)
        }
    } else {
        this.small.Enqueue(num)
    }

    // Rebalance
    if this.small.Size() > this.large.Size()+1 {
        val, _ := this.small.Dequeue()
        this.large.Enqueue(val)
    }
    if this.large.Size() > this.small.Size()+1 {
        val, _ := this.large.Dequeue()
        this.small.Enqueue(val)
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.small.Size() > this.large.Size() {
        val, _ := this.small.Peek()
        return float64(val.(int))
    }
    if this.large.Size() > this.small.Size() {
        val, _ := this.large.Peek()
        return float64(val.(int))
    }
    smallVal, _ := this.small.Peek()
    largeVal, _ := this.large.Peek()
    return float64(smallVal.(int)+largeVal.(int)) / 2.0
}
```

## アルゴリズムなど解説

### 基本戦略

2つのヒープを使って、データを「小さい半分」と「大きい半分」に分ける。中央値は常にこの境界にある。

### 2つのヒープの役割

```
small (MaxHeap): 小さい方の半分を管理
                 最大値がトップに来る

large (MinHeap): 大きい方の半分を管理
                 最小値がトップに来る

ソート済み配列をイメージ:
[1, 2, 3, | 4, 5, 6]
 ←small→   ←large→
   max=3     min=4

中央値 = (3 + 4) / 2 = 3.5
```

### なぜ2つのヒープか

```
ナイーブな方法:
- 毎回ソート → O(n log n)
- ソート済み配列に挿入 → O(n)

2つのヒープ:
- 追加 → O(log n)
- 中央値取得 → O(1)

データストリームで効率的！
```

### ヒープの構造

```
MaxHeap (small):        MinHeap (large):
     3                       4
    / \                     / \
   1   2                   5   6

small.Peek() = 3        large.Peek() = 4

small は「小さい半分の最大」を即座に取得
large は「大きい半分の最小」を即座に取得
```

### AddNum の動作

```go
func (this *MedianFinder) AddNum(num int) {
    // 1. どちらのヒープに追加するか決定
    if this.large.Size() > 0 {
        largeTop, _ := this.large.Peek()
        if num > largeTop.(int) {
            this.large.Enqueue(num)  // 大きい方へ
        } else {
            this.small.Enqueue(num)  // 小さい方へ
        }
    } else {
        this.small.Enqueue(num)  // 最初はsmallへ
    }

    // 2. リバランス（サイズ差を1以内に保つ）
    if this.small.Size() > this.large.Size()+1 {
        val, _ := this.small.Dequeue()
        this.large.Enqueue(val)
    }
    if this.large.Size() > this.small.Size()+1 {
        val, _ := this.large.Dequeue()
        this.small.Enqueue(val)
    }
}
```

### 視覚的な理解

```
操作: addNum(1), addNum(3), addNum(2)

━━━ Step 1: addNum(1) ━━━
large is empty → small に追加

small: [1]  (max=1)
large: []

━━━ Step 2: addNum(3) ━━━
3 > large.top? → large is empty, なので small へ

small: [1, 3]? → いや、まずsmallに入れる
small: [3, 1]  (max=3)
large: []

リバランス: small.Size(2) > large.Size(0) + 1
→ smallから取り出してlargeへ

small: [1]  (max=1)
large: [3]  (min=3)

━━━ Step 3: addNum(2) ━━━
2 > large.top(3)? → No
→ small に追加

small: [2, 1]  (max=2)
large: [3]     (min=3)

リバランス: small.Size(2) > large.Size(1) + 1? → No
           large.Size(1) > small.Size(2) + 1? → No
→ OK

━━━ findMedian() ━━━
small.Size(2) > large.Size(1)
→ small.Peek() = 2

中央値 = 2.0
```

### FindMedian の動作

```go
func (this *MedianFinder) FindMedian() float64 {
    // ケース1: smallが多い → smallのトップが中央値
    if this.small.Size() > this.large.Size() {
        val, _ := this.small.Peek()
        return float64(val.(int))
    }
    // ケース2: largeが多い → largeのトップが中央値
    if this.large.Size() > this.small.Size() {
        val, _ := this.large.Peek()
        return float64(val.(int))
    }
    // ケース3: 同数 → 両方のトップの平均
    smallVal, _ := this.small.Peek()
    largeVal, _ := this.large.Peek()
    return float64(smallVal.(int)+largeVal.(int)) / 2.0
}
```

### 3つのケース

```
ケース1: 奇数個でsmallが多い
small: [2, 1]  large: [3]
       ↑
     中央値

ケース2: 奇数個でlargeが多い
small: [1]  large: [2, 3]
                   ↑
                 中央値

ケース3: 偶数個で同数
small: [2, 1]  large: [3, 4]
       ↑              ↑
中央値 = (2 + 3) / 2 = 2.5
```

### 計算量

| 操作 | 時間計算量 | 説明 |
|------|-----------|------|
| AddNum | O(log n) | ヒープ操作 |
| FindMedian | O(1) | トップを見るだけ |

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 空間計算量 | O(n) | 全要素を保持 |

### 別解：標準ライブラリ版（container/heap使用）

```go
import "container/heap"

type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }  // 大きい方が優先
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }  // 小さい方が優先
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

type MedianFinder struct {
    small *MaxHeap
    large *MinHeap
}

func Constructor() MedianFinder {
    small := &MaxHeap{}
    large := &MinHeap{}
    heap.Init(small)
    heap.Init(large)
    return MedianFinder{small: small, large: large}
}

func (this *MedianFinder) AddNum(num int) {
    if this.large.Len() > 0 && num > (*this.large)[0] {
        heap.Push(this.large, num)
    } else {
        heap.Push(this.small, num)
    }

    // Rebalance
    if this.small.Len() > this.large.Len()+1 {
        heap.Push(this.large, heap.Pop(this.small))
    }
    if this.large.Len() > this.small.Len()+1 {
        heap.Push(this.small, heap.Pop(this.large))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() > this.large.Len() {
        return float64((*this.small)[0])
    }
    if this.large.Len() > this.small.Len() {
        return float64((*this.large)[0])
    }
    return float64((*this.small)[0]+(*this.large)[0]) / 2.0
}
```

### 重要なポイント

```
1. 2つのヒープで半分ずつ管理
   - small (MaxHeap): 小さい半分 → 最大値がトップ
   - large (MinHeap): 大きい半分 → 最小値がトップ

2. リバランスでサイズ差を1以内に保つ
   → 中央値は常にトップにある

3. 中央値の取得
   - 奇数個: 多い方のトップ
   - 偶数個: 両方のトップの平均

4. なぜヒープ？
   - 最大/最小の取得が O(1)
   - 追加が O(log n)
   - データストリームに最適
```

### 関連問題

| 問題 | 共通点 |
|------|--------|
| Kth Largest Element | ヒープの活用 |
| Sliding Window Median | この問題の拡張版 |
| Top K Frequent Elements | ヒープでのランキング |

この問題は**Hard**で、2つのヒープを使う典型的なパターン。面接でよく出題される！
