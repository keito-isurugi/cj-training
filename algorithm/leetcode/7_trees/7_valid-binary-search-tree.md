# Validate Binary Search Tree (Medium)

## 問題内容

二分木の `root` が与えられたとき、有効な二分探索木（BST）であれば `true` を、そうでなければ `false` を返す。

有効なBSTの条件：
- 左部分木のすべてのノードの値 < 親の値
- 右部分木のすべてのノードの値 > 親の値
- 左右の部分木もそれぞれ有効なBST

### 例

```
Input:
    2
   / \
  1   3

Output: true
```

```
Input:
    1
   / \
  2   3

Output: false
Explanation: ルート(1)の左の子(2)が親より大きい
```

### 制約

- `1 <= ノード数 <= 1000`
- `-1000 <= Node.val <= 1000`

## ソースコード

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type QueueItem struct {
    node  *TreeNode
    left  int64
    right int64
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    queue := []QueueItem{{root, math.MinInt64, math.MaxInt64}}

    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]

        val := int64(item.node.Val)
        if val <= item.left || val >= item.right {
            return false
        }

        if item.node.Left != nil {
            queue = append(queue, QueueItem{item.node.Left, item.left, val})
        }
        if item.node.Right != nil {
            queue = append(queue, QueueItem{item.node.Right, val, item.right})
        }
    }

    return true
}
```

## アルゴリズムなど解説

### 基本戦略

各ノードに「取りうる値の範囲（left, right）」を持たせ、その範囲内にあるかをチェックする。

### よくある間違い

```
❌ 親と子だけを比較する

    5
   / \
  1   6
     / \
    3   7

親(6) > 左の子(3) → OK？

でも 3 < 5（ルート）なのでNG！
3は5の右部分木にあるので、5より大きくないといけない
```

### 正しいアプローチ：範囲を管理

```
各ノードの値は (left, right) の範囲内でなければならない

       5
      / \
     1   6
    / \
   ?   ?

ノード5: (-∞, +∞) の範囲 → 5はOK
ノード1: (-∞, 5) の範囲 → 1はOK（5より小さい）
ノード6: (5, +∞) の範囲 → 6はOK（5より大きい）
```

### 範囲の更新ルール

```go
// 左の子に行くとき
// → 現在のノードの値が新しい上限になる
QueueItem{item.node.Left, item.left, val}  // 範囲: (left, 現在の値)

// 右の子に行くとき
// → 現在のノードの値が新しい下限になる
QueueItem{item.node.Right, val, item.right}  // 範囲: (現在の値, right)
```

### 視覚的な理解

```
        5
       / \
      3   7
     / \ / \
    1  4 6  8

各ノードの有効範囲:

ノード5: (-∞, +∞) → 5はOK
├── ノード3: (-∞, 5) → 3はOK
│   ├── ノード1: (-∞, 3) → 1はOK
│   └── ノード4: (3, 5) → 4はOK
└── ノード7: (5, +∞) → 7はOK
    ├── ノード6: (5, 7) → 6はOK
    └── ノード8: (7, +∞) → 8はOK

すべてOK → return true
```

### 無効な例

```
        5
       / \
      1   6
         / \
        3   7

ノード5: (-∞, +∞) → OK
├── ノード1: (-∞, 5) → OK
└── ノード6: (5, +∞) → OK
    ├── ノード3: (5, 6) → 3 <= 5 なのでNG!
    └── ノード7: (6, +∞) → OK

return false
```

### なぜ int64 を使うか

```go
queue := []QueueItem{{root, math.MinInt64, math.MaxInt64}}
val := int64(item.node.Val)
```

```
制約: -1000 <= Node.val <= 1000

初期範囲を (-∞, +∞) にしたい
→ int32の範囲外の値が必要
→ int64 を使用

※ Node.valがint32の最小値/最大値の場合も正しく動作
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ処理 |
| 空間計算量 | O(n) | キューの最大サイズ |

### 別解：再帰版（DFS）

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, left, right int64) bool {
    if node == nil {
        return true
    }

    val := int64(node.Val)
    if val <= left || val >= right {
        return false
    }

    return validate(node.Left, left, val) && validate(node.Right, val, right)
}
```

### 別解：中順走査（In-order）

BSTを中順走査すると昇順になる性質を利用。

```go
func isValidBST(root *TreeNode) bool {
    var prev *int64
    return inorder(root, &prev)
}

func inorder(node *TreeNode, prev **int64) bool {
    if node == nil {
        return true
    }

    // 左部分木
    if !inorder(node.Left, prev) {
        return false
    }

    // 現在のノード：前の値より大きいか
    val := int64(node.Val)
    if *prev != nil && val <= **prev {
        return false
    }
    *prev = &val

    // 右部分木
    return inorder(node.Right, prev)
}
```

```
中順走査の順序:
    5
   / \
  3   7
 / \
1   4

走査順: 1 → 3 → 4 → 5 → 7
昇順になっていれば有効なBST！
```

### 解法の比較

| 解法 | 特徴 |
|------|------|
| 範囲チェック（BFS） | 直感的、各ノードの有効範囲が明確 |
| 範囲チェック（DFS） | シンプル、再帰で簡潔に書ける |
| 中順走査 | BSTの性質を活用、別の視点 |

範囲チェックの考え方は他の問題でも応用できる重要なテクニック！
