# Rotate Image (Medium)

## 問題内容

整数からなる `n x n` の正方行列 `matrix` が与えられる。行列を90度時計回りに回転させる。

行列を **インプレース** で回転させなければならない。別の2次元配列を確保してはいけない。

### 例

```
Input: matrix = [
  [1,2],
  [3,4]
]

Output: [
  [3,1],
  [4,2]
]
```

```
Input: matrix = [
  [1,2,3],
  [4,5,6],
  [7,8,9]
]

Output: [
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
```

### 制約

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

## ソースコード

```go
func rotate(matrix [][]int) {
    l, r := 0, len(matrix)-1

    for l < r {
        for i := 0; i < r-l; i++ {
            top, bottom := l, r

            // save the topleft
            topLeft := matrix[top][l+i]

            // move bottom left into top left
            matrix[top][l+i] = matrix[bottom-i][l]

            // move bottom right into bottom left
            matrix[bottom-i][l] = matrix[bottom][r-i]

            // move top right into bottom right
            matrix[bottom][r-i] = matrix[top+i][r]

            // move top left into top right
            matrix[top+i][r] = topLeft
        }
        r--
        l++
    }
}
```

## アルゴリズムなど解説

### 基本戦略（4セルローテーション）

**発想**: 行列を **レイヤーごと** に処理し、外側から内側へ回転させる。

各レイヤーで4つの要素を同時に入れ替える：
- **左上 → 右上**
- **右上 → 右下**
- **右下 → 左下**
- **左下 → 左上**

### 動作の仕組み

1. **2つのポインタで境界を管理**
   ```go
   l = 0       // 左境界
   r = n - 1   // 右境界
   ```

2. **外側から内側へ処理**
   ```go
   for l < r {
       for i := 0; i < r-l; i++ {
           // 4つの要素を回転
       }
       r--
       l++
   }
   ```

3. **4要素の循環入れ替え**
   ```go
   topLeft := matrix[top][l+i]              // 左上を保存
   matrix[top][l+i] = matrix[bottom-i][l]   // 左下 → 左上
   matrix[bottom-i][l] = matrix[bottom][r-i] // 右下 → 左下
   matrix[bottom][r-i] = matrix[top+i][r]   // 右上 → 右下
   matrix[top+i][r] = topLeft               // 左上 → 右上
   ```

### 具体例

```
3x3行列の場合:

元:     回転後:
1 2 3   7 4 1
4 5 6 → 8 5 2
7 8 9   9 6 3

外側レイヤー:
- (0,0)→(0,2)→(2,2)→(2,0) の4要素を回転
- (0,1)→(1,2)→(2,1)→(1,0) の4要素を回転

内側（中央）は回転不要
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n²) | 全要素を1回処理 |
| 空間計算量 | O(1) | インプレース操作 |

### 別解：転置 + 反転

```go
func rotate(matrix [][]int) {
    // 行列を上下反転
    for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
        matrix[i], matrix[j] = matrix[j], matrix[i]
    }

    // 転置（対角線で入れ替え）
    for i := 0; i < len(matrix); i++ {
        for j := i + 1; j < len(matrix); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}
```

**この方法の考え方**:
1. 行列を上下反転（reverse）
2. 転置（transpose）

この2ステップで90度時計回り回転と同じ結果になる。

### なぜ転置 + 反転が有効か

```
元:     反転後:   転置後:
1 2 3   7 8 9    7 4 1
4 5 6 → 4 5 6  → 8 5 2
7 8 9   1 2 3    9 6 3
```

数学的に、90度時計回り回転は「上下反転」と「転置」の合成と等価。

### 2つの解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|------------|------------|------|
| 4セルローテーション | O(n²) | O(1) | 直感的、1パスで完了 |
| 転置 + 反転 | O(n²) | O(1) | コードがシンプル |

### Matrix パターン

この問題は **行列操作** の基本パターン：
- レイヤーごとの処理
- インプレース入れ替え
- 座標変換の理解

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Rotate Image** | 90度回転 |
| Spiral Matrix | スパイラル走査 |
| Set Matrix Zeroes | 条件付き更新 |
| Transpose Matrix | 転置操作 |
