# Reorder List (Medium)

## 問題内容

リンクリストを以下の順序に並び替える（ノードの値ではなく、ノード自体を並び替える）。

```
元:     [0, 1, 2, 3, 4, 5, 6]
並替後: [0, 6, 1, 5, 2, 4, 3]

一般化: [0, n-1, 1, n-2, 2, n-3, ...]
```

### 例

```
Input: head = [2,4,6,8]
Output: [2,8,4,6]

Input: head = [2,4,6,8,10]
Output: [2,10,4,8,6]
```

### 制約

- `1 <= リストの長さ <= 1000`
- `1 <= Node.val <= 1000`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    // Step 1: 中間点を見つける
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Step 2: 後半を反転
    second := slow.Next
    slow.Next = nil
    var prev *ListNode
    for second != nil {
        tmp := second.Next
        second.Next = prev
        prev = second
        second = tmp
    }

    // Step 3: 前半と後半をマージ
    first := head
    second = prev
    for second != nil {
        tmp1, tmp2 := first.Next, second.Next
        first.Next = second
        second.Next = tmp1
        first, second = tmp1, tmp2
    }
}
```

## アルゴリズムなど解説

### 基本戦略

3つのステップに分解する：
1. リストの中間点を見つける
2. 後半を反転する
3. 前半と後半を交互にマージする

### 全体の流れ

```
元のリスト:
[2] → [4] → [6] → [8] → nil

Step 1: 中間点で分割
前半: [2] → [4] → nil
後半: [6] → [8] → nil

Step 2: 後半を反転
前半: [2] → [4] → nil
後半: [8] → [6] → nil

Step 3: 交互にマージ
結果: [2] → [8] → [4] → [6] → nil
```

### Step 1: 中間点を見つける（Fast & Slow）

```go
slow, fast := head, head.Next
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
```

```
[2] → [4] → [6] → [8] → nil
 s     f

[2] → [4] → [6] → [8] → nil
       s           f

fast.Next == nil → 終了
slow = 中間点（前半の最後）
```

### Step 2: 後半を反転

```go
second := slow.Next    // 後半の開始点
slow.Next = nil        // 前半と後半を切り離す
var prev *ListNode
for second != nil {
    tmp := second.Next
    second.Next = prev
    prev = second
    second = tmp
}
```

```
切り離し後:
前半: [2] → [4] → nil
後半: [6] → [8] → nil

反転後:
前半: [2] → [4] → nil
後半: [8] → [6] → nil  (prevが先頭)
```

### Step 3: 交互にマージ

```go
first := head
second = prev
for second != nil {
    tmp1, tmp2 := first.Next, second.Next
    first.Next = second
    second.Next = tmp1
    first, second = tmp1, tmp2
}
```

#### 詳細なトレース

```
初期:
first = [2], second = [8]

ステップ1:
tmp1 = [4], tmp2 = [6]
first.Next = second  → [2] → [8]
second.Next = tmp1   → [2] → [8] → [4]
first = tmp1 = [4]
second = tmp2 = [6]

結果: [2] → [8] → [4]

ステップ2:
tmp1 = nil, tmp2 = nil
first.Next = second  → [4] → [6]
second.Next = tmp1   → [4] → [6] → nil
first = tmp1 = nil
second = tmp2 = nil

結果: [2] → [8] → [4] → [6] → nil

second == nil → 終了
```

### 奇数長の場合

```
[2] → [4] → [6] → [8] → [10] → nil

Step 1: 中間点
[2] → [4] → [6]       slow = [6]
[8] → [10] → nil

Step 2: 反転
前半: [2] → [4] → [6] → nil
後半: [10] → [8] → nil

Step 3: マージ
[2] → [10] → [4] → [8] → [6] → nil
                         ↑
                       前半の余り（そのまま残る）
```

### なぜ3ステップに分けるか

```
直接的なアプローチ:
- 末尾から要素を取得するのはO(n)
- 毎回末尾を探すとO(n²)になる

このアプローチ:
- Step 1: O(n)
- Step 2: O(n)
- Step 3: O(n)
- 合計: O(n)
```

### 既出テクニックの組み合わせ

| ステップ | テクニック | 元の問題 |
|----------|-----------|---------|
| Step 1 | Fast & Slow | Linked List Cycle / Find Middle |
| Step 2 | リスト反転 | Reverse Linked List |
| Step 3 | リストマージ | Merge Two Sorted Lists |

この問題は**3つの基本テクニックを組み合わせた応用問題**。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各ステップがO(n) |
| 空間計算量 | O(1) | ポインタのみ使用 |
