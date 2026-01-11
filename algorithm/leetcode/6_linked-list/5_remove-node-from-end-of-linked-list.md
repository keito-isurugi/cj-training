# Remove Nth Node From End of List (Medium)

## 問題内容

リンクリストの先頭 `head` と整数 `n` が与えられたとき、末尾から n 番目のノードを削除し、リストの先頭を返す。

### 例

```
Input: head = [1,2,3,4], n = 2
Output: [1,2,4]
Explanation: 末尾から2番目の「3」を削除

Input: head = [5], n = 1
Output: []
Explanation: 唯一のノードを削除

Input: head = [1,2], n = 2
Output: [2]
Explanation: 先頭ノードを削除
```

### 制約

- `1 <= sz <= 30`（リストの長さ）
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    left := dummy
    right := head

    for n > 0 {
        right = right.Next
        n--
    }

    for right != nil {
        left = left.Next
        right = right.Next
    }

    left.Next = left.Next.Next
    return dummy.Next
}
```

## アルゴリズムなど解説

### 基本戦略

2つのポインタを n 個分離して同時に進め、right が末尾に達したとき left が削除対象の1つ前にいるようにする。

### 核心的なアイデア

```
末尾から n 番目 = 先頭から (長さ - n + 1) 番目

しかし、リストの長さを知るには1回走査が必要...

Two Pointer で1回の走査で解決！
- right を n 個先に置く
- left と right を同時に進める
- right が末尾に達したとき、left は削除対象の1つ前
```

### 動作の仕組み

1. **ダミーヘッドを作成**
   ```go
   dummy := &ListNode{Next: head}
   left := dummy
   right := head
   ```
   - ダミーを使うことで、先頭ノードの削除も同じロジックで処理できる

2. **right を n 個先に進める**
   ```go
   for n > 0 {
       right = right.Next
       n--
   }
   ```

3. **両ポインタを同時に進める**
   ```go
   for right != nil {
       left = left.Next
       right = right.Next
   }
   ```

4. **ノードを削除**
   ```go
   left.Next = left.Next.Next
   ```

### 視覚的な理解

```
head = [1,2,3,4], n = 2

初期状態:
dummy → [1] → [2] → [3] → [4] → nil
  l      r

Step 1: right を n=2 個先に進める
dummy → [1] → [2] → [3] → [4] → nil
  l             r

Step 2: 両方同時に進める（rightがnilになるまで）
dummy → [1] → [2] → [3] → [4] → nil
          l             r

dummy → [1] → [2] → [3] → [4] → nil
                l             r (= nil)

right == nil → 終了

Step 3: left.Next を削除
left.Next = [3]（削除対象）
left.Next = left.Next.Next

dummy → [1] → [2] ──────→ [4] → nil
                l

return dummy.Next = [1] → [2] → [4]
```

### 先頭ノードを削除する場合

```
head = [1,2], n = 2

初期状態:
dummy → [1] → [2] → nil
  l      r

Step 1: right を n=2 個先に進める
dummy → [1] → [2] → nil
  l                 r (= nil)

Step 2: 両方同時に進める
right == nil なのでスキップ

Step 3: left.Next を削除
left.Next = [1]（削除対象 = head!）
left.Next = left.Next.Next

dummy ──────→ [2] → nil
  l

return dummy.Next = [2]
```

ダミーヘッドがあるから、先頭削除も特別扱い不要！

### なぜ left は dummy から始めるか

```
削除には「1つ前のノード」が必要

もし left = head から始めると:
- 削除対象の「上」に止まってしまう
- 1つ前のノードにアクセスできない（単方向リスト）

left = dummy から始めると:
- 削除対象の「1つ前」に止まる
- left.Next で削除対象にアクセス可能
```

### ポインタ間の距離

```
right を n 個先に進めた後:

dummy → [1] → [2] → [3] → [4] → nil
  l             r

left と right の間には n+1 個のノードがある
（left自体はdummyなので）

right が nil に達したとき:
- left は末尾から n+1 番目
- left.Next は末尾から n 番目（= 削除対象）
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | リストを1回走査 |
| 空間計算量 | O(1) | ポインタのみ使用 |

### 別解：2パスアプローチ

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 1パス目: 長さを計算
    length := 0
    curr := head
    for curr != nil {
        length++
        curr = curr.Next
    }

    // 先頭を削除する場合
    if length == n {
        return head.Next
    }

    // 2パス目: (length - n - 1) 番目まで進む
    curr = head
    for i := 0; i < length-n-1; i++ {
        curr = curr.Next
    }
    curr.Next = curr.Next.Next

    return head
}
```

| 比較 | Two Pointer | 2パス |
|------|-------------|-------|
| 走査回数 | 1回 | 2回 |
| コード複雑度 | シンプル | 先頭削除の分岐が必要 |

Two Pointerの方がエレガント。
