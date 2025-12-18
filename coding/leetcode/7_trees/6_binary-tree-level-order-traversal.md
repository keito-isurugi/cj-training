# Binary Tree Level Order Traversal (Medium)

## 問題内容

二分木の `root` が与えられたとき、レベル順走査（幅優先探索）の結果をネストされたリストで返す。各サブリストは同じレベルのノードの値を左から右の順で含む。

### 例

```
Input:
        1
       / \
      2   3
     / \ / \
    4  5 6  7

Output: [[1],[2,3],[4,5,6,7]]
```

```
Input: root = [1]
Output: [[1]]

Input: root = []
Output: []
```

### 制約

- `0 <= ノード数 <= 1000`
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
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }

    q := []*TreeNode{root}

    for len(q) > 0 {
        qLen := len(q)
        level := []int{}

        for i := 0; i < qLen; i++ {
            node := q[0]
            q = q[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }

        res = append(res, level)
    }

    return res
}
```

## アルゴリズムなど解説

### 基本戦略

BFS（幅優先探索）を使い、キューで各レベルのノードを管理。レベルごとに結果を配列に追加する。

### 動作の仕組み

1. **キューを初期化**
   ```go
   q := []*TreeNode{root}
   ```

2. **レベルごとに処理**
   ```go
   for len(q) > 0 {
       qLen := len(q)       // 現在のレベルのノード数を固定
       level := []int{}     // 現在のレベルの値を格納
   ```

3. **現在のレベルの全ノードを処理**
   ```go
   for i := 0; i < qLen; i++ {
       node := q[0]         // キューの先頭を取得
       q = q[1:]            // デキュー
       level = append(level, node.Val)

       // 子ノードをキューに追加
       if node.Left != nil {
           q = append(q, node.Left)
       }
       if node.Right != nil {
           q = append(q, node.Right)
       }
   }
   ```

4. **レベルの結果を追加**
   ```go
   res = append(res, level)
   ```

### 視覚的な理解

```
        1           Level 0
       / \
      2   3         Level 1
     / \ / \
    4  5 6  7       Level 2

初期: q = [1], res = []

━━━ Level 0 ━━━
qLen = 1, level = []
- node=1 を処理 → level = [1]
- 子(2,3)をキューに追加 → q = [2, 3]
res = [[1]]

━━━ Level 1 ━━━
qLen = 2, level = []
- node=2 を処理 → level = [2]
- 子(4,5)をキューに追加 → q = [3, 4, 5]
- node=3 を処理 → level = [2, 3]
- 子(6,7)をキューに追加 → q = [4, 5, 6, 7]
res = [[1], [2, 3]]

━━━ Level 2 ━━━
qLen = 4, level = []
- node=4 を処理 → level = [4]
- node=5 を処理 → level = [4, 5]
- node=6 を処理 → level = [4, 5, 6]
- node=7 を処理 → level = [4, 5, 6, 7]
- 子なし → q = []
res = [[1], [2, 3], [4, 5, 6, 7]]

q が空 → 終了
```

### なぜ qLen を先に固定するか

```go
qLen := len(q)  // ここで固定！

for i := 0; i < qLen; i++ {
    // 処理中にqに追加しても
    // このループは現在のレベル分だけ回る
}
```

```
もし固定しないと:
q = [2, 3]
- 2を処理 → q = [3, 4, 5]
- 3を処理 → q = [4, 5, 6, 7]
- 4を処理 → ... ← 次のレベルも同じループで処理されてしまう！

qLen を固定することで:
q = [2, 3], qLen = 2
- 2を処理 → q = [3, 4, 5]
- 3を処理 → q = [4, 5, 6, 7]
- ループ終了（2回で終わり）
res に [2, 3] を追加
```

### キューの操作

```go
// デキュー（先頭を取り出す）
node := q[0]
q = q[1:]

// エンキュー（末尾に追加）
q = append(q, node.Left)
```

Goではスライスでキューを実装できる。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ処理 |
| 空間計算量 | O(n) | キュー + 結果配列 |

### 別解：再帰版（DFS）

```go
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    dfs(root, 0, &res)
    return res
}

func dfs(node *TreeNode, level int, res *[][]int) {
    if node == nil {
        return
    }

    // 新しいレベルなら配列を追加
    if level >= len(*res) {
        *res = append(*res, []int{})
    }

    // 現在のレベルに値を追加
    (*res)[level] = append((*res)[level], node.Val)

    // 子ノードを再帰処理（レベル+1）
    dfs(node.Left, level+1, res)
    dfs(node.Right, level+1, res)
}
```

```
再帰の流れ:
dfs(1, level=0) → res = [[1]]
├── dfs(2, level=1) → res = [[1], [2]]
│   ├── dfs(4, level=2) → res = [[1], [2], [4]]
│   └── dfs(5, level=2) → res = [[1], [2], [4, 5]]
└── dfs(3, level=1) → res = [[1], [2, 3], [4, 5]]
    ├── dfs(6, level=2) → res = [[1], [2, 3], [4, 5, 6]]
    └── dfs(7, level=2) → res = [[1], [2, 3], [4, 5, 6, 7]]
```

### BFS vs DFS の比較

| 観点 | BFS（この解法） | DFS（再帰） |
|------|---------------|------------|
| 直感性 | レベルごとに処理 | 深さ優先で処理 |
| 順序 | 自然に左→右 | 左→右（引数で制御） |
| 実装 | キュー使用 | 再帰 |
| 面接 | **推奨**（BFSの典型例） | 代替解として有効 |

Level Order Traversal は**BFSの最も基本的な応用**なので、確実に書けるようにしておくべき！
