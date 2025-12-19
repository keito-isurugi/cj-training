# Graph Valid Tree (Medium)

## 問題内容

`0` から `n - 1` までラベル付けされた `n` 個のノードと、**無向** エッジのリストが与えられる。これらのエッジが有効な木を構成するかどうかをチェックする関数を書く。

### 例

```
Input: n = 5, edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
Output: true
```

```
Input: n = 5, edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
Output: false
```

### 制約

- `1 <= n <= 100`
- `0 <= edges.length <= n * (n - 1) / 2`
- 重複エッジはない。`[0, 1]` と `[1, 0]` は同じエッジ

## ソースコード

```go
func validTree(n int, edges [][]int) bool {
    if len(edges) > n-1 {
        return false
    }

    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    visit := make(map[int]bool)
    var dfs func(node, parent int) bool
    dfs = func(node, parent int) bool {
        if visit[node] {
            return false
        }
        visit[node] = true
        for _, nei := range adj[node] {
            if nei == parent {
                continue
            }
            if !dfs(nei, node) {
                return false
            }
        }
        return true
    }

    return dfs(0, -1) && len(visit) == n
}
```

## アルゴリズムなど解説

### 基本戦略

グラフが有効な **木** であるための条件：
1. **サイクルがない**
2. **完全に連結されている**（すべてのノードが到達可能）

追加の洞察：`n` ノードの木は正確に `n - 1` 本のエッジを持つ。

### 木の性質

| 性質 | 説明 |
|------|------|
| エッジ数 | 正確に n - 1 本 |
| サイクル | なし |
| 連結性 | 完全に連結 |
| 任意の2ノード間のパス | 正確に1つ |

### 動作の仕組み

1. **エッジ数チェック**
   ```go
   if len(edges) > n-1 {
       return false
   }
   ```
   - エッジが n-1 より多い → 必ずサイクルがある
   - 早期リターンで効率化

2. **隣接リストの構築**
   ```go
   adj := make([][]int, n)
   for _, edge := range edges {
       u, v := edge[0], edge[1]
       adj[u] = append(adj[u], v)
       adj[v] = append(adj[v], u)
   }
   ```

3. **DFS関数（親を追跡）**
   ```go
   dfs = func(node, parent int) bool {
       if visit[node] {
           return false  // サイクル検出
       }
       visit[node] = true
       for _, nei := range adj[node] {
           if nei == parent {
               continue  // 親へは戻らない
           }
           if !dfs(nei, node) {
               return false
           }
       }
       return true
   }
   ```

4. **最終チェック**
   ```go
   return dfs(0, -1) && len(visit) == n
   ```
   - DFS成功（サイクルなし）かつ
   - すべてのノードを訪問（連結している）

### なぜ親を追跡するのか

無向グラフでは、エッジ `[u, v]` は両方向に走査可能：
```
u → v → u (戻る)
```

親を追跡しないと、この「戻り」をサイクルと誤検出してしまう。

```go
if nei == parent {
    continue  // 親への戻りはスキップ
}
```

### 具体例

```
n = 5, edges = [[0, 1], [0, 2], [0, 3], [1, 4]]

    0
   /|\
  1 2 3
  |
  4

dfs(0, -1): visit={0}
  → dfs(1, 0): visit={0,1}
    → dfs(4, 1): visit={0,1,4}
      → 隣接=[1], parent=1 → スキップ
      → return true
    → return true
  → dfs(2, 0): visit={0,1,2,4}
    → return true
  → dfs(3, 0): visit={0,1,2,3,4}
    → return true
  → return true

len(visit)=5 == n=5 → true
結果: true
```

```
n = 5, edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]

サイクル: 1 → 2 → 3 → 1

dfs(0, -1): visit={0}
  → dfs(1, 0): visit={0,1}
    → dfs(2, 1): visit={0,1,2}
      → dfs(3, 2): visit={0,1,2,3}
        → 隣接=[2,1]
        → 2 == parent → スキップ
        → dfs(1, 3): visit[1]=true → サイクル検出!
        → return false

結果: false
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(V + E) | 各ノードとエッジを1回処理 |
| 空間計算量 | O(V + E) | 隣接リスト + 訪問セット |

> V = ノード数、E = エッジ数

### 別解：Union-Find

```go
type DSU struct {
    Parent []int
    Size   []int
    Comps  int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n+1)
    size := make([]int, n+1)
    for i := 0; i <= n; i++ {
        parent[i] = i
        size[i] = 1
    }
    return &DSU{Parent: parent, Size: size, Comps: n}
}

func (dsu *DSU) Find(node int) int {
    if dsu.Parent[node] != node {
        dsu.Parent[node] = dsu.Find(dsu.Parent[node])
    }
    return dsu.Parent[node]
}

func (dsu *DSU) Union(u, v int) bool {
    pu, pv := dsu.Find(u), dsu.Find(v)
    if pu == pv {
        return false  // 同じ成分 → サイクル発生
    }
    dsu.Comps--
    if dsu.Size[pu] < dsu.Size[pv] {
        pu, pv = pv, pu
    }
    dsu.Size[pu] += dsu.Size[pv]
    dsu.Parent[pv] = pu
    return true
}

func validTree(n int, edges [][]int) bool {
    if len(edges) > n-1 {
        return false
    }
    dsu := NewDSU(n)
    for _, edge := range edges {
        if !dsu.Union(edge[0], edge[1]) {
            return false  // サイクル検出
        }
    }
    return dsu.Comps == 1  // 連結性チェック
}
```

**Union-Findのアイデア**：
- 各エッジを処理：既に同じ成分 → サイクル
- 最終的に1つの成分 → 連結

### グラフ問題のパターン

この問題は **木の検証** パターン：
- サイクル検出 + 連結性チェック
- DFS（親追跡）または Union-Find で解ける
- エッジ数チェックで早期枝刈り可能

### 関連問題との比較

| 問題 | 確認すること |
|------|-------------|
| Number of Islands | 連結成分の数 |
| Number of Connected Components | 連結成分の数 |
| Course Schedule | サイクルの有無（有向） |
| **Graph Valid Tree** | サイクルなし + 1つの連結成分 |
