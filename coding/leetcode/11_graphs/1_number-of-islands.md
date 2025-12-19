# Number of Islands (Medium)

## 問題内容

`'1'`（陸地）と `'0'`（水）で構成される `m x n` の2次元グリッドが与えられる。島の数を返す。

**島** は水に囲まれており、隣接する陸地を水平方向または垂直方向に接続することで形成される。グリッドの4辺はすべて水に囲まれていると仮定できる。

### 例

```
Input: grid = [
    ["0","1","1","1","0"],
    ["0","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
]
Output: 1
```

```
Input: grid = [
    ["1","1","0","0","1"],
    ["1","1","0","0","1"],
    ["0","0","1","0","0"],
    ["0","0","0","1","1"]
]
Output: 4
```

### 制約

- `1 <= m, n <= 100`
- `grid[i][j]` は `'0'` または `'1'`

## ソースコード

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    rows, cols := len(grid), len(grid[0])
    islands := 0

    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
            return
        }
        grid[r][c] = '0'  // mark as visited
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                islands++
                dfs(r, c)
            }
        }
    }

    return islands
}
```

## アルゴリズムなど解説

### 基本戦略

グリッドを走査し、未訪問の陸地 `'1'` を見つけるたびに島のカウントを増やし、DFS/BFSでその島全体を訪問済みにする。

### DFS（深さ優先探索）アプローチ

1. グリッドの各セルを走査
2. `'1'` を見つけたら：
   - 島のカウントをインクリメント
   - DFSでその島に属するすべての `'1'` を `'0'` に変更（訪問済みマーク）
3. すべてのセルを処理したら、島の数を返す

### 動作の仕組み

1. **初期化**
   ```go
   rows, cols := len(grid), len(grid[0])
   islands := 0
   ```

2. **DFS関数**
   ```go
   var dfs func(r, c int)
   dfs = func(r, c int) {
       if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
           return
       }
       grid[r][c] = '0'  // mark as visited
       dfs(r+1, c)
       dfs(r-1, c)
       dfs(r, c+1)
       dfs(r, c-1)
   }
   ```
   - 境界チェックと水 `'0'` のチェック
   - 現在のセルを `'0'` に変更して訪問済みにする
   - 4方向に再帰的に探索

3. **メインループ**
   ```go
   for r := 0; r < rows; r++ {
       for c := 0; c < cols; c++ {
           if grid[r][c] == '1' {
               islands++
               dfs(r, c)
           }
       }
   }
   ```
   - 新しい島を見つけるたびにカウントアップ
   - DFSでその島全体を訪問済みにする

### 具体例

```
grid = [
    ["1","1","0","0","1"],
    ["1","1","0","0","1"],
    ["0","0","1","0","0"],
    ["0","0","0","1","1"]
]

(0,0): '1' → islands=1, DFSで(0,0),(0,1),(1,0),(1,1)を訪問済みに
(0,4): '1' → islands=2, DFSで(0,4),(1,4)を訪問済みに
(2,2): '1' → islands=3, DFSで(2,2)を訪問済みに
(3,3): '1' → islands=4, DFSで(3,3),(3,4)を訪問済みに

最終結果: 4
```

### なぜこのアルゴリズムが正しいか

- 各島は連結した `'1'` の集合
- DFSで連結した `'1'` をすべて訪問済みにすることで、同じ島を複数回カウントしない
- 新しい `'1'` を見つけるたびに、それは新しい島の始点

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n) | 各セルを最大2回訪問（メインループ + DFS） |
| 空間計算量 | O(m × n) | 最悪の場合、再帰スタックの深さ |

### 別解：BFS版

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    rows, cols := len(grid), len(grid[0])
    islands := 0
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    bfs := func(r, c int) {
        queue := [][]int{{r, c}}
        grid[r][c] = '0'

        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]

            for _, d := range directions {
                nr, nc := curr[0]+d[0], curr[1]+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
                    grid[nr][nc] = '0'
                    queue = append(queue, []int{nr, nc})
                }
            }
        }
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                islands++
                bfs(r, c)
            }
        }
    }

    return islands
}
```

### グラフ問題のパターン

この問題は **Grid DFS/BFS** パターンの代表的な問題：
- グリッドをグラフとして扱う
- 各セルはノード、隣接セルはエッジ
- 連結成分（島）の数を数える
