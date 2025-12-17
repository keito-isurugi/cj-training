# Merge Two Sorted Lists (Easy)

## 問題内容

2つのソート済みリンクリスト `list1` と `list2` が与えられたとき、これらをマージして1つのソート済みリストを返す。

### 例

```
Input: list1 = [1,2,4], list2 = [1,3,5]
Output: [1,1,2,3,4,5]

Input: list1 = [], list2 = [1,2]
Output: [1,2]

Input: list1 = [], list2 = []
Output: []
```

### 制約

- `0 <= 各リストの長さ <= 100`
- `-100 <= Node.val <= 100`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    node := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            node.Next = list1
            list1 = list1.Next
        } else {
            node.Next = list2
            list2 = list2.Next
        }
        node = node.Next
    }

    node.Next = list1
    if list1 == nil {
        node.Next = list2
    }

    return dummy.Next
}
```

## アルゴリズムなど解説

### 基本戦略

ダミーヘッドを使い、2つのリストを先頭から比較しながら小さい方を繋げていく。

### 動作の仕組み

1. **ダミーヘッドを作成**
   ```go
   dummy := &ListNode{}
   node := dummy
   ```
   - `dummy`: 結果リストの先頭の前（番兵）
   - `node`: 結果リストの現在位置

2. **両リストを比較しながらマージ**
   ```go
   for list1 != nil && list2 != nil {
       if list1.Val < list2.Val {
           node.Next = list1
           list1 = list1.Next
       } else {
           node.Next = list2
           list2 = list2.Next
       }
       node = node.Next
   }
   ```

3. **残りを繋げる**
   ```go
   node.Next = list1
   if list1 == nil {
       node.Next = list2
   }
   ```
   - どちらかが先に終わったら、残りをそのまま繋げる

### 視覚的な理解

```
list1: 1 → 2 → 4 → nil
list2: 1 → 3 → 5 → nil
dummy: [X] →
       node
```

#### ステップ1: 1 vs 1

```
list1.Val(1) >= list2.Val(1) → list2を選択

list1: 1 → 2 → 4 → nil
list2:     3 → 5 → nil
dummy: [X] → [1] →
             node
```

#### ステップ2: 1 vs 3

```
list1.Val(1) < list2.Val(3) → list1を選択

list1:     2 → 4 → nil
list2:     3 → 5 → nil
dummy: [X] → [1] → [1] →
                   node
```

#### ステップ3: 2 vs 3

```
list1.Val(2) < list2.Val(3) → list1を選択

list1:         4 → nil
list2:     3 → 5 → nil
dummy: [X] → [1] → [1] → [2] →
                         node
```

#### ステップ4: 4 vs 3

```
list1.Val(4) >= list2.Val(3) → list2を選択

list1:         4 → nil
list2:             5 → nil
dummy: [X] → [1] → [1] → [2] → [3] →
                               node
```

#### ステップ5: 4 vs 5

```
list1.Val(4) < list2.Val(5) → list1を選択

list1:             nil
list2:             5 → nil
dummy: [X] → [1] → [1] → [2] → [3] → [4] →
                                     node
```

#### ステップ6: list1が終了

```
list1 == nil → ループ終了
残りのlist2を繋げる

dummy: [X] → [1] → [1] → [2] → [3] → [4] → [5] → nil

return dummy.Next = [1] → [1] → [2] → [3] → [4] → [5]
```

### なぜダミーヘッドを使うか

```
ダミーヘッドなしの場合:
- 最初のノードを決めるための特別な処理が必要
- headがnilかどうかの分岐が必要

ダミーヘッドありの場合:
- 常にnode.Nextに繋げるだけ
- 最後にdummy.Nextを返せば良い
- コードがシンプルになる
```

### 残りを繋げる処理の意味

```go
node.Next = list1
if list1 == nil {
    node.Next = list2
}
```

```
ループ終了時:
- list1が残っている → そのまま繋げる
- list2が残っている → list1はnilなのでlist2を繋げる
- 両方nil → node.Next = nil（何もしなくてOK）

// より簡潔な書き方
if list1 != nil {
    node.Next = list1
} else {
    node.Next = list2
}
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n + m) | 両リストを1回ずつ走査 |
| 空間計算量 | O(1) | ポインタのみ使用（新しいノードを作らない） |

### 別解：再帰版

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }

    if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}
```

- 空間計算量: O(n + m)（コールスタック）
- 反復版の方がメモリ効率が良い
