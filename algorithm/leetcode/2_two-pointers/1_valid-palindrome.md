# Valid Palindrome (Easy)

## 問題内容

文字列 `s` が与えられたとき、それが回文（palindrome）であれば `true` を、そうでなければ `false` を返す。

回文とは、前から読んでも後ろから読んでも同じ文字列のことを指す。大文字・小文字は区別せず、英数字以外の文字は無視する。

### 例

```
Input: s = "Was it a car or a cat I saw?"
Output: true

説明: 英数字のみを考慮すると "wasitacaroracatisaw" となり、回文である。
```

```
Input: s = "tab a cat"
Output: false

説明: "tabacat" は回文ではない。
```

### 制約

- `1 <= s.length <= 1000`
- `s` は印刷可能なASCII文字のみで構成される

## ソースコード

```go
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !alphaNum(s[l]) {
			l++
		}
		for r > l && !alphaNum(s[r]) {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func alphaNum(c byte) bool {
	return ('A' <= c && c <= 'Z') ||
		('a' <= c && c <= 'z') ||
		('0' <= c && c <= '9')
}
```

## アルゴリズムなど解説

### 基本戦略

Two Pointers（2つのポインタ）テクニックを使用して、文字列の両端から中央に向かって比較を行う。

### Two Pointersテクニックとは

配列や文字列に対して、2つのポインタを使って効率的に探索を行うテクニック。このテクニックにより、追加のメモリを使わずに問題を解決できる。

### 動作の仕組み

1. **2つのポインタを初期化**
   ```go
   l, r := 0, len(s)-1
   ```
   - `l`: 左端からスタート
   - `r`: 右端からスタート

2. **非英数字をスキップ**
   ```go
   for l < r && !alphaNum(s[l]) {
       l++
   }
   for r > l && !alphaNum(s[r]) {
       r--
   }
   ```
   - 英数字以外の文字は回文判定に関係ないのでスキップ

3. **文字を比較**
   ```go
   if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
       return false
   }
   ```
   - 小文字に変換して比較（大文字小文字を区別しない）
   - 一致しなければ回文ではない

4. **ポインタを移動**
   ```go
   l, r = l+1, r-1
   ```
   - 両端から中央に向かって進める

### 具体例

```
s = "A man, a plan, a canal: Panama"

処理の流れ:
l=0 ('A'), r=29 ('a') → 'a' == 'a' ✓
l=1 (' ') → スキップ → l=2 ('m')
r=28 ('m') → 'm' == 'm' ✓
l=3 ('a'), r=27 ('a') → 'a' == 'a' ✓
...（続く）...
最終的に l >= r となり、true を返す
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 文字列を1回スキャン |
| 空間計算量 | O(1) | 追加のメモリを使用しない |

### 別解：新しい文字列を作成する方法

```go
func isPalindrome(s string) bool {
    newStr := ""
    for _, c := range s {
        if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
            newStr += string(c)
        } else if 'A' <= c && c <= 'Z' {
            newStr += string(c + 'a' - 'A')
        }
    }

    reversedStr := reverse(newStr)
    return newStr == reversedStr
}

func reverse(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}
```

- 時間計算量: O(n)
- 空間計算量: O(n) - 新しい文字列を作成するため

Two Pointersを使う方法の方が空間効率が良い。
