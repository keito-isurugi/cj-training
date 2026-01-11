# Number of Connected Components in an Undirected Graph (Medium)

## 問題内容

`n` 個のノードを持つ無向グラフがある。`edges[i] = [a, b]` はノード `a` とノード `b` の間にエッジがあることを意味する `edges` 配列も与えられる。

ノードは `0` から `n - 1` まで番号付けされている。

グラフ内の **連結成分の総数** を返す。

### 例

```
Input: n = 3, edges = [[0,1], [0,2]]
Output: 1
```

```
Input: n = 6, edges = [[0,1], [1,2], [2,3], [4,5]]
Output: 2
```

### 制約

- `1 <= n <= 100`
- `0 <= edges.length <= n * (n - 1) / 2`

## ソースコード

```go
func countComponents(n int, edges [][]int) int {
    adj := make([][]int, n)
    visit := make([]bool, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    var dfs func(int)
    dfs = func(node int) {
        for _, nei := range adj[node] {
            if !visit[nei] {
                visit[nei] = true
                dfs(nei)
            }
        }
    }

    res := 0
    for node := 0; node < n; node++ {
        if !visit[node] {
            visit[node] = true
            dfs(node)
            res++
        }
    }
    return res
}
```

## アルゴリズムなど解説

### 基本戦略

**連結成分** とは、グループ内のすべてのノードが互いに到達可能なノードの集合。

DFSを使用：
- 未訪問のノードからDFSを開始すると、その連結成分内の **すべてのノード** を訪問する
- 新しい未訪問ノードからDFSを開始するたびに、**1つの新しい連結成分** を発見

### 動作の仕組み

1. **隣接リストの構築**
   ```go
   adj := make([][]int, n)
   for _, edge := range edges {
       u, v := edge[0], edge[1]
       adj[u] = append(adj[u], v)
       adj[v] = append(adj[v], u)
   }
   ```
   - 無向グラフなので両方向にエッジを追加

2. **訪問済み配列**
   ```go
   visit := make([]bool, n)
   ```
   - どのノードを訪問したかを追跡

3. **DFS関数**
   ```go
   var dfs func(int)
   dfs = func(node int) {
       for _, nei := range adj[node] {
           if !visit[nei] {
               visit[nei] = true
               dfs(nei)
           }
       }
   }
   ```
   - 隣接する未訪問ノードをすべて訪問

4. **メインループ**
   ```go
   res := 0
   for node := 0; node < n; node++ {
       if !visit[node] {
           visit[node] = true
           dfs(node)
           res++
       }
   }
   ```
   - 未訪問ノードを見つけるたびに連結成分カウントを増加
   - DFSでその成分全体を訪問済みにする

### 具体例

```
n = 6, edges = [[0,1], [1,2], [2,3], [4,5]]

隣接リスト:
0: [1]
1: [0, 2]
2: [1, 3]
3: [2]
4: [5]
5: [4]

処理:
node=0: 未訪問 → res=1, DFSで0,1,2,3を訪問
node=1: 訪問済み → スキップ
node=2: 訪問済み → スキップ
node=3: 訪問済み → スキップ
node=4: 未訪問 → res=2, DFSで4,5を訪問
node=5: 訪問済み → スキップ

結果: 2つの連結成分
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(V + E) | 各ノードとエッジを1回処理 |
| 空間計算量 | O(V + E) | 隣接リスト + 訪問配列 |

> V = ノード数、E = エッジ数

### 別解：Union-Find（素集合データ構造）

```go
type DSU struct {
    parent []int
    rank   []int
}

func NewDSU(n int) *DSU {
    dsu := &DSU{
        parent: make([]int, n),
        rank:   make([]int, n),
    }
    for i := 0; i < n; i++ {
        dsu.parent[i] = i
        dsu.rank[i] = 1
    }
    return dsu
}

func (dsu *DSU) Find(node int) int {
    cur := node
    for cur != dsu.parent[cur] {
        dsu.parent[cur] = dsu.parent[dsu.parent[cur]]
        cur = dsu.parent[cur]
    }
    return cur
}

func (dsu *DSU) Union(u, v int) bool {
    pu := dsu.Find(u)
    pv := dsu.Find(v)
    if pu == pv {
        return false
    }
    if dsu.rank[pv] > dsu.rank[pu] {
        pu, pv = pv, pu
    }
    dsu.parent[pv] = pu
    dsu.rank[pu] += dsu.rank[pv]
    return true
}

func countComponents(n int, edges [][]int) int {
    dsu := NewDSU(n)
    res := n
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        if dsu.Union(u, v) {
            res--
        }
    }
    return res
}
```

**Union-Findのアイデア**：
- 最初は各ノードが独自の成分（成分数 = n）
- エッジを処理するたびに2つのノードを結合
- 異なる成分のノードを結合すると、成分数が1減る

### DFS vs Union-Find

| 観点 | DFS/BFS | Union-Find |
|------|---------|------------|
| 時間計算量 | O(V + E) | O(E × α(V)) ≈ O(E) |
| 空間計算量 | O(V + E) | O(V) |
| 実装の簡単さ | より直感的 | やや複雑 |
| 用途 | 静的グラフ | 動的エッジ追加に強い |

### グラフ問題のパターン

この問題は **連結成分カウント** パターン：
- グラフの構造を把握する基本問題
- DFS/BFS または Union-Find で解ける
- Number of Islands と同じ考え方（グリッド vs グラフの違い）
