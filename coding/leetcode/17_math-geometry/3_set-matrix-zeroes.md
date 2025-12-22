# Set Matrix Zeroes (Medium)

## 問題内容

整数からなる `m x n` の行列 `matrix` が与えられる。要素が `0` であれば、その行と列全体を `0` にする。

行列を **インプレース** で更新しなければならない。

**フォローアップ**: `O(1)` の空間計算量で解けるか？

### 例

```
Input: matrix = [
  [0,1],
  [1,0]
]

Output: [
  [0,0],
  [0,0]
]
```

```
Input: matrix = [
  [1,2,3],
  [4,0,5],
  [6,7,8]
]

Output: [
  [1,0,3],
  [0,0,0],
  [6,0,8]
]
```

### 制約

- `1 <= matrix.length, matrix[0].length <= 100`
- `-2^31 <= matrix[i][j] <= (2^31) - 1`

## ソースコード

```go
func setZeroes(matrix [][]int) {
    ROWS, COLS := len(matrix), len(matrix[0])
    rowZero := false

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if matrix[r][c] == 0 {
                matrix[0][c] = 0
                if r > 0 {
                    matrix[r][0] = 0
                } else {
                    rowZero = true
                }
            }
        }
    }

    for r := 1; r < ROWS; r++ {
        for c := 1; c < COLS; c++ {
            if matrix[0][c] == 0 || matrix[r][0] == 0 {
                matrix[r][c] = 0
            }
        }
    }

    if matrix[0][0] == 0 {
        for r := 0; r < ROWS; r++ {
            matrix[r][0] = 0
        }
    }

    if rowZero {
        for c := 0; c < COLS; c++ {
            matrix[0][c] = 0
        }
    }
}
```

## アルゴリズムなど解説

### 基本戦略（O(1) 空間）

**問題点**: 0を見つけて即座に行・列を0にすると、新たに作った0に反応してしまう

**解決策**: 行列の **最初の行と最初の列** をマーカーとして使用

- 最初の行 → どの列を0にするか記録
- 最初の列 → どの行を0にするか記録
- `rowZero` フラグ → 最初の行自体を0にするか記録

### 動作の仕組み

1. **マーカーの設定（1回目のスキャン）**
   ```go
   for r := 0; r < ROWS; r++ {
       for c := 0; c < COLS; c++ {
           if matrix[r][c] == 0 {
               matrix[0][c] = 0  // 列のマーカー
               if r > 0 {
                   matrix[r][0] = 0  // 行のマーカー
               } else {
                   rowZero = true  // 最初の行用フラグ
               }
           }
       }
   }
   ```

2. **内部要素の更新（2回目のスキャン）**
   ```go
   for r := 1; r < ROWS; r++ {
       for c := 1; c < COLS; c++ {
           if matrix[0][c] == 0 || matrix[r][0] == 0 {
               matrix[r][c] = 0
           }
       }
   }
   ```

3. **最初の列の更新**
   ```go
   if matrix[0][0] == 0 {
       for r := 0; r < ROWS; r++ {
           matrix[r][0] = 0
       }
   }
   ```

4. **最初の行の更新**
   ```go
   if rowZero {
       for c := 0; c < COLS; c++ {
           matrix[0][c] = 0
       }
   }
   ```

### 具体例

```
元の行列:
1 2 3
4 0 5
6 7 8

1回目スキャン後（マーカー設定）:
1 0 3    ← matrix[0][1] = 0 (列1のマーカー)
0 0 5    ← matrix[1][0] = 0 (行1のマーカー)
6 7 8

2回目スキャン後（内部更新）:
1 0 3
0 0 0    ← 行1全体が0に
6 0 8    ← 列1が0に

最終結果:
1 0 3
0 0 0
6 0 8
```

### なぜ rowZero フラグが必要か

`matrix[0][0]` は最初の行と最初の列の両方の交点。
- 最初の行に0があるか
- 最初の列に0があるか

この2つの情報を1つのセルで管理できないため、`rowZero` フラグで最初の行を別管理。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m×n) | 行列を2回スキャン |
| 空間計算量 | O(1) | 追加の配列なし |

### 別解1：O(m+n) 空間

```go
func setZeroes(matrix [][]int) {
    ROWS, COLS := len(matrix), len(matrix[0])
    rows := make([]bool, ROWS)
    cols := make([]bool, COLS)

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if matrix[r][c] == 0 {
                rows[r] = true
                cols[c] = true
            }
        }
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if rows[r] || cols[c] {
                matrix[r][c] = 0
            }
        }
    }
}
```

この方法はシンプルだが、追加の配列が必要。

### 3つの解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|------------|------------|------|
| O(1) 空間 | O(m×n) | O(1) | 最適だが複雑 |
| 配列使用 | O(m×n) | O(m+n) | シンプル |
| コピー使用 | O(m×n) | O(m×n) | 最もシンプル |

### Matrix パターン

この問題は **インプレース更新** パターン：
- 行列自体をマーカーとして使用
- 2パス処理（マーク→更新）
- 境界ケースの注意深い処理

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Rotate Image | 90度回転 |
| Spiral Matrix | スパイラル走査 |
| **Set Matrix Zeroes** | 条件付き更新 |
| Game of Life | 同時更新 |
