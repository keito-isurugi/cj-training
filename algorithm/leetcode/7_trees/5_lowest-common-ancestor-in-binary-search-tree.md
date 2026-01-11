# Lowest Common Ancestor of a Binary Search Tree (Medium)

## 問題内容

二分探索木（BST）と2つのノード `p`, `q` が与えられたとき、両方のノードの最小共通祖先（LCA: Lowest Common Ancestor）を返す。

LCAとは、pとqの両方を子孫に持つ最も深いノード。ノード自身も自分の子孫とみなせる。

### 例

```
Input:
           5
          / \
         3   8
        / \ / \
       1  4 7  9
        \
         2

p = 3, q = 8
Output: 5
Explanation: 3と8の共通祖先は5のみ
```

```
Input: 同じ木で p = 3, q = 4
Output: 3
Explanation: 3は4の祖先であり、自身も子孫とみなせるのでLCAは3
```

### 制約

- `2 <= ノード数 <= 100`
- `-100 <= Node.val <= 100`
- `p != q`
- `p` と `q` は必ずBST内に存在

## ソースコード

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    cur := root

    for cur != nil {
        if p.Val > cur.Val && q.Val > cur.Val {
            cur = cur.Right
        } else if p.Val < cur.Val && q.Val < cur.Val {
            cur = cur.Left
        } else {
            return cur
        }
    }
    return nil
}
```

## アルゴリズムなど解説

### 基本戦略

BSTの特性（左 < 親 < 右）を利用し、pとqが「分岐」するノードを見つける。

### BSTの特性を活用

```
BSTの性質:
- 左部分木のすべてのノード < 親
- 右部分木のすべてのノード > 親

この性質から:
- p, q 両方が cur より大きい → LCAは右部分木にある
- p, q 両方が cur より小さい → LCAは左部分木にある
- そうでない（分岐点） → cur がLCA！
```

### 動作の仕組み

```go
for cur != nil {
    if p.Val > cur.Val && q.Val > cur.Val {
        cur = cur.Right   // 両方とも右側にいる
    } else if p.Val < cur.Val && q.Val < cur.Val {
        cur = cur.Left    // 両方とも左側にいる
    } else {
        return cur        // 分岐点 = LCA
    }
}
```

### 「分岐点」とは

```
3つのケース:

1. p < cur < q （pは左、qは右）
           cur ← LCA
          /   \
        p      q

2. cur == p （pが祖先）
           p ← LCA
          / \
         ?   q

3. cur == q （qが祖先）
           q ← LCA
          / \
        p    ?
```

### 視覚的な理解

```
           5
          / \
         3   8
        / \ / \
       1  4 7  9

例1: p=3, q=8

cur=5: 3 < 5, 8 > 5 → 分岐！return 5

例2: p=3, q=4

cur=5: 3 < 5, 4 < 5 → 両方左 → cur=3
cur=3: 3 == 3, 4 > 3 → 分岐！return 3

例3: p=1, q=4

cur=5: 1 < 5, 4 < 5 → 両方左 → cur=3
cur=3: 1 < 3, 4 > 3 → 分岐！return 3
```

### なぜこれで正しいのか

```
BSTの特性により:

1. p, q が両方とも cur の右にいる
   → LCAは cur ではない（curは片方の祖先にしかなれない）
   → 右に進む

2. p, q が両方とも cur の左にいる
   → LCAは cur ではない
   → 左に進む

3. p, q が cur を挟んでいる（または一方が cur と等しい）
   → cur が両方の祖先になれる最も深いノード
   → cur がLCA！
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(h) | 木の高さ分だけ探索 |
| 空間計算量 | O(1) | 変数のみ使用 |

h = 木の高さ（バランス木ならO(log n)、偏った木ならO(n)）

### 別解：再帰版

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else {
        return root
    }
}
```

### 通常の二分木との違い

| 木の種類 | LCAの探し方 | 時間計算量 |
|----------|------------|-----------|
| BST（この問題） | 値の比較で方向を決定 | O(h) |
| 通常の二分木 | 両部分木を探索 | O(n) |

BSTでは値の大小関係で探索方向が一意に決まるため、効率的！

### 通常の二分木でのLCA（参考）

```go
// 普通の二分木の場合（BSTではない）
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root  // pとqが左右に分かれている → rootがLCA
    }
    if left != nil {
        return left
    }
    return right
}
```

BSTの特性を活かすと、通常の二分木より効率的に解ける！
