# Minimum Window Substring (Hard)

## 問題内容

2つの文字列 `s` と `t` が与えられる。`t` の全ての文字（重複含む）を含む `s` の最短部分文字列を返す。

そのような部分文字列が存在しない場合は空文字列 `""` を返す。

正解は常に一意であると仮定してよい。

### 例

```
Input: s = "OUZODYXAZV", t = "XYZ"
Output: "YXAZ"
```
説明: "YXAZ" は `t` の "X", "Y", "Z" を全て含む最短部分文字列

```
Input: s = "xyz", t = "xyz"
Output: "xyz"
```

```
Input: s = "x", t = "xy"
Output: ""
```

### 制約

- `1 <= s.length <= 1000`
- `1 <= t.length <= 1000`
- `s` と `t` は大文字・小文字の英字で構成される

## ソースコード

```go
func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	countT := make(map[rune]int)
	for _, c := range t {
		countT[c]++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	resLen := math.MaxInt32
	l := 0
	window := make(map[rune]int)

	for r := 0; r < len(s); r++ {
		c := rune(s[r])
		window[c]++

		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}

			window[rune(s[l])]--
			if countT[rune(s[l])] > 0 && window[rune(s[l])] < countT[rune(s[l])] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]
}
```

## アルゴリズムなど解説

### 基本戦略

Sliding Window を使用し、`t` の全文字を含むウィンドウを見つけたら、最小化のため左から縮小する。

### 核心となるアイデア

1. **have / need カウンター**
   - `need`: `t` に含まれるユニーク文字の数
   - `have`: 現在のウィンドウで条件を満たしているユニーク文字の数
   - `have == need` のとき、ウィンドウは有効

2. **ウィンドウの拡張と縮小**
   - 右ポインタで拡張して文字を追加
   - 有効になったら左ポインタで縮小して最小化

### 動作の仕組み

1. **`t` の文字頻度を記録**
   ```go
   countT := make(map[rune]int)
   for _, c := range t {
       countT[c]++
   }
   ```

2. **初期化**
   ```go
   have, need := 0, len(countT)
   res := []int{-1, -1}
   resLen := math.MaxInt32
   l := 0
   window := make(map[rune]int)
   ```
   - `need` = `t` のユニーク文字数
   - `window` = 現在のウィンドウ内の文字頻度

3. **右ポインタでスキャン**
   ```go
   for r := 0; r < len(s); r++ {
       c := rune(s[r])
       window[c]++

       if countT[c] > 0 && window[c] == countT[c] {
           have++
       }
   ```
   - 文字をウィンドウに追加
   - その文字が `t` に含まれ、必要な数に達したら `have` をインクリメント

4. **ウィンドウが有効なら縮小を試みる**
   ```go
   for have == need {
       if (r - l + 1) < resLen {
           res = []int{l, r}
           resLen = r - l + 1
       }

       window[rune(s[l])]--
       if countT[rune(s[l])] > 0 && window[rune(s[l])] < countT[rune(s[l])] {
           have--
       }
       l++
   }
   ```
   - 現在のウィンドウが最小なら結果を更新
   - 左端の文字を削除
   - その文字が `t` に含まれ、必要な数を下回ったら `have` をデクリメント

### 具体例

```
s = "OUZODYXAZV", t = "XYZ"
countT = {'X':1, 'Y':1, 'Z':1}, need = 3

r=0: 'O', window={'O':1}, have=0
r=1: 'U', window={'O':1,'U':1}, have=0
r=2: 'Z', window={..,'Z':1}, have=1 ← 'Z'条件達成
r=3: 'O', window={..}, have=1
r=4: 'D', window={..}, have=1
r=5: 'Y', window={..,'Y':1}, have=2 ← 'Y'条件達成
r=6: 'X', window={..,'X':1}, have=3 ← 'X'条件達成, have==need!
     → ウィンドウ有効! s[0:7]="OUZODYX", len=7
     → 縮小: 'O'削除, l=1, have=3（'O'はtに無関係）
     → 縮小: 'U'削除, l=2, have=3
     → 縮小: 'Z'削除, l=3, have=2 ← 'Z'が不足, ループ終了
r=7: 'A', window={..}, have=2
r=8: 'Z', window={..,'Z':1}, have=3 ← have==need!
     → ウィンドウ有効! s[3:9]="ODYXAZ", len=6
     → 縮小: 'O'削除, l=4, have=3
     → 縮小: 'D'削除, l=5, have=3
     → s[5:9]="YXAZ", len=4 ← これが最小!
     → 縮小: 'Y'削除, l=6, have=2
r=9: 'V', window={..}, have=2

最終結果: "YXAZ"
```

### なぜ `have/need` パターンが効率的か

- 全文字の頻度を毎回比較する必要がない（O(m) の比較が不要）
- 条件を満たす文字が増減したときだけカウンターを更新
- `have == need` のチェックは O(1)

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n + m) | n = sの長さ, m = s,t のユニーク文字数 |
| 空間計算量 | O(m) | 2つのハッシュマップ |

### 類似問題との比較

| 問題 | ウィンドウの条件 | 目的 |
|------|------------------|------|
| Longest Substring Without Repeating | 重複なし | 最大化 |
| Longest Repeating Character Replacement | 置換回数 ≤ k | 最大化 |
| **Minimum Window Substring** | t の全文字を含む | **最小化** |

この問題は「条件を満たす最小ウィンドウ」を求めるため、有効になったら縮小を試みる点が特徴的。
