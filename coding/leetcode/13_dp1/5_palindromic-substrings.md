# Palindromic Substrings (Medium)

## 問題内容

文字列 `s` が与えられたとき、`s` の中の**回文である部分文字列の数**を返す。

**回文**とは、前から読んでも後ろから読んでも同じ文字列のこと。

### 例

```
Input: s = "abc"
Output: 3
```
説明: "a", "b", "c" の3つ。

```
Input: s = "aaa"
Output: 6
```
説明: "a", "a", "a", "aa", "aa", "aaa" の6つ。同じ内容でも異なる位置の部分文字列は別々にカウントする。

### 制約

- `1 <= s.length <= 1000`
- `s` は小文字の英字のみを含む

## ソースコード

```go
func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        // 奇数長の回文
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            res++
            l--
            r++
        }

        // 偶数長の回文
        l, r = i, i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            res++
            l--
            r++
        }
    }

    return res
}
```

## アルゴリズムなど解説

### 基本戦略

すべての回文には**中心**がある：
- **奇数長**の回文 → 1文字が中心
- **偶数長**の回文 → 2文字の間が中心

各インデックスを中心として、**左右に拡張**しながら回文を数える。
拡張が成功するたびに**新しい回文**が1つ見つかる。

### 核心となる洞察

```
中心から拡張するたびに、その時点での部分文字列は回文である
→ 拡張ごとにカウントを+1
```

### 動作の仕組み

1. **奇数長の回文カウント**
   ```go
   l, r := i, i
   for l >= 0 && r < len(s) && s[l] == s[r] {
       res++
       l--
       r++
   }
   ```
   - 中心の1文字自体が回文（長さ1）
   - 拡張が成功するたびに新しい回文

2. **偶数長の回文カウント**
   ```go
   l, r = i, i+1
   for l >= 0 && r < len(s) && s[l] == s[r] {
       res++
       l--
       r++
   }
   ```
   - 2文字が一致すれば回文（長さ2）
   - 同様に拡張

### 具体例

```
s = "aaa"

i=0 ('a'):
  奇数: l=0,r=0 → "a" (res=1)
         l=-1,r=1 → 終了
  偶数: l=0,r=1 → "aa" (res=2)
         l=-1,r=2 → 終了

i=1 ('a'):
  奇数: l=1,r=1 → "a" (res=3)
         l=0,r=2 → "aaa" (res=4)
         l=-1,r=3 → 終了
  偶数: l=1,r=2 → "aa" (res=5)
         l=0,r=3 → 終了

i=2 ('a'):
  奇数: l=2,r=2 → "a" (res=6)
         l=1,r=3 → 終了
  偶数: l=2,r=3 → 範囲外

結果: 6
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n²) | 各中心からO(n)の拡張 |
| 空間計算量 | O(1) | 追加の配列不使用 |

### 別解：動的計画法

```go
func countSubstrings(s string) int {
    n := len(s)
    res := 0

    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    for i := n - 1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
                dp[i][j] = true
                res++
            }
        }
    }

    return res
}
```

**DPのアイデア**：
- `dp[i][j] = true` は `s[i:j+1]` が回文
- 部分文字列 `s[i:j+1]` が回文の条件：
  - `s[i] == s[j]` かつ
  - 長さが3以下、または `dp[i+1][j-1]` が真

### 別解：ブルートフォース

```go
func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            l, r := i, j
            for l < r && s[l] == s[r] {
                l++
                r--
            }
            if l >= r {
                res++
            }
        }
    }

    return res
}
```
- 時間計算量: O(n³)
- すべての部分文字列をチェック

### 回文のカウントパターン

この問題は **中心拡張法によるカウント** パターン：
- 最長を求めるのではなく、**すべての回文を数える**
- 拡張が成功するたびにカウント
- Longest Palindromic Substring と同じ手法

### 関連問題との比較

| 問題 | 目的 | 戻り値 |
|------|------|--------|
| Longest Palindromic Substring | 最長の回文 | 文字列 |
| **Palindromic Substrings** | 回文の数 | 整数 |
| Valid Palindrome | 回文判定 | bool |
| Count Binary Substrings | 特定パターンの数 | 整数 |
