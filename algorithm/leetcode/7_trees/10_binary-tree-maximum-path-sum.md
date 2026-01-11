# Binary Tree Maximum Path Sum (Hard)

## 問題内容

非空の二分木の `root` が与えられたとき、任意の非空パスの最大パス和を返す。

パスとは、隣接するノード同士が辺で接続されているノードの列。各ノードはパスに一度しか登場できない。パスは必ずしもルートを含む必要はない。

### 例

```
Input: root = [1,2,3]

    1
   / \
  2   3

Output: 6
Explanation: パス 2 → 1 → 3 の和 = 2 + 1 + 3 = 6
```

```
Input: root = [-15,10,20,null,null,15,5,-5]

       -15
       /  \
     10    20
          /  \
         15   5
             /
           -5

Output: 40
Explanation: パス 15 → 20 → 5 の和 = 15 + 20 + 5 = 40
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
func maxPathSum(root *TreeNode) int {
    res := []int{root.Val}

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftMax := dfs(node.Left)
        rightMax := dfs(node.Right)

        leftMax = max(leftMax, 0)
        rightMax = max(rightMax, 0)

        res[0] = max(res[0], node.Val+leftMax+rightMax)

        return node.Val + max(leftMax, rightMax)
    }

    dfs(root)
    return res[0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## アルゴリズムなど解説

### 基本戦略

DFS（深さ優先探索）で各ノードを訪問し、そのノードを「経由点」とした最大パス和を計算。グローバルな最大値を更新していく。

### パスの種類

```
パスには2種類ある:

1. 「経由パス」: ノードを経由して左右に分岐
       node
      /    \
   left    right
   パス: left → node → right

2. 「片側パス」: ノードから片側のみに伸びる
       node
      /
   left
   パス: left → node（親に渡せる）
```

### 核心的なアイデア

```
各ノードで2つのことを行う:

1. 「経由パス」の和を計算 → グローバル最大を更新
   node.Val + leftMax + rightMax

2. 「片側パス」の最大を返す → 親ノードで使用
   node.Val + max(leftMax, rightMax)
```

### なぜ「経由パス」は親に返せないか

```
        1
       / \
      2   3
     /
    4

もし 4 → 2 → 1 → 3 を1つのパスとして親に渡そうとすると...

   親
    |
    1         ← 親から見ると1に来るが
   / \
  2   3       ← 2にも3にも行く必要がある
 /
4

パスは「折り返せない」ので、これは不可能！
だから経由パスはそのノードで「完結」し、
片側パスだけを親に返す。
```

### 負の値の処理

```go
leftMax = max(leftMax, 0)
rightMax = max(rightMax, 0)
```

```
子からの最大が負なら、その部分木は使わない（0と比較）

例:
    5
   / \
 -3   2

左の子は -3（負）→ 使わない方が良い
最大パス = 5 + 0 + 2 = 7
（5 → 2 のパス）
```

### 視覚的な理解

```
       -15
       /  \
     10    20
          /  \
         15   5
             /
           -5

━━━ DFS探索（後順走査） ━━━

Step 1: ノード -5
- leftMax = 0, rightMax = 0
- 経由パス = -5 + 0 + 0 = -5
- 片側パス = -5 + 0 = -5
- res = max(-15, -5) = -5
- return -5 → 負なので親は0として扱う

Step 2: ノード 5
- leftMax = max(-5, 0) = 0
- rightMax = 0
- 経由パス = 5 + 0 + 0 = 5
- 片側パス = 5 + 0 = 5
- res = max(-5, 5) = 5
- return 5

Step 3: ノード 15
- leftMax = 0, rightMax = 0
- 経由パス = 15 + 0 + 0 = 15
- 片側パス = 15 + 0 = 15
- res = max(5, 15) = 15
- return 15

Step 4: ノード 20
- leftMax = 15, rightMax = 5
- 経由パス = 20 + 15 + 5 = 40 ← 最大！
- 片側パス = 20 + max(15, 5) = 35
- res = max(15, 40) = 40
- return 35

Step 5: ノード 10
- leftMax = 0, rightMax = 0
- 経由パス = 10 + 0 + 0 = 10
- 片側パス = 10 + 0 = 10
- res = max(40, 10) = 40
- return 10

Step 6: ノード -15（ルート）
- leftMax = max(10, 0) = 10
- rightMax = max(35, 0) = 35
- 経由パス = -15 + 10 + 35 = 30
- 片側パス = -15 + 35 = 20
- res = max(40, 30) = 40

最終結果: 40
```

### なぜ res をスライスで持つか

```go
res := []int{root.Val}
```

```
Goのクロージャで外部変数を更新するため:

❌ 値渡しでは更新が反映されない
res := root.Val
dfs = func(node *TreeNode) int {
    res = max(res, ...)  // ローカルな変更
}

✅ スライス（参照型）なら更新が反映される
res := []int{root.Val}
dfs = func(node *TreeNode) int {
    res[0] = max(res[0], ...)  // 元のスライスを更新
}

✅ ポインタでもOK
var res *int
*res = max(*res, ...)
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ訪問 |
| 空間計算量 | O(h) | 再帰スタックの深さ（木の高さ） |

### 別解：グローバル変数版

```go
var maxSum int

func maxPathSum(root *TreeNode) int {
    maxSum = root.Val
    dfs(root)
    return maxSum
}

func dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }

    leftMax := max(dfs(node.Left), 0)
    rightMax := max(dfs(node.Right), 0)

    // 経由パスでグローバル最大を更新
    maxSum = max(maxSum, node.Val+leftMax+rightMax)

    // 片側パスを親に返す
    return node.Val + max(leftMax, rightMax)
}
```

### 重要なポイントまとめ

```
1. 各ノードで「経由パス」と「片側パス」を区別する

2. 経由パス（左 + ノード + 右）
   → グローバル最大の更新に使用
   → 親には返せない（折り返せないから）

3. 片側パス（ノード + 左or右の大きい方）
   → 親ノードに返す
   → 親がこれを使ってさらに大きなパスを作る

4. 負の部分木は使わない（max(子, 0)）
   → 負の値を足すとパス和が減るから
```

### 関連問題

| 問題 | 共通点 |
|------|--------|
| Maximum Depth of Binary Tree | DFSで各ノードの値を計算 |
| Diameter of Binary Tree | 左右の最大を合わせる考え方 |
| Path Sum | パスの和を求める |

この問題は**Hard**だが、「経由パス」と「片側パス」の区別を理解すれば解ける！
