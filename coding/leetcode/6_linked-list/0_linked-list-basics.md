# Linked List 基礎

## Linked Listとは

配列と異なり、メモリ上で連続していないデータ構造。各要素（ノード）が次の要素への参照を持つ。

```
配列:     [0][1][2][3]  ← メモリ上で連続
Linked:   [0]→[1]→[2]→[3]→nil  ← ポインタで繋がる
```

## 基本構造

### 単方向リンクリスト（Singly Linked List）

```go
type ListNode struct {
    Val  int        // 値
    Next *ListNode  // 次のノードへのポインタ
}
```

```
┌─────┬──────┐    ┌─────┬──────┐    ┌─────┬──────┐
│ Val │ Next │ →  │ Val │ Next │ →  │ Val │ Next │ → nil
└─────┴──────┘    └─────┴──────┘    └─────┴──────┘
```

### 双方向リンクリスト（Doubly Linked List）

```go
type DoublyListNode struct {
    Val  int
    Prev *DoublyListNode  // 前のノード
    Next *DoublyListNode  // 次のノード
}
```

```
nil ← [Prev|Val|Next] ⇄ [Prev|Val|Next] ⇄ [Prev|Val|Next] → nil
```

## 配列との比較

| 操作 | 配列 | Linked List |
|------|------|-------------|
| インデックスアクセス | O(1) | O(n) |
| 先頭への挿入 | O(n) | O(1) |
| 末尾への挿入 | O(1)* | O(n)** |
| 中間への挿入 | O(n) | O(1)*** |
| 検索 | O(n) | O(n) |
| メモリ | 連続した領域が必要 | 分散可能 |

\* 動的配列の場合、リサイズ時はO(n)
\** tailポインタがあればO(1)
\*** 挿入位置が分かっている場合

## 基本操作

### ノードの作成

```go
node := &ListNode{Val: 1}
```

### リストの作成

```go
// 手動で作成
head := &ListNode{Val: 0}
head.Next = &ListNode{Val: 1}
head.Next.Next = &ListNode{Val: 2}

// 結果: 0 → 1 → 2 → nil
```

### リストの走査

```go
func traverse(head *ListNode) {
    curr := head
    for curr != nil {
        fmt.Println(curr.Val)
        curr = curr.Next
    }
}
```

### 先頭への挿入

```go
func insertAtHead(head *ListNode, val int) *ListNode {
    newNode := &ListNode{Val: val}
    newNode.Next = head
    return newNode  // 新しいheadを返す
}

// 0 → 1 → 2
// ↓ insertAtHead(head, -1)
// -1 → 0 → 1 → 2
```

### 末尾への挿入

```go
func insertAtTail(head *ListNode, val int) *ListNode {
    newNode := &ListNode{Val: val}

    if head == nil {
        return newNode
    }

    curr := head
    for curr.Next != nil {
        curr = curr.Next
    }
    curr.Next = newNode
    return head
}
```

### ノードの削除

```go
func deleteNode(head *ListNode, val int) *ListNode {
    // ダミーヘッドを使うテクニック
    dummy := &ListNode{Next: head}
    curr := dummy

    for curr.Next != nil {
        if curr.Next.Val == val {
            curr.Next = curr.Next.Next  // スキップして削除
            break
        }
        curr = curr.Next
    }
    return dummy.Next
}
```

## よく使うテクニック

### 1. ダミーヘッド（Sentinel Node）

先頭ノードの特別扱いを避けるテクニック。

```go
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}  // ダミーを先頭に置く
    curr := dummy

    for curr.Next != nil {
        if curr.Next.Val == val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return dummy.Next  // 本当のheadを返す
}
```

### 2. 2ポインタ（Fast & Slow）

サイクル検出や中間点を見つけるのに使用。

```go
// リストの中間点を見つける
func findMiddle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next       // 1歩進む
        fast = fast.Next.Next  // 2歩進む
    }
    return slow  // slowが中間点
}

// 0 → 1 → 2 → 3 → 4
//           ↑
//         中間点
```

### 3. サイクル検出（Floyd's Algorithm）

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true  // サイクルあり
        }
    }
    return false
}
```

### 4. リストの反転

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    return prev
}
```

## LeetCode頻出パターン

| パターン | 例題 |
|----------|------|
| 反転 | Reverse Linked List |
| マージ | Merge Two Sorted Lists |
| サイクル | Linked List Cycle |
| 2ポインタ | Remove Nth Node From End |
| 再帰 | Reverse Linked List (再帰版) |

## 注意点

### NilPointerエラー

```go
// NG: nilチェックなし
curr.Next.Val  // currやcurr.Nextがnilだとパニック

// OK: nilチェックあり
if curr != nil && curr.Next != nil {
    curr.Next.Val
}
```

### headの更新を忘れない

```go
// 先頭を削除する場合、headが変わる
func deleteHead(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    return head.Next  // 新しいheadを返す
}
```

## 計算量まとめ

| 操作 | 時間計算量 | 空間計算量 |
|------|-----------|-----------|
| 走査 | O(n) | O(1) |
| 検索 | O(n) | O(1) |
| 先頭挿入 | O(1) | O(1) |
| 末尾挿入 | O(n) | O(1) |
| 削除 | O(n) | O(1) |
| 反転 | O(n) | O(1) |
