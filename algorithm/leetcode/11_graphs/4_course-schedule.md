# Course Schedule (Medium)

## 問題内容

`prerequisites[i] = [a, b]` はコース `a` を取る前にコース `b` を取る必要があることを示す配列 `prerequisites` が与えられる。

`0` から `numCourses - 1` までラベル付けされた合計 `numCourses` 個のコースを取る必要がある。

すべてのコースを修了することが可能であれば `true` を、そうでなければ `false` を返す。

### 例

```
Input: numCourses = 2, prerequisites = [[0,1]]
Output: true
```
説明: まずコース1を取り（前提条件なし）、次にコース0を取る。

```
Input: numCourses = 2, prerequisites = [[0,1],[1,0]]
Output: false
```
説明: コース1を取るにはコース0が必要で、コース0を取るにはコース1が必要。不可能。

### 制約

- `1 <= numCourses <= 1000`
- `0 <= prerequisites.length <= 1000`
- `prerequisites[i].length == 2`
- `0 <= a[i], b[i] < numCourses`
- すべての前提条件ペアは一意

## ソースコード

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }
    for _, prereq := range prerequisites {
        crs, pre := prereq[0], prereq[1]
        preMap[crs] = append(preMap[crs], pre)
    }

    visiting := make(map[int]bool)

    var dfs func(crs int) bool
    dfs = func(crs int) bool {
        if visiting[crs] {
            return false
        }
        if len(preMap[crs]) == 0 {
            return true
        }

        visiting[crs] = true
        for _, pre := range preMap[crs] {
            if !dfs(pre) {
                return false
            }
        }
        visiting[crs] = false
        preMap[crs] = []int{}
        return true
    }

    for c := 0; c < numCourses; c++ {
        if !dfs(c) {
            return false
        }
    }
    return true
}
```

## アルゴリズムなど解説

### 基本戦略

各コースをノード、前提条件を有向エッジとして有向グラフを構築する。すべてのコースを修了できる ⟺ **グラフにサイクルがない**。

### 核心となる洞察

サイクルが存在すると：
- コースA → コースB → コースC → コースA
- どれも先に取ることができない → 不可能

サイクル検出には **DFS with 訪問中追跡** を使用：
- DFS中に「現在の再帰パス」にあるノードを追跡
- 同じパス上のノードを再訪問 → サイクル発見

### 動作の仕組み

1. **隣接リストの構築**
   ```go
   preMap := make(map[int][]int)
   for _, prereq := range prerequisites {
       crs, pre := prereq[0], prereq[1]
       preMap[crs] = append(preMap[crs], pre)
   }
   ```
   - 各コースの前提条件リストを作成

2. **訪問中セット**
   ```go
   visiting := make(map[int]bool)
   ```
   - **現在の再帰パス** にあるノードを追跡
   - 通常の「訪問済み」とは異なる

3. **DFS関数**
   ```go
   dfs = func(crs int) bool {
       if visiting[crs] {
           return false  // サイクル検出
       }
       if len(preMap[crs]) == 0 {
           return true   // 前提条件なし = OK
       }

       visiting[crs] = true
       for _, pre := range preMap[crs] {
           if !dfs(pre) {
               return false
           }
       }
       visiting[crs] = false
       preMap[crs] = []int{}  // 最適化: 処理済みマーク
       return true
   }
   ```

4. **最適化: preMap[crs] = []int{}**
   ```go
   preMap[crs] = []int{}
   ```
   - このコースのすべての前提条件が検証済み
   - 次回同じコースを訪問時、すぐに `true` を返せる
   - 同じパスを何度も辿ることを防ぐ

### 具体例

```
numCourses = 4
prerequisites = [[1,0], [2,1], [3,2]]

グラフ: 0 ← 1 ← 2 ← 3

dfs(0): visiting={}, preMap[0]=[] → true
dfs(1): visiting={1}, preMap[1]=[0]
  → dfs(0): true
  → preMap[1]=[] → true
dfs(2): visiting={2}, preMap[2]=[1]
  → dfs(1): preMap[1]=[] → true
  → preMap[2]=[] → true
dfs(3): visiting={3}, preMap[3]=[2]
  → dfs(2): preMap[2]=[] → true
  → preMap[3]=[] → true

結果: true
```

```
numCourses = 2
prerequisites = [[0,1], [1,0]]

グラフ: 0 ⟷ 1 (サイクル)

dfs(0): visiting={0}, preMap[0]=[1]
  → dfs(1): visiting={0,1}, preMap[1]=[0]
    → dfs(0): visiting[0]=true → サイクル検出!
    → return false

結果: false
```

### visiting vs visited の違い

| 状態 | 意味 | 用途 |
|------|------|------|
| visiting | 現在のDFSパス上にある | サイクル検出 |
| visited | 以前に処理完了 | 重複処理の回避 |

このコードでは `preMap[crs] = []int{}` が visited の役割を果たす。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(V + E) | 各ノードとエッジを1回処理 |
| 空間計算量 | O(V + E) | 隣接リスト + 再帰スタック |

> V = コース数、E = 前提条件数

### 別解：トポロジカルソート（カーン法）

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make([]int, numCourses)
    adj := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        adj[i] = []int{}
    }

    for _, prereq := range prerequisites {
        src, dst := prereq[0], prereq[1]
        indegree[dst]++
        adj[src] = append(adj[src], dst)
    }

    q := []int{}
    for n := 0; n < numCourses; n++ {
        if indegree[n] == 0 {
            q = append(q, n)
        }
    }

    finish := 0
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        finish++
        for _, nei := range adj[node] {
            indegree[nei]--
            if indegree[nei] == 0 {
                q = append(q, nei)
            }
        }
    }

    return finish == numCourses
}
```

**カーン法のアイデア**：
- 入次数（前提条件の数）が0のコースから始める
- そのコースを処理し、依存コースの入次数を減らす
- すべてのコースを処理できれば、サイクルなし

### グラフ問題のパターン

この問題は **サイクル検出 / トポロジカルソート** パターン：
- 有向グラフのサイクル検出
- DFSで「訪問中」と「訪問済み」を区別
- トポロジカル順序が存在するか確認
