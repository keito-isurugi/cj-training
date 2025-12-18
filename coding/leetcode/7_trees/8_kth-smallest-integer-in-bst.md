# Kth Smallest Element in a BST (Medium)

## 問題内容

二分探索木（BST）の `root` と整数 `k` が与えられたとき、木の中でk番目に小さい値（1-indexed）を返す。

### 例

```
Input:
    2
   / \
  1   3

k = 1
Output: 1
Explanation: 最小値は1
```

```
Input:
      4
     / \
    3   5
   /
  2

k = 4
Output: 5
Explanation: 昇順: [2,3,4,5] → 4番目は5
```

### 制約

- `1 <= k <= ノード数 <= 1000`
- `0 <= Node.val <= 1000`

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
func kthSmallest(root *TreeNode, k int) int {
    curr := root
    for {
        if curr.Left == nil {
            k--
            if k == 0 {
                return curr.Val
            }
            curr = curr.Right
        } else {
            pred := curr.Left
            for pred.Right != nil && pred.Right != curr {
                pred = pred.Right
            }
            if pred.Right == nil {
                pred.Right = curr
                curr = curr.Left
            } else {
                pred.Right = nil
                k--
                if k == 0 {
                    return curr.Val
                }
                curr = curr.Right
            }
        }
    }
}
```

## アルゴリズムなど解説

### 基本戦略

Morris Traversal（モリス走査）を使用して、O(1)の空間計算量で中順走査を行い、k番目の要素を見つける。

### BSTの重要な性質

```
BSTを中順走査（In-order）すると昇順になる

      4
     / \
    2   6
   / \ / \
  1  3 5  7

中順走査: 1 → 2 → 3 → 4 → 5 → 6 → 7
         ↑
        昇順！

k番目に小さい = 中順走査でk番目に訪問するノード
```

### Morris Traversal とは

スタックや再帰を使わずに木を走査するアルゴリズム。一時的にポインタを変更して「戻る道」を作る。

```
通常の中順走査:
- 再帰 or スタックで「戻る場所」を記憶
- 空間計算量 O(h)（木の高さ）

Morris Traversal:
- ポインタを一時的に変更して「戻る道」を作る
- 空間計算量 O(1)
```

### Morris Traversalの動作

```go
// 左の子がない場合
if curr.Left == nil {
    // currを処理（k番目かチェック）
    k--
    if k == 0 {
        return curr.Val
    }
    curr = curr.Right  // 右に進む
}
```

```go
// 左の子がある場合
else {
    // currの「中順先行者」を見つける
    // = 左部分木の最右ノード
    pred := curr.Left
    for pred.Right != nil && pred.Right != curr {
        pred = pred.Right
    }

    if pred.Right == nil {
        // 初めて訪問: 戻るリンクを作成
        pred.Right = curr
        curr = curr.Left
    } else {
        // 2回目の訪問: リンクを削除して処理
        pred.Right = nil
        k--
        if k == 0 {
            return curr.Val
        }
        curr = curr.Right
    }
}
```

### 視覚的な理解

```
      4
     / \
    2   6
   / \
  1   3

k = 3 を探す

Step 1: curr=4
- 左の子(2)あり
- pred = 4の中順先行者を探す → 3
- 3.Right = nil → リンク作成: 3.Right = 4
- curr = 2

      4 ←──┐
     / \    │ (一時リンク)
    2   6   │
   / \      │
  1   3 ───┘

Step 2: curr=2
- 左の子(1)あり
- pred = 2の中順先行者 → 1
- 1.Right = nil → リンク作成: 1.Right = 2
- curr = 1

Step 3: curr=1
- 左の子なし
- k-- → k=2
- curr = 1.Right = 2

Step 4: curr=2
- 左の子(1)あり
- pred = 1, 1.Right = 2（curr）→ 2回目の訪問
- リンク削除: 1.Right = nil
- k-- → k=1
- curr = 3

Step 5: curr=3
- 左の子なし
- k-- → k=0
- return 3 ← 答え！
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各ノードを最大2回訪問 |
| 空間計算量 | O(1) | 追加の空間不要（ポインタのみ） |

### 別解：再帰版（シンプル）

```go
func kthSmallest(root *TreeNode, k int) int {
    result := 0
    count := 0

    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }

        inorder(node.Left)

        count++
        if count == k {
            result = node.Val
            return
        }

        inorder(node.Right)
    }

    inorder(root)
    return result
}
```

### 別解：スタック版（反復）

```go
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    curr := root

    for {
        // 左に行けるだけ行く
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        // スタックからpop
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // k番目か確認
        k--
        if k == 0 {
            return curr.Val
        }

        // 右に進む
        curr = curr.Right
    }
}
```

### 解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|-----------|-----------|------|
| Morris Traversal | O(n) | **O(1)** | 最も空間効率が良い |
| 再帰 | O(n) | O(h) | シンプル |
| スタック | O(n) | O(h) | 反復的で制御しやすい |

**面接ではスタック版や再帰版で十分**。Morris Traversalは「O(1)空間で解けるか」と聞かれた場合の最適化として知っておくと良い。
