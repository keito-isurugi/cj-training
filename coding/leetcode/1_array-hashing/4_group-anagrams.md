# Group Anagrams (Medium)

## 問題内容

文字列配列 `strs` が与えられたとき、すべてのアナグラムをグループ化してサブリストにまとめる。出力の順序は任意。

**アナグラム**とは、別の文字列の文字を並び替えて作られる文字列のこと。

### 例

```
Input: strs = ["act","pots","tops","cat","stop","hat"]
Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

Input: strs = ["x"]
Output: [["x"]]

Input: strs = [""]
Output: [[""]]
```

### 制約

- `1 <= strs.length <= 1000`
- `0 <= strs[i].length <= 100`
- `strs[i]` は小文字の英字のみ

## ソースコード

```go
func groupAnagrams(strs []string) [][]string {
	ans := make(map[[26]int][]string)
	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[c-'a']++
		}
		ans[count] = append(ans[count], s)
	}

	result := make([][]string, 0, len(ans))
	for _, group := range ans {
		result = append(result, group)
	}
	return result
}
```

## アルゴリズムなど解説

### 基本戦略

各文字列の文字頻度を[26]int配列として計算し、その配列をキーとしてHash Mapでグループ化する。

### 動作の仕組み

1. **頻度配列の作成**
   ```go
   count := [26]int{}
   for _, c := range s {
       count[c-'a']++
   }
   ```
   - 各文字列について、26文字分の頻度配列を作成
   - `'a'`は0番目、`'b'`は1番目... という形でカウント

2. **グループ化**
   ```go
   ans[count] = append(ans[count], s)
   ```
   - 同じ頻度配列を持つ文字列は同じキーになる
   - つまりアナグラム同士は同じグループに入る

3. **結果の構築**
   ```go
   for _, group := range ans {
       result = append(result, group)
   }
   ```
   - Map の値（各グループ）を結果配列に追加

### 具体例

#### 例: `strs = ["act", "cat", "pots"]`

```
"act" → count = [1,0,1,0,...,0,0,1,0,...] (a:1, c:1, t:1)
"cat" → count = [1,0,1,0,...,0,0,1,0,...] (a:1, c:1, t:1) ← 同じ！
"pots" → count = [0,0,0,0,...,1,1,0,1,0,...] (o:1, p:1, s:1, t:1)

結果:
  key1 → ["act", "cat"]
  key2 → ["pots"]
```

### なぜ文字頻度配列をキーに使うのか

```
アナグラムの特性:
- 同じ文字を同じ回数使用
- 順序は異なる可能性がある

キーの選択肢:
1. ソート済み文字列: "act" → "act", "cat" → "act"
   - 時間計算量: O(n * m log m) ← ソートが必要
2. 文字頻度配列: [1,0,1,...,1,...]
   - 時間計算量: O(n * m) ← より効率的

Goでは配列は比較可能なので直接Mapのキーに使える
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m * n) | m: 文字列数、n: 最長文字列の長さ |
| 空間計算量 | O(m) | 追加スペース（出力を除く） |
