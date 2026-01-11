# Reverse Linked List (Easy)

## 問題内容

単方向リンクリストの先頭 `head` が与えられたとき、リストを反転して新しい先頭を返す。

### 例

```
Input: head = [0,1,2,3]
Output: [3,2,1,0]

Input: head = []
Output: []
```

### 制約

- `0 <= リストの長さ <= 1000`
- `-1000 <= Node.val <= 1000`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

## アルゴリズムなど解説

### 基本戦略

3つのポインタ（prev, curr, temp）を使って、各ノードの向きを1つずつ反転させる。

### 動作の仕組み

1. **初期化**
   ```go
   var prev *ListNode  // nil（反転後の末尾になる）
   curr := head        // 現在処理中のノード
   ```

2. **各ノードを反転**
   ```go
   for curr != nil {
       temp := curr.Next  // 次のノードを保存
       curr.Next = prev   // 向きを反転
       prev = curr        // prevを進める
       curr = temp        // currを進める
   }
   ```

3. **新しい先頭を返す**
   ```go
   return prev  // prevが新しい先頭
   ```

### 視覚的な理解

```
初期状態:
prev = nil
curr = 0

nil    0 → 1 → 2 → 3 → nil
 ↑     ↑
prev  curr
```

#### ステップ1

```
temp = curr.Next (1を保存)
curr.Next = prev (0→nilに変更)
prev = curr (prevを0に移動)
curr = temp (currを1に移動)

nil ← 0    1 → 2 → 3 → nil
      ↑    ↑
     prev curr
```

#### ステップ2

```
temp = curr.Next (2を保存)
curr.Next = prev (1→0に変更)
prev = curr (prevを1に移動)
curr = temp (currを2に移動)

nil ← 0 ← 1    2 → 3 → nil
          ↑    ↑
         prev curr
```

#### ステップ3

```
temp = curr.Next (3を保存)
curr.Next = prev (2→1に変更)
prev = curr (prevを2に移動)
curr = temp (currを3に移動)

nil ← 0 ← 1 ← 2    3 → nil
              ↑    ↑
             prev curr
```

#### ステップ4

```
temp = curr.Next (nilを保存)
curr.Next = prev (3→2に変更)
prev = curr (prevを3に移動)
curr = temp (currをnilに移動)

nil ← 0 ← 1 ← 2 ← 3    nil
                  ↑     ↑
                 prev  curr

curr == nil → ループ終了
return prev (3)
```

### なぜtempが必要か

```
curr.Next = prev を実行すると、
元々のcurr.Next（次のノード）への参照が失われる

temp = curr.Next で先に保存しておかないと
リストの残りにアクセスできなくなる
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各ノードを1回ずつ処理 |
| 空間計算量 | O(1) | ポインタ3つのみ使用 |

### 別解：再帰版

```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil

    return newHead
}
```

- 空間計算量: O(n)（コールスタック）
- 反復版の方がメモリ効率が良い
