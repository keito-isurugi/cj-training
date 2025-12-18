# Word Search (Medium)

## 問題内容

2次元の文字グリッド `board` と文字列 `word` が与えられる。`word` がグリッド内に存在すれば `true`、そうでなければ `false` を返す。

単語が存在するとは、水平または垂直に隣接するセルでパスを形成できること。同じセルは1つの単語内で複数回使用できない。

### 例

```
Input:
board = [
  ["A","B","C","D"],
  ["S","A","A","T"],
  ["A","C","A","E"]
],
word = "CAT"

Output: true

パス: C(0,2) → A(1,2) → T(1,3)
```

```
Input:
board = [
  ["A","B","C","D"],
  ["S","A","A","T"],
  ["A","C","A","E"]
],
word = "BAT"

Output: false
（Bの隣にAはあるが、その隣にTがない）
```

### 制約

- `1 <= board.length, board[i].length <= 5`
- `1 <= word.length <= 10`
- `board` と `word` は英大文字・小文字のみ

## ソースコード

```go
func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])

    var dfs func(r, c, i int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || c < 0 || r >= rows || c >= cols ||
           board[r][c] != word[i] || board[r][c] == '#' {
            return false
        }

        temp := board[r][c]
        board[r][c] = '#'
        res := dfs(r+1, c, i+1) ||
               dfs(r-1, c, i+1) ||
               dfs(r, c+1, i+1) ||
               dfs(r, c-1, i+1)
        board[r][c] = temp

        return res
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if dfs(r, c, 0) {
                return true
            }
        }
    }
    return false
}
```

## アルゴリズムなど解説

### 基本戦略

全てのセルを開始点として試し、DFS + バックトラッキングで4方向に探索。訪問済みセルを一時的にマークして再訪問を防ぐ。

### 探索のイメージ

```
board:              word = "CAT"
A B C D
S A A T
A C A E

C(0,2)から開始:
    C ← 一致！
    ↓
    A ← 一致！
    ↓
    ? → 上下左右を探索

C(0,2) → A(1,2) → T(1,3) ✓
```

### DFSの終了条件

```go
// 成功: 単語を全て見つけた
if i == len(word) {
    return true
}

// 失敗: 範囲外、不一致、または訪問済み
if r < 0 || c < 0 || r >= rows || c >= cols ||
   board[r][c] != word[i] || board[r][c] == '#' {
    return false
}
```

### 訪問済みマーキング

```go
temp := board[r][c]    // 元の値を保存
board[r][c] = '#'      // 訪問済みマーク

res := dfs(...)        // 探索

board[r][c] = temp     // 元に戻す（バックトラック）
```

```
なぜ '#' を使うか？

- 英字以外の文字
- board[r][c] != word[i] の条件で自動的に false になる
- 別途 visited 配列が不要 → 空間効率が良い
```

### 視覚的な理解

```
board:          word = "CAT"
A B C D
S A A T
A C A E

━━━ 探索開始: (0,0)='A' ━━━
A != 'C' → false

━━━ 探索開始: (0,2)='C' ━━━
Step 1: (0,2)='C' == word[0]='C' ✓
        board[0][2] = '#'

        A B # D
        S A A T
        A C A E

Step 2: 4方向を探索
        上: (-1,2) → 範囲外
        下: (1,2)='A' == word[1]='A' ✓
        左: (0,1)='B' != 'A'
        右: (0,3)='D' != 'A'

        → 下に進む

Step 3: (1,2)='A' == word[1]='A' ✓
        board[1][2] = '#'

        A B # D
        S A # T
        A C A E

Step 4: 4方向を探索
        上: (0,2)='#' → 訪問済み
        下: (2,2)='A' != 'T'
        左: (1,1)='A' != 'T'
        右: (1,3)='T' == word[2]='T' ✓

        → 右に進む

Step 5: (1,3)='T' == word[2]='T' ✓
        i+1 == len(word) → return true!

結果: true
```

### なぜバックトラックが必要か

```
board:          word = "AAB"
A A
A B

パス1を試す: (0,0) → (0,1) → ?
A(0,0) → A(0,1) → B がない → 失敗

もしバックトラックしないと:
# #
A B
↑ (0,0)と(0,1)が '#' のまま

パス2が試せない: (0,0) → (1,0) → (1,1)='B'

バックトラックすると:
A A  ← 元に戻る
A B

パス2: A(0,0) → A(1,0) → B(1,1) ✓
```

### 4方向探索

```go
res := dfs(r+1, c, i+1) ||  // 下
       dfs(r-1, c, i+1) ||  // 上
       dfs(r, c+1, i+1) ||  // 右
       dfs(r, c-1, i+1)     // 左
```

```
短絡評価（||）:
- 1つでも true が見つかれば残りは評価しない
- 効率的！
```

### 全セルから開始を試す

```go
for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
        if dfs(r, c, 0) {
            return true
        }
    }
}
return false
```

```
単語の最初の文字がどこにあるか分からない
→ 全セルを開始点として試す

最適化: 最初の文字と一致するセルだけ試す
if board[r][c] == word[0] && dfs(r, c, 0) {
    return true
}
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n × 4^L) | m×n=グリッドサイズ, L=単語長 |
| 空間計算量 | O(L) | 再帰スタックの深さ |

### 別解：visited配列を使用

```go
func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }

    var dfs func(r, c, i int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || c < 0 || r >= rows || c >= cols ||
           visited[r][c] || board[r][c] != word[i] {
            return false
        }

        visited[r][c] = true
        res := dfs(r+1, c, i+1) ||
               dfs(r-1, c, i+1) ||
               dfs(r, c+1, i+1) ||
               dfs(r, c-1, i+1)
        visited[r][c] = false

        return res
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if dfs(r, c, 0) {
                return true
            }
        }
    }
    return false
}
```

### 2つのアプローチの比較

| 方法 | 空間 | 特徴 |
|------|------|------|
| board書き換え | O(L) | 空間効率◎、元データを変更 |
| visited配列 | O(m×n) | 元データを保持、分かりやすい |

### バックトラッキングのパターン

```go
func backtrack(位置, 状態) bool {
    // 1. 成功条件
    if 完了 {
        return true
    }

    // 2. 失敗条件
    if 無効 {
        return false
    }

    // 3. 選択（マーク）
    visited[位置] = true

    // 4. 再帰で探索
    for 隣接 in 4方向 {
        if backtrack(隣接, 次の状態) {
            return true
        }
    }

    // 5. 選択を取り消す（バックトラック）
    visited[位置] = false

    return false
}
```

### 関連問題

| 問題 | 共通点 |
|------|--------|
| Word Search II | 複数単語を同時に探索（Trie使用） |
| Number of Islands | グリッドDFS |
| Surrounded Regions | グリッドDFS + マーキング |

グリッド探索 + バックトラッキングの典型問題！
