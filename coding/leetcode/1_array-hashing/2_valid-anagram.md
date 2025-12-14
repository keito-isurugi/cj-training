# Valid Anagram (Easy)

## 問題内容

2つの文字列 `s` と `t` が与えられたとき、`t` が `s` のアナグラムである場合は `true` を返し、そうでない場合は `false` を返す。

**アナグラム**とは、別の文字列の文字を並び替えて作られる文字列のこと（文字の出現回数が同じ）。

### 例

```
Input: s = "racecar", t = "carrace"
Output: true

Input: s = "jar", t = "jam"
Output: false
```

### 制約

- `s` と `t` は小文字の英字のみで構成される

## ソースコード

```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS, countT := make(map[rune]int), make(map[rune]int)

    for i := range s {
        countS[rune(s[i])] = 1 + countS[rune(s[i])]
        countT[rune(t[i])] = 1 + countT[rune(t[i])]
    }

    if len(countS) != len(countT) {
        return false
    }
    for k, v := range countS {
        if countT[k] != v {
            return false
        }
    }
    return true
}
```

## アルゴリズムなど解説

### 基本戦略

2つのHash Mapを使って各文字の出現頻度をカウントし、それらが一致するかを確認する。

### 動作の仕組み

1. **長さチェック**
   ```go
   if len(s) != len(t) {
       return false
   }
   ```
   - 長さが異なればアナグラムではない

2. **頻度カウント**
   ```go
   for i := range s {
       countS[rune(s[i])] = 1 + countS[rune(s[i])]
       countT[rune(t[i])] = 1 + countT[rune(t[i])]
   }
   ```
   - 各文字列の文字の出現回数をカウント

3. **頻度比較**
   ```go
   for k, v := range countS {
       if countT[k] != v {
           return false
       }
   }
   ```
   - すべての文字について、出現回数が一致するか確認

### 具体例

#### 例1: `s = "racecar"`, `t = "carrace"`

```
sの頻度: {r:2, a:2, c:2, e:1}
tの頻度: {c:2, a:2, r:2, e:1}

すべての文字の頻度が一致 → true
```

#### 例2: `s = "jar"`, `t = "jam"`

```
sの頻度: {j:1, a:1, r:1}
tの頻度: {j:1, a:1, m:1}

'r' と 'm' が異なる → false
```

### なぜHash Mapが適切か

```
アナグラムの定義:
- 同じ文字を同じ回数使用している
- 順序は関係ない

Hash Mapの利点:
- 各文字の出現回数を効率的に記録
- O(1)での挿入・検索が可能
- ソートよりも効率的（O(n) vs O(n log n)）
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n + m) | 両文字列を1回ずつスキャン |
| 空間計算量 | O(1) | 最大26文字分のスペース（小文字英字のみ） |
