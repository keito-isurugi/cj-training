# Same Tree (Easy)

## 問題内容

2つの二分木 `p` と `q` が与えられたとき、両方の木が等価（同じ構造で同じ値を持つ）であれば `true` を、そうでなければ `false` を返す。

### 例

```
Input:
p:      1           q:      1
       / \                 / \
      2   3               2   3

Output: true
```

```
Input:
p:      4           q:      4
       /                     \
      7                       7

Output: false
Explanation: 構造が異なる（左 vs 右）
```

```
Input:
p:      1           q:      1
       / \                 / \
      2   3               3   2

Output: false
Explanation: 左右の値が入れ替わっている
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
    queue1 := []*TreeNode{p}
    queue2 := []*TreeNode{q}

    for len(queue1) > 0 && len(queue2) > 0 {
        for i := len(queue1); i > 0; i-- {
            nodeP := queue1[0]
            nodeQ := queue2[0]
            queue1 = queue1[1:]
            queue2 = queue2[1:]

            if nodeP == nil && nodeQ == nil {
                continue
            }
            if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
                return false
            }

            queue1 = append(queue1, nodeP.Left, nodeP.Right)
            queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
        }
    }

    return len(queue1) == 0 && len(queue2) == 0
}
```

## アルゴリズムなど解説

### 基本戦略

BFS（幅優先探索）で2つの木を同時に走査し、各ノードを比較する。

### 動作の仕組み

1. **2つのキューを初期化**
   ```go
   queue1 := []*TreeNode{p}
   queue2 := []*TreeNode{q}
   ```

2. **同時に走査して比較**
   ```go
   nodeP := queue1[0]
   nodeQ := queue2[0]
   ```

3. **3つのケースをチェック**
   ```go
   // 両方nil → OK、次へ
   if nodeP == nil && nodeQ == nil {
       continue
   }
   // 片方だけnil、または値が違う → NG
   if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
       return false
   }
   ```

4. **子ノードをキューに追加**
   ```go
   queue1 = append(queue1, nodeP.Left, nodeP.Right)
   queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
   ```

### 視覚的な理解

```
p:      1           q:      1
       / \                 / \
      2   3               2   3

初期: queue1 = [1], queue2 = [1]

ラウンド1:
  nodeP = 1, nodeQ = 1 → 値が同じ ✓
  queue1 = [2, 3], queue2 = [2, 3]

ラウンド2:
  nodeP = 2, nodeQ = 2 → 値が同じ ✓
  queue1 = [3, nil, nil], queue2 = [3, nil, nil]

  nodeP = 3, nodeQ = 3 → 値が同じ ✓
  queue1 = [nil, nil, nil, nil], queue2 = [nil, nil, nil, nil]

ラウンド3:
  全てnil → continue

両方空 → return true
```

### 失敗例

```
p:      4           q:      4
       /                     \
      7                       7

初期: queue1 = [4], queue2 = [4]

ラウンド1:
  nodeP = 4, nodeQ = 4 → 値が同じ ✓
  queue1 = [7, nil], queue2 = [nil, 7]

ラウンド2:
  nodeP = 7, nodeQ = nil → 片方だけnil！
  return false
```

### なぜnilも追加するか

```go
queue1 = append(queue1, nodeP.Left, nodeP.Right)
// nilも追加する！

理由:
- 構造の違いを検出するため
- [7, nil] と [nil, 7] は違う構造

もしnilを追加しないと:
- 両方 [7] になってしまい
- 構造の違いが検出できない
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ比較 |
| 空間計算量 | O(n) | キューの最大サイズ |

### 別解：再帰版（DFS）

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    // 両方nil
    if p == nil && q == nil {
        return true
    }
    // 片方だけnil、または値が違う
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }
    // 左右の部分木を再帰的に比較
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

```
再帰の流れ:

isSameTree(1, 1)
├── 1 == 1 ✓
├── isSameTree(2, 2)
│   ├── 2 == 2 ✓
│   ├── isSameTree(nil, nil) = true
│   └── isSameTree(nil, nil) = true
│   └── return true && true = true
└── isSameTree(3, 3)
    ├── 3 == 3 ✓
    ├── isSameTree(nil, nil) = true
    └── isSameTree(nil, nil) = true
    └── return true && true = true
└── return true && true = true
```

### BFS vs DFS の比較

| 観点 | BFS（この解法） | DFS（再帰） |
|------|---------------|------------|
| コード量 | やや多い | **シンプル** |
| 直感性 | レベルごとに比較 | 構造的に比較 |
| 早期終了 | 可能 | 可能 |

**DFS（再帰）の方がシンプルでおすすめ**。わずか数行で書ける。
