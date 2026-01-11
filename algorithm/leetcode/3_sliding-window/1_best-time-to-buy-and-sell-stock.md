# Best Time to Buy and Sell Stock (Easy)

## 問題内容

整数配列 `prices` が与えられる。`prices[i]` は `i` 日目の株価を表す。

**1日だけ**株を買い、**異なる未来の日**に売ることができる。

得られる最大利益を返す。利益を得られない場合は `0` を返す。

### 例

```
Input: prices = [10,1,5,6,7,1]
Output: 6
```
説明: `prices[1]` で買い、`prices[4]` で売る。利益 = 7 - 1 = 6

```
Input: prices = [10,8,7,5,2]
Output: 0
```
説明: 利益が出る取引がないため、最大利益は 0

### 制約

- `1 <= prices.length <= 100`
- `0 <= prices[i] <= 100`

## ソースコード

```go
func maxProfit(prices []int) int {
	res := 0

	lowest := prices[0]
	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		res = max(res, price-lowest)
	}
	return res
}
```

## アルゴリズムなど解説

### 基本戦略

配列を1回スキャンしながら、「これまでの最安値」と「現在の価格との差（利益）」を追跡する。

### Sliding Windowの考え方

この問題はSliding Windowパターンの一種として捉えられる：
- **左ポインタ（買い日）**: これまでの最安値の位置
- **右ポインタ（売り日）**: 現在スキャンしている位置

ウィンドウを「買い日から売り日まで」と考え、最大利益を追跡する。

### 動作の仕組み

1. **初期化**
   ```go
   res := 0
   lowest := prices[0]
   ```
   - 最大利益を0で初期化
   - 最安値を最初の価格で初期化

2. **配列をスキャン**
   ```go
   for _, price := range prices {
       if price < lowest {
           lowest = price
       }
       res = max(res, price-lowest)
   }
   ```
   - 現在の価格が最安値より低ければ、最安値を更新
   - 現在の価格 - 最安値 = 現在売った場合の利益
   - 最大利益を更新

### 具体例

```
prices = [10, 1, 5, 6, 7, 1]

i=0: price=10, lowest=10, profit=0,  res=0
i=1: price=1,  lowest=1,  profit=0,  res=0  ← 最安値更新
i=2: price=5,  lowest=1,  profit=4,  res=4
i=3: price=6,  lowest=1,  profit=5,  res=5
i=4: price=7,  lowest=1,  profit=6,  res=6  ← 最大利益
i=5: price=1,  lowest=1,  profit=0,  res=6

最終結果: 6
```

### なぜこのアルゴリズムが正しいか

- **最大利益 = 売値 - 買値** を最大化したい
- 各日を「売り日」と仮定すると、最大利益を得るには**その日より前の最安値で買う**のが最適
- 配列を順にスキャンしながら最安値を更新することで、各位置での最適な買い日を O(1) で参照できる

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回スキャン |
| 空間計算量 | O(1) | 変数のみ使用 |

### 別解：Two Pointers版

```go
func maxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0

    for r < len(prices) {
        if prices[l] < prices[r] {
            profit := prices[r] - prices[l]
            if profit > maxP {
                maxP = profit
            }
        } else {
            l = r
        }
        r++
    }
    return maxP
}
```

- `l` = 買い日（最安値の候補）
- `r` = 売り日（現在スキャン中）
- `prices[r] < prices[l]` の場合、より安い買い日が見つかったので `l` を `r` に移動
