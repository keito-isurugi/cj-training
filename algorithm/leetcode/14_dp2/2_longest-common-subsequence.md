# Longest Common Subsequence (Medium)

## 問題内容

2つの文字列 `text1` と `text2` が与えられたとき、両者の**最長共通部分列**の長さを返す。存在しない場合は `0` を返す。

**部分列**とは、元の文字列から一部の文字を削除（または削除しない）して、残りの文字の相対的な順序を変えずに得られる列のこと。

- 例: `"cat"` は `"crabt"` の部分列

**共通部分列**とは、両方の文字列に存在する部分列のこと。

### 例

```
Input: text1 = "cat", text2 = "crabt"
Output: 3
```
説明: 最長共通部分列は "cat" で、長さは 3。

```
Input: text1 = "abcd", text2 = "abcd"
Output: 4
```

```
Input: text1 = "abcd", text2 = "efgh"
Output: 0
```

### 制約

- `1 <= text1.length, text2.length <= 1000`
- `text1` と `text2` は小文字のアルファベットのみで構成される

## ソースコード

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    if len(text1) < len(text2) {
        text1, text2 = text2, text1
    }

    dp := make([]int, len(text2)+1)

    for i := len(text1) - 1; i >= 0; i-- {
        prev := 0
        for j := len(text2) - 1; j >= 0; j-- {
            temp := dp[j]
            if text1[i] == text2[j] {
                dp[j] = 1 + prev
            } else {
                dp[j] = max(dp[j], dp[j+1])
            }
            prev = temp
        }
    }

    return dp[0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## アルゴリズムなど解説

### 基本戦略

両方の文字列を同時に走査し、各位置で：
- **文字が一致する場合**: LCSに含め、両方のインデックスを進める
- **一致しない場合**: どちらかのインデックスを進める2つの選択肢から最大値を取る

### 核心となる洞察

```
if text1[i] == text2[j]:
    dp[i][j] = 1 + dp[i+1][j+1]
else:
    dp[i][j] = max(dp[i+1][j], dp[i][j+1])
```

この漸化式は：
- 文字が一致 → LCS長に1を加算
- 文字が不一致 → どちらかを飛ばした場合の最大値を採用

### 動作の仕組み

1. **空間最適化のアイデア**
   ```go
   dp := make([]int, len(text2)+1)
   ```
   - 2次元テーブルの代わりに1次元配列を使用
   - 各行は次の行のみに依存するため

2. **前の値の保持**
   ```go
   prev := 0
   for j := len(text2) - 1; j >= 0; j-- {
       temp := dp[j]
       // ... 更新処理 ...
       prev = temp
   }
   ```
   - `prev` は `dp[i+1][j+1]`（右下の値）を保持
   - 更新前の値を `temp` に退避

3. **更新ロジック**
   ```go
   if text1[i] == text2[j] {
       dp[j] = 1 + prev  // 対角線の値 + 1
   } else {
       dp[j] = max(dp[j], dp[j+1])  // 下 or 右の最大値
   }
   ```

### 具体例

```
text1 = "cat", text2 = "crabt"

DPテーブル（2次元版で説明）:
     c  r  a  b  t  ""
c  [ 3  2  2  1  1  0 ]
a  [ 2  2  2  1  1  0 ]
t  [ 1  1  1  1  1  0 ]
"" [ 0  0  0  0  0  0 ]

- dp[0][0] = 3 → 最長共通部分列の長さ
- 共通部分列: "c", "a", "t" → "cat"
```

### LCS構築の追跡

```
text1 = "cat", text2 = "crabt"

位置 (0,0): text1[0]='c' == text2[0]='c' ✓ → 'c' を含む
位置 (1,2): text1[1]='a' == text2[2]='a' ✓ → 'a' を含む
位置 (2,4): text1[2]='t' == text2[4]='t' ✓ → 't' を含む

LCS = "cat"
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n) | 両文字列の全ての組み合わせを処理 |
| 空間計算量 | O(min(m, n)) | 短い方の文字列長の配列のみ使用 |

### 別解：2次元DP配列

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if text1[i] == text2[j] {
                dp[i][j] = 1 + dp[i+1][j+1]
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            }
        }
    }

    return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### 別解：再帰 + メモ化（Top-Down）

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    memo := make([][]int, m+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == m || j == n {
            return 0
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }

        if text1[i] == text2[j] {
            memo[i][j] = 1 + dfs(i+1, j+1)
        } else {
            memo[i][j] = max(dfs(i+1, j), dfs(i, j+1))
        }

        return memo[i][j]
    }

    return dfs(0, 0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### DPパターン

この問題は **LCS（最長共通部分列）** パターン：
- 2つの文字列の比較問題
- 状態: `dp[i][j]` = `text1[i:]` と `text2[j:]` のLCS長
- 遷移: 文字の一致/不一致で分岐

### なぜ部分「列」であって部分「文字列」ではないのか

- **部分列（Subsequence）**: 連続していなくてよい
  - `"ace"` は `"abcde"` の部分列
- **部分文字列（Substring）**: 連続している必要がある
  - `"ace"` は `"abcde"` の部分文字列ではない

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Longest Common Subsequence** | LCS DP |
| Longest Common Substring | 連続部分文字列 DP |
| Edit Distance | LCS の応用 |
| Shortest Common Supersequence | LCS + 構築 |
