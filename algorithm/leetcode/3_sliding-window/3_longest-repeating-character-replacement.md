# Longest Repeating Character Replacement (Medium)

## 問題内容

大文字英字のみで構成される文字列 `s` と整数 `k` が与えられる。文字列から最大 `k` 文字を選び、任意の大文字英字に置換できる。

最大 `k` 回の置換を行った後、同一文字のみで構成される最長部分文字列の長さを返す。

### 例

```
Input: s = "XYYX", k = 2
Output: 4
```
説明: 'X' を 'Y' に置換するか、'Y' を 'X' に置換することで全て同じ文字にできる

```
Input: s = "AAABABB", k = 1
Output: 5
```

### 制約

- `1 <= s.length <= 1000`
- `0 <= k <= s.length`

## ソースコード

```go
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)

	l := 0
	maxf := 0
	r := 0
	for r = 0; r < len(s); r++ {
		count[s[r]] = 1 + count[s[r]]
		maxf = max(maxf, count[s[r]])
		if (r-l+1)-maxf > k {
			count[s[l]]--
			l++
		}
	}
	return r - l
}
```

## アルゴリズムなど解説

### 基本戦略

Sliding Window を使用し、「ウィンドウサイズ - 最頻文字の出現回数 <= k」という条件を維持する。

### 核心となる洞察

ウィンドウ内で全ての文字を同一にするには：
- **最も出現回数が多い文字を残す**
- **それ以外の文字を置換する**

したがって、必要な置換回数は：
```
置換回数 = ウィンドウサイズ - 最頻文字の出現回数
```

この値が `k` 以下ならウィンドウは有効。

### 動作の仕組み

1. **初期化**
   ```go
   count := make(map[byte]int)
   l := 0
   maxf := 0
   ```
   - `count`: 各文字の出現回数
   - `l`: 左ポインタ
   - `maxf`: ウィンドウ内の最大出現回数

2. **右ポインタでスキャン**
   ```go
   for r = 0; r < len(s); r++ {
       count[s[r]] = 1 + count[s[r]]
       maxf = max(maxf, count[s[r]])
   ```
   - 新しい文字の出現回数を更新
   - 最大出現回数を更新

3. **ウィンドウが無効なら縮小**
   ```go
   if (r-l+1)-maxf > k {
       count[s[l]]--
       l++
   }
   ```
   - 置換回数が `k` を超えたら左から縮小
   - `if` を使用（`while` ではない）— これがポイント

4. **結果を返す**
   ```go
   return r - l
   ```
   - 最終的なウィンドウサイズが答え

### `if` vs `while` の重要な違い

このコードでは `while` ではなく `if` を使用している。これは：
- ウィンドウサイズは**縮小しても最大値は更新されない**
- 一度見つけた最大サイズ以下のウィンドウは結果に影響しない
- したがって、1つ縮小するだけで十分

### 具体例

```
s = "AAABABB", k = 1

r=0: s[r]='A', count={'A':1}, maxf=1, window=1, 置換=0 ≤ 1 OK
r=1: s[r]='A', count={'A':2}, maxf=2, window=2, 置換=0 ≤ 1 OK
r=2: s[r]='A', count={'A':3}, maxf=3, window=3, 置換=0 ≤ 1 OK
r=3: s[r]='B', count={'A':3,'B':1}, maxf=3, window=4, 置換=1 ≤ 1 OK
r=4: s[r]='A', count={'A':4,'B':1}, maxf=4, window=5, 置換=1 ≤ 1 OK
r=5: s[r]='B', count={'A':4,'B':2}, maxf=4, window=6, 置換=2 > 1 NG!
     → count['A']-- → count={'A':3,'B':2}, l=1
r=6: s[r]='B', count={'A':3,'B':3}, maxf=4, window=6, 置換=2 > 1 NG!
     → count['A']-- → count={'A':2,'B':3}, l=2

最終: r=7, l=2 → 結果 = 7 - 2 = 5
```

### maxf を減らさない理由

`maxf` は厳密には現在のウィンドウ内の最大値ではなく、「これまでに見た最大値」を保持している。

これが正しく動作する理由：
- より大きな有効ウィンドウを見つけるには、`maxf` が増加する必要がある
- `maxf` が増加しない限り、ウィンドウは拡大できない
- 古い `maxf` を使うことで、不要な縮小を避けられる

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各文字を1回処理 |
| 空間計算量 | O(m) | m = ユニーク文字数（最大26） |

### 別解：各文字をターゲットにする方法

```go
func characterReplacement(s string, k int) int {
    res := 0
    charSet := make(map[byte]bool)

    for i := 0; i < len(s); i++ {
        charSet[s[i]] = true
    }

    for c := range charSet {
        count, l := 0, 0
        for r := 0; r < len(s); r++ {
            if s[r] == c {
                count++
            }

            for (r - l + 1) - count > k {
                if s[l] == c {
                    count--
                }
                l++
            }

            res = max(res, r - l + 1)
        }
    }

    return res
}
```

- 各文字をターゲットとして、その文字だけにするために必要な置換を考える
- より直感的だが、計算量は O(26 * n) = O(n)
