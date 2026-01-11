# Pacific Atlantic Water Flow (Medium)

## 問題内容

`heights[r][c]` が座標 `(r, c)` のセルの **海抜高度** を表す長方形の島 `heights` が与えられる。

島は **上辺と左辺** で太平洋に、**下辺と右辺** で大西洋に接している。

水は4方向（上下左右）に、**高さが同じか低い** 隣接セルに流れることができる。また、海に隣接するセルから海に流れ出ることもできる。

**太平洋と大西洋の両方** に水が流れることができるすべてのセルを見つけ、`[r, c]` のリストとして返す。

### 例

```
Input: heights = [
  [4,2,7,3,4],
  [7,4,6,4,7],
  [6,3,5,3,6]
]

Output: [[0,2],[0,4],[1,0],[1,1],[1,2],[1,3],[1,4],[2,0]]
```

```
Input: heights = [[1],[1]]

Output: [[0,0],[0,1]]
```

### 制約

- `1 <= heights.length, heights[r].length <= 100`
- `0 <= heights[r][c] <= 1000`

## ソースコード

```go
func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
    pac := make(map[[2]int]bool)
    atl := make(map[[2]int]bool)

    var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)
    dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
        coord := [2]int{r, c}
        if visit[coord] ||
           r < 0 || c < 0 ||
           r == rows || c == cols ||
           heights[r][c] < prevHeight {
            return
        }

        visit[coord] = true

        dfs(r+1, c, visit, heights[r][c])
        dfs(r-1, c, visit, heights[r][c])
        dfs(r, c+1, visit, heights[r][c])
        dfs(r, c-1, visit, heights[r][c])
    }

    for c := 0; c < cols; c++ {
        dfs(0, c, pac, heights[0][c])
        dfs(rows-1, c, atl, heights[rows-1][c])
    }

    for r := 0; r < rows; r++ {
        dfs(r, 0, pac, heights[r][0])
        dfs(r, cols-1, atl, heights[r][cols-1])
    }

    result := make([][]int, 0)
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            coord := [2]int{r, c}
            if pac[coord] && atl[coord] {
                result = append(result, []int{r, c})
            }
        }
    }

    return result
}
```

## アルゴリズムなど解説

### 基本戦略

各セルからDFSを始めるのではなく、**逆方向に考える**：海岸線から内陸に向かって「水が逆流できるセル」を探す。

### 核心となる洞察

**順方向の考え方**（非効率）：
- 各セルから両方の海に到達できるか確認 → O((m × n)²)

**逆方向の考え方**（効率的）：
- 太平洋の海岸線から「登れる」セルを全て記録 → `pac` セット
- 大西洋の海岸線から「登れる」セルを全て記録 → `atl` セット
- 両方のセットに含まれるセル = 両方の海に流れることができる

なぜ逆方向が有効か：
- 水は低い方へ流れる
- 逆に見ると：海から「登って」到達できるセル = そのセルから海へ流れ落ちることができる

### 動作の仕組み

1. **2つの訪問セットを用意**
   ```go
   pac := make(map[[2]int]bool)  // 太平洋に到達可能
   atl := make(map[[2]int]bool)  // 大西洋に到達可能
   ```

2. **DFS関数（逆流ルール）**
   ```go
   dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
       // ...
       if heights[r][c] < prevHeight {
           return  // 登れない（逆流できない）
       }
       // ...
   }
   ```
   - `heights[r][c] >= prevHeight` の場合のみ進める
   - これは実際の水の流れとは逆方向

3. **太平洋側からDFS**
   ```go
   for c := 0; c < cols; c++ {
       dfs(0, c, pac, heights[0][c])      // 上辺
   }
   for r := 0; r < rows; r++ {
       dfs(r, 0, pac, heights[r][0])      // 左辺
   }
   ```

4. **大西洋側からDFS**
   ```go
   for c := 0; c < cols; c++ {
       dfs(rows-1, c, atl, heights[rows-1][c])  // 下辺
   }
   for r := 0; r < rows; r++ {
       dfs(r, cols-1, atl, heights[r][cols-1])  // 右辺
   }
   ```

5. **両方に含まれるセルを抽出**
   ```go
   if pac[coord] && atl[coord] {
       result = append(result, []int{r, c})
   }
   ```

### 具体例

```
heights = [
  [4,2,7,3,4],
  [7,4,6,4,7],
  [6,3,5,3,6]
]

太平洋（上・左）から到達可能:
  P P P P P
  P P P P P
  P . . . .

大西洋（下・右）から到達可能:
  . . A A A
  A A A A A
  A A A A A

両方に含まれる:
  . . X . X
  X X X X X
  X . . . .

結果: (0,2), (0,4), (1,0), (1,1), (1,2), (1,3), (1,4), (2,0)
```

### なぜ逆流アプローチが正しいか

- セルAから海Bに水が流れる条件：Aから海岸まで「高さが減少または同じ」の経路がある
- 逆に見ると：海岸Bからセルまで「高さが増加または同じ」の経路がある
- 海岸から「登って」到達可能 = そのセルから海へ「下って」到達可能

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n) | 各セルを最大2回訪問 |
| 空間計算量 | O(m × n) | 2つの訪問セット |

### 別解：BFS版

```go
func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    bfs := func(sources [][2]int) map[[2]int]bool {
        visited := make(map[[2]int]bool)
        queue := sources
        for _, s := range sources {
            visited[s] = true
        }

        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]
            r, c := curr[0], curr[1]

            for _, d := range directions {
                nr, nc := r+d[0], c+d[1]
                coord := [2]int{nr, nc}
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                   !visited[coord] && heights[nr][nc] >= heights[r][c] {
                    visited[coord] = true
                    queue = append(queue, coord)
                }
            }
        }
        return visited
    }

    // Build source lists for each ocean
    var pacSources, atlSources [][2]int
    for c := 0; c < cols; c++ {
        pacSources = append(pacSources, [2]int{0, c})
        atlSources = append(atlSources, [2]int{rows-1, c})
    }
    for r := 0; r < rows; r++ {
        pacSources = append(pacSources, [2]int{r, 0})
        atlSources = append(atlSources, [2]int{r, cols-1})
    }

    pac := bfs(pacSources)
    atl := bfs(atlSources)

    result := make([][]int, 0)
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            coord := [2]int{r, c}
            if pac[coord] && atl[coord] {
                result = append(result, []int{r, c})
            }
        }
    }
    return result
}
```

### グラフ問題のパターン

この問題は **Multi-Source BFS/DFS** パターン：
- 複数の始点から同時に探索
- 逆方向の発想で効率化
- 2つの条件の交差（AND条件）を求める
