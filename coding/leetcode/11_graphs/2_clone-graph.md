# Clone Graph (Medium)

## 問題内容

連結した無向グラフ内のノードへの参照が与えられる。グラフの **ディープコピー**（クローン）を返す。

グラフの各ノードは整数値と隣接ノードのリストを含む。

```go
type Node struct {
    Val       int
    Neighbors []*Node
}
```

ノードの値は `1` から `n` までの番号で、`n` はグラフ内のノードの総数。入力ノードは常に最初のノード（値 `1`）。

### 例

```
Input: adjList = [[2],[1,3],[2]]
Output: [[2],[1,3],[2]]
```
説明: 3つのノード。ノード1の隣接=[2]、ノード2の隣接=[1,3]、ノード3の隣接=[2]

```
Input: adjList = [[]]
Output: [[]]
```
説明: 隣接ノードのない1つのノード

```
Input: adjList = []
Output: []
```
説明: 空のグラフ

### 制約

- `0 <= ノード数 <= 100`
- `1 <= Node.val <= 100`
- 重複エッジや自己ループはない

## ソースコード

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    oldToNew := make(map[*Node]*Node)

    var dfs func(*Node) *Node
    dfs = func(node *Node) *Node {
        if node == nil {
            return nil
        }

        if _, found := oldToNew[node]; found {
            return oldToNew[node]
        }

        copy := &Node{Val: node.Val}
        oldToNew[node] = copy
        for _, nei := range node.Neighbors {
            copy.Neighbors = append(copy.Neighbors, dfs(nei))
        }
        return copy
    }

    return dfs(node)
}
```

## アルゴリズムなど解説

### 基本戦略

DFSでグラフを走査しながら、各ノードのクローンを作成する。サイクルがある場合に無限ループを防ぐため、オリジナル→クローンのマッピングを保持する。

### 核心となる洞察

グラフには **サイクル** が含まれる可能性があるため、単純に再帰的にコピーするだけでは無限ループに陥る。

解決策：**ハッシュマップ（old → new）** を使用
- ノードを初めて見たとき、そのコピーを作成
- 同じノードを再度見たとき、既に作成したコピーを再利用
- これにより無限ループを回避し、各ノードを正確に1回だけクローン

### 動作の仕組み

1. **マッピング用のハッシュマップ**
   ```go
   oldToNew := make(map[*Node]*Node)
   ```
   - オリジナルノード → クローンノードの対応を保存

2. **DFS関数**
   ```go
   var dfs func(*Node) *Node
   dfs = func(node *Node) *Node {
       if node == nil {
           return nil
       }
       if _, found := oldToNew[node]; found {
           return oldToNew[node]
       }
       // ...
   }
   ```
   - nullチェック
   - 既にクローン済みなら、そのクローンを返す

3. **ノードのクローンと隣接ノードの処理**
   ```go
   copy := &Node{Val: node.Val}
   oldToNew[node] = copy
   for _, nei := range node.Neighbors {
       copy.Neighbors = append(copy.Neighbors, dfs(nei))
   }
   return copy
   ```
   - 新しいノードを作成
   - マップに登録（**重要**: 隣接ノードを処理する前に登録）
   - 隣接ノードを再帰的にクローンして追加

### 具体例

```
adjList = [[2],[1,3],[2]]

Node1 → neighbors: [Node2]
Node2 → neighbors: [Node1, Node3]
Node3 → neighbors: [Node2]

dfs(Node1):
  - copy1 = {Val: 1}, oldToNew[Node1] = copy1
  - dfs(Node2):
    - copy2 = {Val: 2}, oldToNew[Node2] = copy2
    - dfs(Node1): 既に存在 → copy1を返す
    - dfs(Node3):
      - copy3 = {Val: 3}, oldToNew[Node3] = copy3
      - dfs(Node2): 既に存在 → copy2を返す
      - copy3.Neighbors = [copy2]
    - copy2.Neighbors = [copy1, copy3]
  - copy1.Neighbors = [copy2]

結果: 完全にクローンされたグラフ
```

### なぜ先にマップに登録するか

```go
copy := &Node{Val: node.Val}
oldToNew[node] = copy  // ← 隣接ノード処理前に登録
for _, nei := range node.Neighbors {
    copy.Neighbors = append(copy.Neighbors, dfs(nei))
}
```

- サイクルがある場合、隣接ノードが現在のノードを参照している可能性がある
- 先に登録しておくことで、再帰呼び出しで同じノードに戻ってきたときに既存のクローンを返せる
- 無限ループを防止

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(V + E) | 各ノードと各エッジを1回処理 |
| 空間計算量 | O(V) | ハッシュマップ + 再帰スタック |

> V = ノード数、E = エッジ数

### 別解：BFS版

```go
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    oldToNew := make(map[*Node]*Node)
    oldToNew[node] = &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    queue := make([]*Node, 0)
    queue = append(queue, node)

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        for _, nei := range cur.Neighbors {
            if _, exists := oldToNew[nei]; !exists {
                oldToNew[nei] = &Node{Val: nei.Val, Neighbors: make([]*Node, 0)}
                queue = append(queue, nei)
            }
            oldToNew[cur].Neighbors = append(oldToNew[cur].Neighbors, oldToNew[nei])
        }
    }

    return oldToNew[node]
}
```

### グラフ問題のパターン

この問題は **グラフ走査 + クローン** パターン：
- ハッシュマップで訪問済み管理とマッピングを同時に行う
- サイクル対策が重要
- DFS/BFSどちらでも解ける
