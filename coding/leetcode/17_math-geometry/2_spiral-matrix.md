# Spiral Matrix (Medium)

## 問題内容

整数からなる `m x n` の行列 `matrix` が与えられる。行列のすべての要素を **スパイラル順** でリストとして返す。

### 例

```
Input: matrix = [[1,2],[3,4]]

Output: [1,2,4,3]
```

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]

Output: [1,2,3,6,9,8,7,4,5]
```

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]

Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

### 制約

- `1 <= matrix.length, matrix[i].length <= 10`
- `-100 <= matrix[i][j] <= 100`

## ソースコード

```go
func spiralOrder(matrix [][]int) []int {
    res := []int{}
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return res
    }

    left, right := 0, len(matrix[0])
    top, bottom := 0, len(matrix)

    for left < right && top < bottom {
        for i := left; i < right; i++ {
            res = append(res, matrix[top][i])
        }
        top++

        for i := top; i < bottom; i++ {
            res = append(res, matrix[i][right-1])
        }
        right--

        if !(left < right && top < bottom) {
            break
        }

        for i := right - 1; i >= left; i-- {
            res = append(res, matrix[bottom-1][i])
        }
        bottom--

        for i := bottom - 1; i >= top; i-- {
            res = append(res, matrix[i][left])
        }
        left++
    }

    return res
}
```

## アルゴリズムなど解説

### 基本戦略（4つの境界）

**スパイラル順**: 右 → 下 → 左 → 上 を繰り返し、内側に進む

4つの境界ポインタで未処理領域を管理：
- `top` - 上の行境界
- `bottom` - 下の行境界（の次）
- `left` - 左の列境界
- `right` - 右の列境界（の次）

### 動作の仕組み

1. **初期化**
   ```go
   left, right := 0, len(matrix[0])
   top, bottom := 0, len(matrix)
   ```

2. **4方向を順番に処理**
   ```go
   for left < right && top < bottom {
       // 1. 上の行を左から右へ
       for i := left; i < right; i++ {
           res = append(res, matrix[top][i])
       }
       top++

       // 2. 右の列を上から下へ
       for i := top; i < bottom; i++ {
           res = append(res, matrix[i][right-1])
       }
       right--

       // 境界チェック（1行または1列の場合）
       if !(left < right && top < bottom) {
           break
       }

       // 3. 下の行を右から左へ
       for i := right - 1; i >= left; i-- {
           res = append(res, matrix[bottom-1][i])
       }
       bottom--

       // 4. 左の列を下から上へ
       for i := bottom - 1; i >= top; i-- {
           res = append(res, matrix[i][left])
       }
       left++
   }
   ```

### 具体例

```
3x3行列:
1 2 3
4 5 6
7 8 9

スパイラル順:
1 → 2 → 3 (上の行)
      ↓
      6 (右の列)
      ↓
      9 → 8 → 7 (下の行)
↓
4 (左の列)
↓
5 (中央)

結果: [1,2,3,6,9,8,7,4,5]
```

### 境界チェックが必要な理由

```
1行の場合: [[1,2,3]]
- 上の行: 1,2,3
- 下の行を処理しようとすると重複

1列の場合: [[1],[2],[3]]
- 左の列を処理しようとすると重複
```

`if !(left < right && top < bottom)` で早期終了。

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m×n) | 全要素を1回処理 |
| 空間計算量 | O(1) | 出力配列を除く |

### 別解：方向ベクトル

```go
func spiralOrder(matrix [][]int) []int {
    res := []int{}
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return res
    }

    directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    steps := []int{len(matrix[0]), len(matrix) - 1}

    r, c, d := 0, -1, 0

    for steps[d&1] > 0 {
        for i := 0; i < steps[d&1]; i++ {
            r += directions[d][0]
            c += directions[d][1]
            res = append(res, matrix[r][c])
        }
        steps[d&1]--
        d = (d + 1) % 4
    }

    return res
}
```

**方向ベクトルの考え方**:
- 4つの方向を配列で管理
- 各方向で進める歩数を管理
- 方向を切り替えるたびに歩数を減らす

### 2つの解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|------------|------------|------|
| 4境界法 | O(m×n) | O(1) | 直感的で理解しやすい |
| 方向ベクトル法 | O(m×n) | O(1) | コンパクト |

### Matrix パターン

この問題は **スパイラル走査** パターン：
- 境界を縮めながら処理
- 4方向を順番に処理
- 早期終了条件の確認

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| Rotate Image | 90度回転 |
| **Spiral Matrix** | スパイラル走査 |
| Spiral Matrix II | スパイラル生成 |
| Set Matrix Zeroes | 条件付き更新 |
