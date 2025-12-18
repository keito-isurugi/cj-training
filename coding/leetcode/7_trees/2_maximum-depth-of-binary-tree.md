# Maximum Depth of Binary Tree (Easy)

## 問題内容

二分木のルート `root` が与えられたとき、木の深さ（最大深度）を返す。

深さは、ルートから最も遠い葉ノードまでのパス上のノード数として定義される。

### 例

```
Input:
        1
       / \
      2   3
           \
            4

Output: 3
Explanation: 1 → 3 → 4 のパスで深さ3
```

```
Input: root = []
Output: 0
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
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    q := linkedlistqueue.New()
    q.Enqueue(root)
    level := 0

    for !q.Empty() {
        size := q.Size()

        for i := 0; i < size; i++ {
            val, _ := q.Dequeue()
            node := val.(*TreeNode)

            if node.Left != nil {
                q.Enqueue(node.Left)
            }
            if node.Right != nil {
                q.Enqueue(node.Right)
            }
        }
        level++
    }

    return level
}
```

## アルゴリズムなど解説

### 基本戦略

BFS（幅優先探索）を使って、レベル（階層）ごとに処理し、レベル数をカウントする。

### 動作の仕組み

1. **キューを初期化**
   ```go
   q := linkedlistqueue.New()
   q.Enqueue(root)
   level := 0
   ```

2. **レベルごとに処理**
   ```go
   for !q.Empty() {
       size := q.Size()  // 現在のレベルのノード数

       for i := 0; i < size; i++ {
           // 現在のレベルの全ノードを処理
       }
       level++  // 1レベル完了
   }
   ```

3. **子ノードをキューに追加**
   ```go
   if node.Left != nil {
       q.Enqueue(node.Left)
   }
   if node.Right != nil {
       q.Enqueue(node.Right)
   }
   ```

### 視覚的な理解

```
        1         Level 1
       / \
      2   3       Level 2
           \
            4     Level 3

BFSの流れ:

初期: queue = [1], level = 0

ラウンド1: size = 1
  - 1を処理 → 2, 3をキューに追加
  - queue = [2, 3], level = 1

ラウンド2: size = 2
  - 2を処理 → 子なし
  - 3を処理 → 4をキューに追加
  - queue = [4], level = 2

ラウンド3: size = 1
  - 4を処理 → 子なし
  - queue = [], level = 3

queue.Empty() → 終了
return level = 3
```

### なぜ size を先に取得するか

```go
size := q.Size()  // ここで固定！

for i := 0; i < size; i++ {
    // 処理中にキューに追加しても
    // このループは現在のレベル分だけ回る
}
```

```
もし size を固定しないと:
queue = [2, 3]
- 2を処理 → queue = [3]
- 3を処理 → queue = [4]
- 4を処理 → queue = []  ← 次のレベルも同じループで処理されてしまう！

size を固定することで:
queue = [2, 3], size = 2
- 2を処理 → queue = [3]
- 3を処理 → queue = [4]
- ループ終了（2回で終わり）
level++
- 次のループで4を処理
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ処理 |
| 空間計算量 | O(w) | キューの最大サイズ（w=木の最大幅） |

### 別解：再帰版（DFS）

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    return 1 + max(left, right)
}
```

```
再帰の流れ:

maxDepth(1)
├── maxDepth(2)
│   ├── maxDepth(nil) = 0
│   └── maxDepth(nil) = 0
│   └── return 1 + max(0, 0) = 1
└── maxDepth(3)
    ├── maxDepth(nil) = 0
    └── maxDepth(4)
        ├── maxDepth(nil) = 0
        └── maxDepth(nil) = 0
        └── return 1 + max(0, 0) = 1
    └── return 1 + max(0, 1) = 2
└── return 1 + max(1, 2) = 3
```

### BFS vs DFS の比較

| 観点 | BFS（この解法） | DFS（再帰） |
|------|---------------|------------|
| コード量 | やや多い | シンプル |
| 空間計算量 | O(w) 木の幅 | O(h) 木の高さ |
| 使用データ構造 | キュー | コールスタック |
| 直感性 | レベルを数える | 高さを計算 |

どちらも正解だが、**DFS（再帰）の方がシンプル**で一般的。
