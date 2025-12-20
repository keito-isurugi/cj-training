# Word Break (Medium)

## 問題内容

文字列 `s` と文字列の辞書 `wordDict` が与えられたとき、`s` を辞書内の単語のスペース区切りのシーケンスに分割できる場合は `true` を返す。

辞書内の単語は**無制限に再利用**できる。すべての辞書単語は一意であると仮定してよい。

### 例

```
Input: s = "neetcode", wordDict = ["neet","code"]
Output: true
```
説明: "neetcode" は "neet" と "code" に分割できる。

```
Input: s = "applepenapple", wordDict = ["apple","pen","ape"]
Output: true
```
説明: "applepenapple" は "apple", "pen", "apple" に分割できる。単語の再利用が可能で、すべての単語を使う必要はない。

```
Input: s = "catsincars", wordDict = ["cats","cat","sin","in","car"]
Output: false
```

### 制約

- `1 <= s.length <= 200`
- `1 <= wordDict.length <= 100`
- `1 <= wordDict[i].length <= 20`
- `s` と `wordDict[i]` は小文字の英字のみを含む

## ソースコード

```go
func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n+1)
    dp[n] = true

    for i := n - 1; i >= 0; i-- {
        for _, w := range wordDict {
            if i+len(w) <= n && s[i:i+len(w)] == w {
                dp[i] = dp[i+len(w)]
            }
            if dp[i] {
                break
            }
        }
    }

    return dp[0]
}
```

## アルゴリズムなど解説

### 基本戦略

各インデックス `i` で：
> "接尾辞 `s[i:]` を辞書の単語に分割できるか？"

辞書の**すべての単語**を試す：
- 単語が位置 `i` から一致するなら
- 残りの部分 `s[i+len(word):]` が分割可能か再帰的にチェック

いずれかのパスが文字列の終わりに到達すれば `true`。

### 核心となる洞察

```
dp[i] = true ⟺ ある単語 w が s[i:] の先頭に一致 かつ dp[i+len(w)] = true
```

### 動作の仕組み

1. **DP配列の初期化**
   ```go
   dp := make([]bool, n+1)
   dp[n] = true  // 空文字列は分割可能
   ```

2. **後ろから前へ計算**
   ```go
   for i := n - 1; i >= 0; i-- {
       for _, w := range wordDict {
           if i+len(w) <= n && s[i:i+len(w)] == w {
               dp[i] = dp[i+len(w)]
           }
           if dp[i] {
               break  // 見つかったら早期終了
           }
       }
   }
   ```

3. **結果**
   ```go
   return dp[0]  // 文字列全体が分割可能か
   ```

### 具体例

```
s = "leetcode", wordDict = ["leet", "code"]

dp[8] = true (空文字列)

i=7 ('e'): どの単語も一致しない → dp[7] = false
i=6 ('d'): どの単語も一致しない → dp[6] = false
i=5 ('o'): どの単語も一致しない → dp[5] = false
i=4 ('c'): "code" が一致 → dp[4] = dp[8] = true
i=3 ('t'): どの単語も一致しない → dp[3] = false
i=2 ('e'): どの単語も一致しない → dp[2] = false
i=1 ('e'): どの単語も一致しない → dp[1] = false
i=0 ('l'): "leet" が一致 → dp[0] = dp[4] = true

結果: true
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n × m × t) | n=文字列長, m=単語数, t=最大単語長 |
| 空間計算量 | O(n) | DP配列 |

### 別解：再帰 + メモ化

```go
func wordBreak(s string, wordDict []string) bool {
    memo := make(map[int]bool)
    memo[len(s)] = true

    var dfs func(i int) bool
    dfs = func(i int) bool {
        if val, ok := memo[i]; ok {
            return val
        }

        for _, w := range wordDict {
            if i+len(w) <= len(s) && s[i:i+len(w)] == w {
                if dfs(i + len(w)) {
                    memo[i] = true
                    return true
                }
            }
        }
        memo[i] = false
        return false
    }

    return dfs(0)
}
```

### 別解：HashSetで高速化

```go
func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool)
    for _, w := range wordDict {
        wordSet[w] = true
    }

    n := len(s)
    dp := make([]bool, n+1)
    dp[n] = true

    for i := n - 1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if wordSet[s[i:j+1]] && dp[j+1] {
                dp[i] = true
                break
            }
        }
    }

    return dp[0]
}
```

**HashSetのアイデア**：
- 単語の存在チェックを O(1) に
- すべての部分文字列を試す

### 別解：Trie

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isWord   bool
}

func wordBreak(s string, wordDict []string) bool {
    root := &TrieNode{children: make(map[rune]*TrieNode)}

    // Trieを構築
    for _, word := range wordDict {
        node := root
        for _, c := range word {
            if _, ok := node.children[c]; !ok {
                node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
            }
            node = node.children[c]
        }
        node.isWord = true
    }

    n := len(s)
    dp := make([]bool, n+1)
    dp[n] = true

    for i := n - 1; i >= 0; i-- {
        node := root
        for j := i; j < n; j++ {
            c := rune(s[j])
            if _, ok := node.children[c]; !ok {
                break
            }
            node = node.children[c]
            if node.isWord && dp[j+1] {
                dp[i] = true
                break
            }
        }
    }

    return dp[0]
}
```

### DPパターン

この問題は **文字列分割DP** パターン：
- 文字列を辞書の単語に分割
- 各位置で「ここから分割可能か」を判定
- 再帰的に残りの部分をチェック

### 関連問題との比較

| 問題 | 目的 | 戻り値 |
|------|------|--------|
| **Word Break** | 分割可能か | bool |
| Word Break II | すべての分割方法 | []string |
| Palindrome Partitioning | 回文で分割 | [][]string |
| Concatenated Words | 連結単語を見つける | []string |
