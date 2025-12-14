# Encode and Decode Strings (Medium)

## 問題内容

文字列のリストを1つの文字列にエンコードし、その文字列から元のリストにデコードするアルゴリズムを設計する。

`encode` と `decode` を実装する。

### 例

```
Input: ["neet","code","love","you"]
Output: ["neet","code","love","you"]

Input: ["we","say",":","yes"]
Output: ["we","say",":","yes"]
```

### 制約

- `0 <= strs.length < 100`
- `0 <= strs[i].length < 200`
- `strs[i]` は UTF-8 文字を含む

## ソースコード

```go
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length
	}
	return res
}
```

## アルゴリズムなど解説

### 基本戦略

各文字列の前に「長さ + デリミタ(#)」を付けることで、文字列の境界を明確にする。

### 動作の仕組み

1. **エンコード**
   ```go
   res += strconv.Itoa(len(str)) + "#" + str
   ```
   - 各文字列を `長さ#文字列` の形式に変換
   - 例: "neet" → "4#neet"

2. **デコード**
   ```go
   for i < len(encoded) {
       j := i
       for encoded[j] != '#' {
           j++
       }
       length, _ := strconv.Atoi(encoded[i:j])
       i = j + 1
       res = append(res, encoded[i:i+length])
       i += length
   }
   ```
   - `#` まで読んで長さを取得
   - 長さ分だけ文字列を読み取る
   - 次の文字列へ進む

### 具体例

#### エンコード: `["neet", "code", "love", "you"]`

```
"neet" → "4#neet"
"code" → "4#code"
"love" → "4#love"
"you"  → "3#you"

結果: "4#neet4#code4#love3#you"
```

#### デコード: `"4#neet4#code4#love3#you"`

```
i=0: j=1で'#'発見 → 長さ4 → "neet" → i=6
i=6: j=7で'#'発見 → 長さ4 → "code" → i=12
i=12: j=13で'#'発見 → 長さ4 → "love" → i=18
i=18: j=19で'#'発見 → 長さ3 → "you" → i=23

結果: ["neet", "code", "love", "you"]
```

### なぜこの方式が適切か

```
課題:
- 文字列内に任意の文字（#や,など）が含まれる可能性
- 単純なデリミタでは区切れない

解決策:
- 長さを明示することで、デリミタの曖昧さを排除
- "#" が文字列内に含まれていても問題なし
  （長さで読み取る文字数が決まるため）

例: ["a#b", "c"] → "3#a#b1#c"
  - "a#b" の長さは3なので、"#" の後3文字読む
  - 文字列内の "#" は問題にならない
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m) | encode/decode共に全文字を1回処理 |
| 空間計算量 | O(m + n) | m: 全文字数、n: 文字列の数 |
