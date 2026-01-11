# Binary Tree 基礎

## Binary Treeとは

各ノードが最大2つの子ノード（左と右）を持つツリー構造。

```
        1          ← ルート（root）
       / \
      2   3        ← 子ノード
     / \ / \
    4  5 6  7      ← 葉ノード（leaf）
```

## 基本構造

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode   // 左の子
    Right *TreeNode   // 右の子
}
```

## 用語

| 用語 | 説明 |
|------|------|
| Root | 最上位のノード |
| Leaf | 子を持たないノード |
| Parent | 親ノード |
| Child | 子ノード |
| Height | ルートから最も深い葉までの距離 |
| Depth | ルートからそのノードまでの距離 |

## 木の種類

### 完全二分木（Complete Binary Tree）

```
最後のレベル以外は完全に埋まっている
最後のレベルは左から詰まっている

        1
       / \
      2   3
     / \
    4   5
```

### 満二分木（Full Binary Tree）

```
すべてのノードが0または2つの子を持つ

        1
       / \
      2   3
     / \
    4   5
```

### 二分探索木（Binary Search Tree）

```
左の子 < 親 < 右の子 の関係が成り立つ

        5
       / \
      3   7
     / \ / \
    2  4 6  8
```

## 走査（Traversal）

### 深さ優先探索（DFS）

#### 前順（Pre-order）: Root → Left → Right

```go
func preorder(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Println(root.Val)  // 1. 現在のノード
    preorder(root.Left)    // 2. 左部分木
    preorder(root.Right)   // 3. 右部分木
}

        1
       / \
      2   3
     / \
    4   5

結果: [1, 2, 4, 5, 3]
```

#### 中順（In-order）: Left → Root → Right

```go
func inorder(root *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left)     // 1. 左部分木
    fmt.Println(root.Val)  // 2. 現在のノード
    inorder(root.Right)    // 3. 右部分木
}

結果: [4, 2, 5, 1, 3]
※ BSTの場合、ソート順になる
```

#### 後順（Post-order）: Left → Right → Root

```go
func postorder(root *TreeNode) {
    if root == nil {
        return
    }
    postorder(root.Left)   // 1. 左部分木
    postorder(root.Right)  // 2. 右部分木
    fmt.Println(root.Val)  // 3. 現在のノード
}

結果: [4, 5, 2, 3, 1]
```

### 幅優先探索（BFS）/ レベル順走査

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        level := []int{}
        size := len(queue)

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, level)
    }
    return result
}

        1
       / \
      2   3
     / \
    4   5

結果: [[1], [2, 3], [4, 5]]
```

## 再帰の基本パターン

### テンプレート

```go
func traverse(root *TreeNode) {
    // Base case
    if root == nil {
        return
    }

    // 現在のノードの処理
    // ...

    // 再帰呼び出し
    traverse(root.Left)
    traverse(root.Right)
}
```

### 値を返す場合

```go
func dfs(root *TreeNode) int {
    if root == nil {
        return 0  // Base case
    }

    left := dfs(root.Left)
    right := dfs(root.Right)

    // 左右の結果を使って計算
    return 1 + left + right  // 例: ノード数
}
```

## よく使うテクニック

### 1. 高さの計算

```go
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(height(root.Left), height(root.Right))
}
```

### 2. ノード数のカウント

```go
func count(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + count(root.Left) + count(root.Right)
}
```

### 3. 葉ノードの判定

```go
func isLeaf(node *TreeNode) bool {
    return node != nil && node.Left == nil && node.Right == nil
}
```

## 計算量まとめ

| 操作 | 時間計算量 | 空間計算量（再帰） |
|------|-----------|-------------------|
| 走査 | O(n) | O(h)* |
| 検索（BST） | O(h) | O(h) |
| 挿入（BST） | O(h) | O(h) |
| 削除（BST） | O(h) | O(h) |

\* h = 木の高さ、最悪O(n)、バランス木でO(log n)
