# Alien Dictionary (Hard)

## 問題内容

ラテンアルファベットを使用する外国語があるが、文字の順序は英語のような "a", "b", "c" ... "z" ではない。

辞書から **辞書順にソートされた** 空でない文字列のリスト `words` が与えられる。

この言語での文字の順序を導き出す。順序が無効な場合は空文字列を返す。複数の有効な順序がある場合は、**どれか1つ** を返す。

文字列 `a` が文字列 `b` より辞書順で小さいとは、以下のいずれかが真の場合：
* 最初に異なる文字が `a` の方が `b` より小さい
* `a` が `b` の接頭辞であり、`a.length < b.length`

### 例

```
Input: ["z","o"]
Output: "zo"
```
説明: "z" と "o" から、'z' < 'o' とわかるので "zo" を返す

```
Input: ["hrn","hrf","er","enn","rfnn"]
Output: "hernf"
```
説明:
* "hrn" と "hrf" から 'n' < 'f'
* "hrf" と "er" から 'h' < 'e'
* "er" と "enn" から 'r' < 'n'
* "enn" と "rfnn" から 'e' < 'r'
* 可能な解の1つは "hernf"

### 制約

- 入力 `words` は小文字の 'a' から 'z' のみを含む
- `1 <= words.length <= 100`
- `1 <= words[i].length <= 100`

## ソースコード

```go
func foreignDictionary(words []string) string {
    adj := make(map[rune]map[rune]struct{})
    for _, w := range words {
        for _, c := range w {
            if _, exists := adj[c]; !exists {
                adj[c] = make(map[rune]struct{})
            }
        }
    }

    for i := 0; i < len(words)-1; i++ {
        w1, w2 := words[i], words[i+1]
        minLen := len(w1)
        if len(w2) < minLen {
            minLen = len(w2)
        }
        if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
            return ""
        }
        for j := 0; j < minLen; j++ {
            if w1[j] != w2[j] {
                adj[rune(w1[j])][rune(w2[j])] = struct{}{}
                break
            }
        }
    }

    visited := make(map[rune]int)  // 0: unvisited, 1: visiting, -1: visited
    var res []rune

    var dfs func(char rune) bool
    dfs = func(char rune) bool {
        if status, exists := visited[char]; exists {
            return status == 1  // cycle detected if visiting
        }

        visited[char] = 1  // mark as visiting

        for neighChar := range adj[char] {
            if dfs(neighChar) {
                return true
            }
        }

        visited[char] = -1  // mark as visited
        res = append(res, char)
        return false
    }

    for char := range adj {
        if dfs(char) {
            return ""
        }
    }

    // reverse result
    var result []byte
    for i := len(res) - 1; i >= 0; i-- {
        result = append(result, byte(res[i]))
    }

    return string(result)
}
```

## アルゴリズムなど解説

### 基本戦略

単語がすでに未知のアルファベット順でソートされているので、**隣接する2つの単語を比較**することで文字間の順序関係を抽出できる。これをグラフ問題として捉え、**トポロジカルソート**で順序を決定する。

### 核心となる洞察

隣接する2つの単語 `w1` と `w2` を比較するとき：
- 最初に異なる位置 `j` で `w1[j] != w2[j]` であれば
- `w1[j]` は `w2[j]` より前に来る → 有向エッジ `w1[j] → w2[j]`

これらの関係を**有向グラフ**として構築し、トポロジカル順序を求める。

### 無効なケース

1. **サイクルがある場合**: 順序が矛盾 → 空文字列を返す
2. **接頭辞の問題**: `w1` が `w2` より長く、`w2` が `w1` の接頭辞
   - 例: `["abc", "ab"]` → "abc" が "ab" より前は不可能

### 動作の仕組み

1. **グラフの構築**
   ```go
   adj := make(map[rune]map[rune]struct{})
   for _, w := range words {
       for _, c := range w {
           adj[c] = make(map[rune]struct{})
       }
   }
   ```
   - すべてのユニーク文字をノードとして登録

2. **エッジの追加**
   ```go
   for i := 0; i < len(words)-1; i++ {
       w1, w2 := words[i], words[i+1]
       // ...
       for j := 0; j < minLen; j++ {
           if w1[j] != w2[j] {
               adj[rune(w1[j])][rune(w2[j])] = struct{}{}
               break
           }
       }
   }
   ```
   - 隣接単語の最初の異なる文字からエッジを作成

3. **DFS（3状態の訪問追跡）**
   ```go
   visited := make(map[rune]int)  // 0: unvisited, 1: visiting, -1: visited

   dfs = func(char rune) bool {
       if status, exists := visited[char]; exists {
           return status == 1  // visiting状態ならサイクル
       }
       visited[char] = 1  // 訪問中にマーク
       // 隣接ノードを探索...
       visited[char] = -1  // 完了にマーク
       res = append(res, char)  // 後順で追加
       return false
   }
   ```

4. **結果の反転**
   ```go
   for i := len(res) - 1; i >= 0; i-- {
       result = append(result, byte(res[i]))
   }
   ```
   - DFSの後順は逆順になるので反転

### 3状態の訪問追跡

| 状態 | 値 | 意味 |
|------|-----|------|
| unvisited | 0 (存在しない) | まだ訪問していない |
| visiting | 1 | 現在のDFSパス上にある |
| visited | -1 | 処理完了 |

visiting 状態のノードを再訪問 → サイクル検出

### 具体例

```
words = ["hrn", "hrf", "er", "enn", "rfnn"]

比較:
"hrn" vs "hrf" → 'n' < 'f' → n → f
"hrf" vs "er"  → 'h' < 'e' → h → e
"er"  vs "enn" → 'r' < 'n' → r → n
"enn" vs "rfnn" → 'e' < 'r' → e → r

グラフ:
h → e → r → n → f

DFS (後順):
f を訪問完了 → res = [f]
n を訪問完了 → res = [f, n]
r を訪問完了 → res = [f, n, r]
e を訪問完了 → res = [f, n, r, e]
h を訪問完了 → res = [f, n, r, e, h]

反転: "hernf"
```

### なぜ後順（Postorder）を使うか

- DFSで深く進んでから戻ってくる順序が後順
- 後順 = すべての後続ノードを処理した後に自分を追加
- これを反転すると、すべての前提条件が満たされた順序（トポロジカル順序）になる

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(N + V + E) | N = 全文字列長の和、V = ユニーク文字数、E = エッジ数 |
| 空間計算量 | O(V + E) | グラフと訪問配列 |

### 別解：カーン法（BFS）

```go
func foreignDictionary(words []string) string {
    adj := make(map[byte]map[byte]struct{})
    indegree := make(map[byte]int)

    for _, word := range words {
        for i := 0; i < len(word); i++ {
            char := word[i]
            if _, exists := adj[char]; !exists {
                adj[char] = make(map[byte]struct{})
            }
            indegree[char] = 0
        }
    }

    for i := 0; i < len(words)-1; i++ {
        w1, w2 := words[i], words[i+1]
        minLen := len(w1)
        if len(w2) < minLen {
            minLen = len(w2)
        }

        if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
            return ""
        }

        for j := 0; j < minLen; j++ {
            if w1[j] != w2[j] {
                if _, exists := adj[w1[j]][w2[j]]; !exists {
                    adj[w1[j]][w2[j]] = struct{}{}
                    indegree[w2[j]]++
                }
                break
            }
        }
    }

    q := []byte{}
    for char := range indegree {
        if indegree[char] == 0 {
            q = append(q, char)
        }
    }

    res := []byte{}
    for len(q) > 0 {
        char := q[0]
        q = q[1:]
        res = append(res, char)

        for neighbor := range adj[char] {
            indegree[neighbor]--
            if indegree[neighbor] == 0 {
                q = append(q, neighbor)
            }
        }
    }

    if len(res) != len(indegree) {
        return ""
    }

    return string(res)
}
```

**カーン法のアイデア**：
- 各ノードの**入次数**（そのノードに入ってくるエッジの数）を計算
- 入次数 0 のノード = 前提条件がない → 先に処理可能
- 処理したノードからのエッジを「削除」し、入次数を減らす
- すべてのノードを処理できなければサイクルあり

### グラフ問題のパターン

この問題は **トポロジカルソート** パターン：
- 順序関係から有向グラフを構築
- DFS後順 または カーン法（BFS）でトポロジカル順序を求める
- サイクル検出が重要

### 関連問題との比較

| 問題 | 目的 | 手法 |
|------|------|------|
| Course Schedule | 全コース完了可能か | サイクル検出 |
| Course Schedule II | コースの順序 | トポロジカルソート |
| **Alien Dictionary** | 文字の順序を復元 | 順序関係からグラフ構築 + トポロジカルソート |
