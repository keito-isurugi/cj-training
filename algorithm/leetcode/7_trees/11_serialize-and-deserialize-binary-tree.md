# Serialize and Deserialize Binary Tree (Hard)

## 問題内容

二分木をシリアライズ・デシリアライズするアルゴリズムを実装する。

シリアライズとは、メモリ上の構造を一連のビット列に変換し、保存やネットワーク越しの送信後に別の環境で再構築できるようにするプロセス。

二分木を文字列にシリアライズし、その文字列から元の木構造をデシリアライズできれば良い。形式は自由。

### 例

```
Input: root = [1,2,3,null,null,4,5]

      1
     / \
    2   3
       / \
      4   5

Output: [1,2,3,null,null,4,5]
```

```
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

type Codec struct{}

func Constructor() Codec {
    return Codec{}
}

// Encodes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "N"
    }
    var res []string
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node == nil {
            res = append(res, "N")
        } else {
            res = append(res, strconv.Itoa(node.Val))
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
    }

    return strings.Join(res, ",")
}

// Decodes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    vals := strings.Split(data, ",")
    if vals[0] == "N" {
        return nil
    }

    rootVal, _ := strconv.Atoi(vals[0])
    root := &TreeNode{Val: rootVal}
    queue := []*TreeNode{root}
    index := 1

    for len(queue) > 0 && index < len(vals) {
        node := queue[0]
        queue = queue[1:]

        if vals[index] != "N" {
            leftVal, _ := strconv.Atoi(vals[index])
            node.Left = &TreeNode{Val: leftVal}
            queue = append(queue, node.Left)
        }
        index++

        if index < len(vals) && vals[index] != "N" {
            rightVal, _ := strconv.Atoi(vals[index])
            node.Right = &TreeNode{Val: rightVal}
            queue = append(queue, node.Right)
        }
        index++
    }

    return root
}
```

## アルゴリズムなど解説

### 基本戦略

BFS（幅優先探索）でレベル順にノードを処理。`null` ノードも明示的に記録することで、木の構造を完全に保存する。

### シリアライズの仕組み

```
      1
     / \
    2   3
       / \
      4   5

BFSでレベル順に走査:

Level 0: 1
Level 1: 2, 3
Level 2: N, N, 4, 5  (2の子はnull)
Level 3: N, N, N, N  (4と5の子はnull)

結果: "1,2,3,N,N,4,5,N,N,N,N"
     （末尾のnullは省略可能な実装もある）
```

### なぜ null を記録するか

```
nullを記録しないと構造が曖昧になる:

木A:          木B:
    1            1
   /            / \
  2            2   3
 /
3

両方とも "1,2,3" になってしまう！

nullを記録すると:
木A: "1,2,N,3,N,N,N"
木B: "1,2,3,N,N,N,N"

区別できる！
```

### serialize の動作

```go
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "N"
    }
    var res []string
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node == nil {
            res = append(res, "N")  // nullは"N"として記録
        } else {
            res = append(res, strconv.Itoa(node.Val))
            queue = append(queue, node.Left)   // 子をキューに追加
            queue = append(queue, node.Right)  // (nullも追加される)
        }
    }

    return strings.Join(res, ",")
}
```

### 視覚的な理解（serialize）

```
      1
     / \
    2   3
       / \
      4   5

初期: queue = [1], res = []

Step 1: node = 1
- res = ["1"]
- queue = [2, 3]

Step 2: node = 2
- res = ["1", "2"]
- queue = [3, nil, nil]  ← 2の子(null)を追加

Step 3: node = 3
- res = ["1", "2", "3"]
- queue = [nil, nil, 4, 5]

Step 4: node = nil
- res = ["1", "2", "3", "N"]
- (nullなので子は追加しない)

Step 5: node = nil
- res = ["1", "2", "3", "N", "N"]

Step 6: node = 4
- res = ["1", "2", "3", "N", "N", "4"]
- queue = [5, nil, nil]

Step 7: node = 5
- res = ["1", "2", "3", "N", "N", "4", "5"]
- queue = [nil, nil, nil, nil]

... (残りのnullを処理)

結果: "1,2,3,N,N,4,5,N,N,N,N"
```

### deserialize の動作

```go
func (this *Codec) deserialize(data string) *TreeNode {
    vals := strings.Split(data, ",")
    if vals[0] == "N" {
        return nil
    }

    rootVal, _ := strconv.Atoi(vals[0])
    root := &TreeNode{Val: rootVal}
    queue := []*TreeNode{root}
    index := 1

    for len(queue) > 0 && index < len(vals) {
        node := queue[0]
        queue = queue[1:]

        // 左の子を処理
        if vals[index] != "N" {
            leftVal, _ := strconv.Atoi(vals[index])
            node.Left = &TreeNode{Val: leftVal}
            queue = append(queue, node.Left)
        }
        index++

        // 右の子を処理
        if index < len(vals) && vals[index] != "N" {
            rightVal, _ := strconv.Atoi(vals[index])
            node.Right = &TreeNode{Val: rightVal}
            queue = append(queue, node.Right)
        }
        index++
    }

    return root
}
```

### 視覚的な理解（deserialize）

```
data = "1,2,3,N,N,4,5"
vals = ["1", "2", "3", "N", "N", "4", "5"]

Step 1: root = TreeNode{Val: 1}
queue = [1], index = 1

Step 2: node = 1
- vals[1] = "2" → node.Left = TreeNode{Val: 2}
- vals[2] = "3" → node.Right = TreeNode{Val: 3}
- queue = [2, 3], index = 3

        1
       / \
      2   3

Step 3: node = 2
- vals[3] = "N" → node.Left = nil
- vals[4] = "N" → node.Right = nil
- queue = [3], index = 5

Step 4: node = 3
- vals[5] = "4" → node.Left = TreeNode{Val: 4}
- vals[6] = "5" → node.Right = TreeNode{Val: 5}
- queue = [4, 5], index = 7

        1
       / \
      2   3
         / \
        4   5

結果: 元の木が復元された！
```

### なぜ BFS が適しているか

```
BFS (この解法):
- レベル順に処理
- 親子関係が明確（2つずつ子を処理）
- 実装がシンプル

DFS:
- 前順走査で実装可能
- 再帰的に処理
- 同様に動作する
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 全ノードを1回ずつ処理 |
| 空間計算量 | O(n) | キュー + 文字列 |

### 別解：DFS版（前順走査）

```go
type Codec struct{}

func Constructor() Codec {
    return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
    var res []string

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            res = append(res, "N")
            return
        }
        res = append(res, strconv.Itoa(node.Val))
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)
    return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
    vals := strings.Split(data, ",")
    index := 0

    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if vals[index] == "N" {
            index++
            return nil
        }
        val, _ := strconv.Atoi(vals[index])
        index++
        node := &TreeNode{Val: val}
        node.Left = dfs()
        node.Right = dfs()
        return node
    }

    return dfs()
}
```

### DFS版の動作

```
      1
     / \
    2   3
       / \
      4   5

前順走査（Root → Left → Right）:
1 → 2 → N → N → 3 → 4 → N → N → 5 → N → N

serialize結果: "1,2,N,N,3,4,N,N,5,N,N"

deserialize:
- "1" → ノード1を作成
  - 左に再帰 → "2" → ノード2を作成
    - 左に再帰 → "N" → null
    - 右に再帰 → "N" → null
  - 右に再帰 → "3" → ノード3を作成
    - 左に再帰 → "4" → ノード4を作成
      - 左に再帰 → "N" → null
      - 右に再帰 → "N" → null
    - 右に再帰 → "5" → ノード5を作成
      - ...
```

### 解法の比較

| 解法 | 特徴 | おすすめ度 |
|------|------|----------|
| BFS（この解法） | レベル順、直感的 | **面接推奨** |
| DFS（前順走査） | 再帰的、コンパクト | シンプルで良い |

### 重要なポイント

```
1. null の扱いが重要
   → 構造を保持するために明示的に記録

2. 区切り文字の選択
   → カンマ","は一般的
   → 値に含まれない文字を選ぶ

3. BFS vs DFS
   → どちらでも実装可能
   → BFSはレベルオーダーと同じパターン

4. 境界条件
   → 空の木 (root == nil)
   → 単一ノード
```

この問題は**Hard**だが、BFSの基本パターンを理解していれば解ける！
