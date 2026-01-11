# Merge k Sorted Lists (Hard)

## 問題内容

k 個のソート済みリンクリストが与えられたとき、それらをすべてマージして1つのソート済みリストを返す。

### 例

```
Input: lists = [[1,2,4],[1,3,5],[3,6]]
Output: [1,1,2,3,3,4,5,6]

Input: lists = []
Output: []

Input: lists = [[]]
Output: []
```

### 制約

- `0 <= lists.length <= 1000`
- `0 <= lists[i].length <= 100`
- `-1000 <= lists[i][j] <= 1000`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    for len(lists) > 1 {
        var mergedLists []*ListNode
        for i := 0; i < len(lists); i += 2 {
            l1 := lists[i]
            var l2 *ListNode
            if i+1 < len(lists) {
                l2 = lists[i+1]
            }
            mergedLists = append(mergedLists, mergeList(l1, l2))
        }
        lists = mergedLists
    }
    return lists[0]
}

func mergeList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }

    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }

    return dummy.Next
}
```

## アルゴリズムなど解説

### 基本戦略

分割統治法（Divide and Conquer）を使用。ペアずつマージを繰り返し、最終的に1つのリストにする。

### マージソートと同じ考え方

```
k個のリストを一度に処理するのは複雑
→ 2つずつマージする操作を繰り返す
→ log(k) ラウンドで完了
```

### 視覚的な理解

```
Input: [L1, L2, L3, L4, L5, L6]

ラウンド1: ペアずつマージ
[L1, L2, L3, L4, L5, L6]
  ↓   ↓    ↓   ↓    ↓
[L1+L2,  L3+L4,  L5+L6]
= [M1,    M2,     M3]

ラウンド2:
[M1, M2, M3]
  ↓   ↓   ↓
[M1+M2,  M3]
= [N1,   N2]

ラウンド3:
[N1, N2]
  ↓   ↓
[N1+N2]
= [Result]

完了！
```

### 具体例: `lists = [[1,2,4],[1,3,5],[3,6]]`

```
初期:
L1: [1] → [2] → [4]
L2: [1] → [3] → [5]
L3: [3] → [6]

ラウンド1:
- L1 と L2 をマージ → M1: [1,1,2,3,4,5]
- L3 は奇数番目なので l2 = nil でマージ → M2: [3,6]

lists = [M1, M2]

ラウンド2:
- M1 と M2 をマージ → Result: [1,1,2,3,3,4,5,6]

lists = [Result]

len(lists) == 1 → 終了
return lists[0]
```

### 動作の仕組み

1. **メインループ: リストが1つになるまで繰り返す**
   ```go
   for len(lists) > 1 {
       var mergedLists []*ListNode
   ```

2. **ペアずつマージ**
   ```go
   for i := 0; i < len(lists); i += 2 {
       l1 := lists[i]
       var l2 *ListNode
       if i+1 < len(lists) {
           l2 = lists[i+1]  // 奇数個の場合の対策
       }
       mergedLists = append(mergedLists, mergeList(l1, l2))
   }
   ```

3. **次のラウンドへ**
   ```go
   lists = mergedLists
   ```

### mergeList関数（2つのリストをマージ）

```go
func mergeList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }

    // 残りを繋げる
    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }

    return dummy.Next
}
```

これは「Merge Two Sorted Lists」と同じ！

### 奇数個のリストの扱い

```go
var l2 *ListNode
if i+1 < len(lists) {
    l2 = lists[i+1]
}
// l2 が nil の場合、mergeList は l1 をそのまま返す
```

```
lists = [L1, L2, L3]

i=0: l1=L1, l2=L2 → マージ
i=2: l1=L3, l2=nil → L3がそのまま残る
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(N log k) | N=全ノード数, k=リスト数 |
| 空間計算量 | O(1)* | ポインタの付け替えのみ |

\* mergedListsの配列を含めるとO(k)

### なぜ O(N log k) か

```
- 各ラウンドで全ノード N 個を1回ずつ処理
- ラウンド数は log(k) 回（半分ずつ減るので）
- 合計: O(N × log k)
```

### 別解との比較

| アプローチ | 時間計算量 | 空間計算量 | 説明 |
|-----------|-----------|-----------|------|
| **分割統治（この解法）** | O(N log k) | O(1) | ペアずつマージ |
| 順次マージ | O(N × k) | O(1) | 1つずつマージしていく |
| Priority Queue | O(N log k) | O(k) | 全リストの先頭を管理 |

### 順次マージ（非効率な方法）

```go
// O(N × k) - 非効率
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    result := lists[0]
    for i := 1; i < len(lists); i++ {
        result = mergeList(result, lists[i])
    }
    return result
}
```

```
問題点:
- 1回目: N1 + N2 を処理
- 2回目: (N1 + N2) + N3 を処理
- ...
- 前のノードが何度も再処理される
```

### Priority Queue を使う方法

```go
import "container/heap"

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
    heap.Init(h)

    for _, l := range lists {
        if l != nil {
            heap.Push(h, l)
        }
    }

    dummy := &ListNode{}
    tail := dummy

    for h.Len() > 0 {
        min := heap.Pop(h).(*ListNode)
        tail.Next = min
        tail = tail.Next
        if min.Next != nil {
            heap.Push(h, min.Next)
        }
    }

    return dummy.Next
}
```

分割統治法はシンプルで効率的な解法！
