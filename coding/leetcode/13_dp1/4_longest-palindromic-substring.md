# Longest Palindromic Substring (Medium)

## 問題内容

文字列 `s` が与えられたとき、`s` の中で**最長の回文部分文字列**を返す。

**回文**とは、前から読んでも後ろから読んでも同じ文字列のこと。

同じ長さの回文部分文字列が複数ある場合は、どれか1つを返せばよい。

### 例

```
Input: s = "ababd"
Output: "bab"
```
説明: "aba" と "bab" の両方が有効な回答。

```
Input: s = "abbc"
Output: "bb"
```

### 制約

- `1 <= s.length <= 1000`
- `s` は数字と英字のみを含む

## ソースコード

```go
func longestPalindrome(s string) string {
    resIdx, resLen := 0, 0

    for i := 0; i < len(s); i++ {
        // 奇数長の回文
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r-l+1 > resLen {
                resIdx = l
                resLen = r - l + 1
            }
            l--
            r++
        }

        // 偶数長の回文
        l, r = i, i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r-l+1 > resLen {
                resIdx = l
                resLen = r - l + 1
            }
            l--
            r++
        }
    }

    return s[resIdx : resIdx+resLen]
}
```

## アルゴリズムなど解説

### 基本戦略

回文は**中心から対称に広がる**性質を持つ。

すべての回文には以下のいずれかの中心がある：
1. **奇数長**の回文 → 1文字が中心（例: "racecar"）
2. **偶数長**の回文 → 2文字の間が中心（例: "abba"）

各インデックスを中心として、**左右に拡張**しながら回文をチェック。

### 核心となる洞察

すべての部分文字列をチェックする代わりに：
- 各インデックスを可能な中心として扱う
- 文字が一致する限り**左右に拡張**
- 拡張中に見つかった最長の回文を追跡

### 動作の仕組み

1. **奇数長の回文チェック**
   ```go
   l, r := i, i
   for l >= 0 && r < len(s) && s[l] == s[r] {
       // 回文を更新
       l--
       r++
   }
   ```
   - 中心を1文字（`i`）として開始
   - 両端が一致する限り拡張

2. **偶数長の回文チェック**
   ```go
   l, r = i, i+1
   for l >= 0 && r < len(s) && s[l] == s[r] {
       // 回文を更新
       l--
       r++
   }
   ```
   - 中心を2文字（`i` と `i+1`）の間として開始
   - 同様に拡張

3. **結果の更新**
   ```go
   if r-l+1 > resLen {
       resIdx = l
       resLen = r - l + 1
   }
   ```
   - より長い回文が見つかったら更新

### 具体例

```
s = "babad"

i=0 ('b'):
  奇数: "b" (長さ1)
  偶数: "ba" → 不一致

i=1 ('a'):
  奇数: "a" → "bab" (長さ3) ← 最長更新
  偶数: "ab" → 不一致

i=2 ('b'):
  奇数: "b" → "aba" (長さ3)
  偶数: "ba" → 不一致

i=3 ('a'):
  奇数: "a" (長さ1)
  偶数: "ad" → 不一致

i=4 ('d'):
  奇数: "d" (長さ1)
  偶数: 範囲外

結果: "bab" (または "aba")
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n²) | 各中心からO(n)の拡張 |
| 空間計算量 | O(1) | 追加の配列不使用 |

### 別解：動的計画法

```go
func longestPalindrome(s string) string {
    n := len(s)
    resIdx, resLen := 0, 0

    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    for i := n - 1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
                dp[i][j] = true
                if resLen < j-i+1 {
                    resIdx = i
                    resLen = j - i + 1
                }
            }
        }
    }

    return s[resIdx : resIdx+resLen]
}
```

**DPのアイデア**：
- `dp[i][j] = true` は `s[i:j+1]` が回文であることを示す
- 条件: `s[i] == s[j]` かつ（長さ≤3 または `dp[i+1][j-1]`が真）

### 別解：Manacher's Algorithm（O(n)）

線形時間で最長回文を見つける高度なアルゴリズム：
- 文字間に特殊文字を挿入して偶数/奇数を統一
- 既に計算した回文情報を再利用
- 実装は複雑だが、時間計算量 O(n)

### 回文のパターン

この問題は **中心拡張法** パターン：
- 各位置を中心として回文を探索
- 奇数長と偶数長の両方をチェック
- 空間効率が良い（O(1)）

### 関連問題との比較

| 問題 | 目的 | 手法 |
|------|------|------|
| **Longest Palindromic Substring** | 最長の回文を返す | 中心拡張 |
| Palindromic Substrings | 回文の数を数える | 中心拡張 |
| Valid Palindrome | 回文かどうか判定 | Two Pointers |
| Shortest Palindrome | 最短で回文にする | KMP / 中心拡張 |
