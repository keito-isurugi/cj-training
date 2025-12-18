# Invert Binary Tree (Easy)

## 問題内容

二分木のルート `root` が与えられたとき、木を左右反転（鏡像化）して返す。

### 例

```
Input:
        1
       / \
      2   3
     / \ / \
    4  5 6  7

Output:
        1
       / \
      3   2
     / \ / \
    7  6 5  4
```

```
Input: root = [3,2,1]
Output: [3,1,2]

Input: root = []
Output: []
```

### 制約

- `0 <= ノード数 <= 100`
- `-100 <= Node.val <= 100`

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
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left, root.Right = root.Right, root.Left

    invertTree(root.Left)
    invertTree(root.Right)

    return root
}
```

## アルゴリズムなど解説

### 基本戦略

再帰を使って、各ノードで左右の子を入れ替える。

### 動作の仕組み

1. **Base case: 空ノード**
   ```go
   if root == nil {
       return nil
   }
   ```

2. **現在のノードの左右を入れ替え**
   ```go
   root.Left, root.Right = root.Right, root.Left
   ```

3. **左右の部分木を再帰的に反転**
   ```go
   invertTree(root.Left)
   invertTree(root.Right)
   ```

### 視覚的な理解

```
元の木:
        1
       / \
      2   3
     / \
    4   5

Step 1: ルート(1)で左右入れ替え
        1
       / \
      3   2      ← 2と3が入れ替わる
     / \
    4   5

Step 2: 左部分木(3)を再帰処理
        1
       / \
      3   2
           ↑ 子がないのでそのまま

Step 3: 右部分木(2)を再帰処理
        1
       / \
      3   2
         / \
        5   4    ← 4と5が入れ替わる

完了！
```

### 再帰の流れ

```
invertTree(1)
├── swap: 2 ↔ 3
├── invertTree(3)  ← 元々の右、今は左
│   └── (子なし、何もしない)
└── invertTree(2)  ← 元々の左、今は右
    ├── swap: 4 ↔ 5
    ├── invertTree(5)
    │   └── (子なし)
    └── invertTree(4)
        └── (子なし)
```

### なぜこれで動くのか

```
ポイント:
1. 各ノードで「自分の子」を入れ替える
2. 全てのノードで同じ操作をする
3. 再帰が深いノードから戻ってくる順序は関係ない
   （各ノードは独立して処理できる）

鏡に映した木 = すべての階層で左右が反転
→ 各ノードで左右を入れ替えれば達成
```

### 入れ替えのタイミング

```go
// 前順（Pre-order）: 入れ替え → 再帰
root.Left, root.Right = root.Right, root.Left
invertTree(root.Left)
invertTree(root.Right)

// 後順（Post-order）でもOK: 再帰 → 入れ替え
invertTree(root.Left)
invertTree(root.Right)
root.Left, root.Right = root.Right, root.Left

// どちらでも同じ結果になる
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ処理 |
| 空間計算量 | O(h) | 再帰のコールスタック（h=木の高さ） |

### 別解：反復版（BFS）

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        // 左右を入れ替え
        node.Left, node.Right = node.Right, node.Left

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }

    return root
}
```

### 有名なエピソード

この問題は「Homebrew」の作者 Max Howell がGoogleの面接で解けなかったことで有名。

> "Google: 90% of our engineers use the software you wrote (Homebrew), but you can't invert a binary tree on a whiteboard so f*** off."

シンプルだが、木の再帰処理の基本を問う良問！
