# Longest Substring Without Repeating Characters (Medium)

## 問題内容

文字列 `s` が与えられたとき、重複する文字を含まない最長部分文字列の長さを返す。

**部分文字列（Substring）** とは、文字列内の連続した文字の並びである。

### 例

```
Input: s = "zxyzxyz"
Output: 3
```
説明: "xyz" が重複文字なしの最長部分文字列

```
Input: s = "xxxx"
Output: 1
```

### 制約

- `0 <= s.length <= 1000`
- `s` は印刷可能なASCII文字で構成される

## ソースコード

```go
func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	res := 0
	for r := 0; r < len(s); r++ {
		for charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}
```

## アルゴリズムなど解説

### 基本戦略

Sliding Window（可変長）を使用して、重複のない部分文字列を維持しながらウィンドウを拡張・縮小する。

### Sliding Windowの考え方

- **右ポインタ `r`**: ウィンドウを拡張（新しい文字を追加）
- **左ポインタ `l`**: ウィンドウを縮小（重複を解消）
- **HashSet**: 現在のウィンドウ内の文字を追跡

ウィンドウ内に重複がある限り左から縮小し、常に有効な（重複なし）ウィンドウを維持する。

### 動作の仕組み

1. **初期化**
   ```go
   charSet := make(map[byte]bool)
   l := 0
   res := 0
   ```
   - HashSetで現在のウィンドウ内の文字を管理
   - 左ポインタと結果を初期化

2. **右ポインタでスキャン**
   ```go
   for r := 0; r < len(s); r++ {
   ```
   - 各文字を順に処理

3. **重複がある場合、左から縮小**
   ```go
   for charSet[s[r]] {
       delete(charSet, s[l])
       l++
   }
   ```
   - 追加しようとしている文字 `s[r]` が既にウィンドウ内にある場合
   - 左端の文字を削除して `l` を進める
   - 重複が解消されるまで繰り返す

4. **文字を追加して結果を更新**
   ```go
   charSet[s[r]] = true
   res = max(res, r-l+1)
   ```
   - 新しい文字をウィンドウに追加
   - ウィンドウサイズ `r-l+1` で最大長を更新

### 具体例

```
s = "zxyzxyz"

r=0: s[r]='z', charSet={}, 重複なし → charSet={'z'}, window="z", res=1
r=1: s[r]='x', charSet={'z'}, 重複なし → charSet={'z','x'}, window="zx", res=2
r=2: s[r]='y', 重複なし → charSet={'z','x','y'}, window="zxy", res=3
r=3: s[r]='z', 重複あり!
     → delete 'z', l=1, charSet={'x','y'}
     → charSet={'x','y','z'}, window="xyz", res=3
r=4: s[r]='x', 重複あり!
     → delete 'x', l=2, charSet={'y','z'}
     → charSet={'y','z','x'}, window="yzx", res=3
r=5: s[r]='y', 重複あり!
     → delete 'y', l=3, charSet={'z','x'}
     → charSet={'z','x','y'}, window="zxy", res=3
r=6: s[r]='z', 重複あり!
     → delete 'z', l=4, charSet={'x','y'}
     → charSet={'x','y','z'}, window="xyz", res=3

最終結果: 3
```

### なぜこのアルゴリズムが正しいか

- 各文字は最大で2回処理される（追加1回、削除1回）
- ウィンドウは常に有効な状態（重複なし）を維持
- すべての有効なウィンドウを効率的に探索できる

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各文字は最大2回処理 |
| 空間計算量 | O(m) | m = ユニーク文字数（最大256 for ASCII） |

### 別解：最適化版（インデックス記録）

```go
func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    l, res := 0, 0

    for r := 0; r < len(s); r++ {
        if idx, found := mp[s[r]]; found {
            l = max(idx+1, l)
        }
        mp[s[r]] = r
        if r - l + 1 > res {
            res = r - l + 1
        }
    }
    return res
}
```

- 各文字の最後のインデックスを記録
- 重複を見つけたら、左ポインタを直接ジャンプ
- 1文字ずつ削除する必要がない
