# Container With Most Water (Medium)

## 問題内容

整数配列 `heights` が与えられる。`heights[i]` は i 番目のバーの高さを表す。

任意の2つのバーを選んでコンテナを形成できる。コンテナが貯められる水の最大量を返す。

### 例

```
Input: height = [1,7,2,5,4,7,3,6]
Output: 36

説明: インデックス1（高さ7）とインデックス7（高さ6）を選ぶと
     幅 = 7 - 1 = 6
     高さ = min(7, 6) = 6
     面積 = 6 × 6 = 36
```

```
Input: height = [2,2,2]
Output: 4

説明: インデックス0とインデックス2を選ぶと
     幅 = 2 - 0 = 2
     高さ = min(2, 2) = 2
     面積 = 2 × 2 = 4
```

### 制約

- `2 <= height.length <= 1000`
- `0 <= height[i] <= 1000`

## ソースコード

```go
func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	res := 0
	for l < r {
		res = max(res, min(heights[l], heights[r])*(r-l))
		if heights[l] < heights[r] {
			l++
		} else if heights[r] <= heights[l] {
			r--
		}
	}

	return res
}
```

## アルゴリズムなど解説

### 基本戦略

Two Pointersを使って、最も幅の広い状態からスタートし、面積を最大化するようにポインタを移動させる。

### なぜTwo Pointersが有効か

- コンテナの面積 = `幅 × 高さ` = `(r - l) × min(heights[l], heights[r])`
- 幅が最大になるのは両端を選んだとき
- ポインタを内側に移動させると幅は必ず減少する
- したがって、面積を増やすには高さを増やす必要がある

### 動作の仕組み

1. **両端にポインタを配置**
   ```go
   l, r := 0, len(heights)-1
   ```
   - 最も幅の広い状態からスタート

2. **現在の面積を計算**
   ```go
   res = max(res, min(heights[l], heights[r])*(r-l))
   ```
   - 面積 = min(左の高さ, 右の高さ) × 幅

3. **低い方のポインタを移動**
   ```go
   if heights[l] < heights[r] {
       l++
   } else {
       r--
   }
   ```

### なぜ低い方を移動するのか

これがこの問題の核心部分。

**例で考える:**
```
heights = [1, 8, 6, 2, 5, 4, 8, 3, 7]
          l=0                    r=8

現在: 面積 = min(1, 7) × 8 = 1 × 8 = 8
```

- 高い方（右、高さ7）を動かしても:
  - 幅は減少（8 → 7）
  - 高さは最大でも1（左が1なので）
  - 面積は絶対に増えない

- 低い方（左、高さ1）を動かすと:
  - 幅は減少（8 → 7）
  - 高さは増える可能性がある（1より大きい値があれば）
  - 面積が増える可能性がある

**結論:** 低い方を動かすことでのみ、面積が増加する可能性がある。

### 具体例

```
heights = [1,7,2,5,4,7,3,6]
           0 1 2 3 4 5 6 7

l=0, r=7: min(1,6) × 7 = 7, res=7
  heights[0]=1 < heights[7]=6 → l++

l=1, r=7: min(7,6) × 6 = 36, res=36
  heights[1]=7 > heights[7]=6 → r--

l=1, r=6: min(7,3) × 5 = 15, res=36
  heights[1]=7 > heights[6]=3 → r--

l=1, r=5: min(7,7) × 4 = 28, res=36
  heights[1]=7 == heights[5]=7 → r--

l=1, r=4: min(7,4) × 3 = 12, res=36
  heights[1]=7 > heights[4]=4 → r--

l=1, r=3: min(7,5) × 2 = 10, res=36
  heights[1]=7 > heights[3]=5 → r--

l=1, r=2: min(7,2) × 1 = 2, res=36
  heights[1]=7 > heights[2]=2 → r--

l=1, r=1: 終了

結果: 36
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各要素を最大1回しか見ない |
| 空間計算量 | O(1) | 追加のメモリを使用しない |

### 正当性の証明

「低い方を動かす」戦略が最適解を見逃さない理由:

1. 現在 `l` と `r` にポインタがあり、`heights[l] < heights[r]` とする
2. `l` を固定したまま `r` を `l+1` から `r-1` のどこに動かしても:
   - 幅は減少する
   - 高さは `min(heights[l], heights[新r])` で、これは最大でも `heights[l]`
   - よって面積は現在以下
3. つまり、`l` を固定した全ての組み合わせの中で、`(l, r)` が最大
4. したがって `l` を進めて、他の可能性を探索すべき

### 別解：総当たり法

```go
func maxArea(heights []int) int {
    res := 0
    for i := 0; i < len(heights); i++ {
        for j := i + 1; j < len(heights); j++ {
            area := min(heights[i], heights[j]) * (j - i)
            if area > res {
                res = area
            }
        }
    }
    return res
}
```

- 時間計算量: O(n²)
- 空間計算量: O(1)

Two Pointersの方が圧倒的に効率的。
