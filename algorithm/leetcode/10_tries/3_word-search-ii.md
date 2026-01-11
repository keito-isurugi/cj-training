# Word Search II (Hard)

## 問題内容

2次元の文字グリッド `board` と文字列リスト `words` が与えられる。グリッド内に存在する全ての単語を返す。

単語が存在するとは、水平または垂直に隣接するセルでパスを形成できること。同じセルは1つの単語内で複数回使用できない。

### 例

```
Input:
board = [
  ["a","b","c","d"],
  ["s","a","a","t"],
  ["a","c","k","e"],
  ["a","c","d","n"]
],
words = ["bat","cat","back","backend","stack"]

Output: ["cat","back","backend"]
```

```
Input:
board = [
  ["x","o"],
  ["x","o"]
],
words = ["xoxo"]

Output: []
```

### 制約

- `1 <= board.length, board[i].length <= 12`
- `board[i]` は小文字英字のみ
- `1 <= words.length <= 30,000`
- `1 <= words[i].length <= 10`
- `words[i]` は小文字英字のみ
- `words` 内の文字列は全て異なる

## ソースコード

```go
type TrieNode struct {
	children [26]*TrieNode
	idx      int
	refs     int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{idx: -1}
}

func (this *TrieNode) addWord(word string, i int) {
	cur := this
	cur.refs++
	for _, ch := range word {
		index := ch - 'a'
		if cur.children[index] == nil {
			cur.children[index] = NewTrieNode()
		}
		cur = cur.children[index]
		cur.refs++
	}
	cur.idx = i
}

func findWords(board [][]byte, words []string) []string {
    root := NewTrieNode()
	for i, word := range words {
		root.addWord(word, i)
	}

	rows, cols := len(board), len(board[0])
	var res []string

	getIndex := func(c byte) int { return int(c - 'a') }

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		if r < 0 || c < 0 || r >= rows || c >= cols ||
           board[r][c] == '*' || node.children[getIndex(board[r][c])] == nil {
			return
		}

		tmp := board[r][c]
		board[r][c] = '*'
		prev := node
		node = node.children[getIndex(tmp)]
		if node.idx != -1 {
			res = append(res, words[node.idx])
			node.idx = -1
			node.refs--
			if node.refs == 0 {
				prev.children[getIndex(tmp)] = nil
				board[r][c] = tmp
				return
			}
		}

		dfs(r+1, c, node)
		dfs(r-1, c, node)
		dfs(r, c+1, node)
		dfs(r, c-1, node)

		board[r][c] = tmp
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return res
}
```

## アルゴリズムなど解説

### 基本戦略

Trie + DFS + バックトラッキング。全単語をTrieに格納し、グリッドを探索しながらTrieを辿る。

### なぜ Trie を使うか

```
ナイーブな方法（Word Search を単語ごとに実行）:
- 30,000単語 × グリッド探索
- TLE（時間超過）確実

Trie を使う方法:
- 全単語を1つのTrieに格納
- 1回の探索で複数単語を同時にチェック
- 共通プレフィックスを共有
```

### TrieNode の構造

```go
type TrieNode struct {
	children [26]*TrieNode
	idx      int   // 単語のインデックス（-1なら単語終端でない）
	refs     int   // このノードを参照する単語数
}
```

```
words = ["cat", "car", "card"]

Trie構造:
      root (refs=3)
       ↓
       c (refs=3)
       ↓
       a (refs=3)
      / \
     t   r (refs=2)
  idx=0  |
         r (refs=1)
       idx=1
         ↓
         d (refs=1)
       idx=2

refs は「このノード以下にいくつの単語があるか」
```

### refs の役割（枝刈り）

```
"cat" を見つけた後:

refs を減らす:
- t.refs: 1 → 0
- a.refs: 3 → 2
- c.refs: 3 → 2
- root.refs: 3 → 2

refs == 0 のノードを削除:
→ 't' ノードを削除

      root (refs=2)
       ↓
       c (refs=2)
       ↓
       a (refs=2)  ← 't' への枝が消えた！
        \
         r (refs=2)
         ...

次回の探索で "cat" のパスは辿らない
→ 無駄な探索を削減！
```

### DFS の動作

```go
dfs = func(r, c int, node *TrieNode) {
    // 終了条件
    if r < 0 || c < 0 || r >= rows || c >= cols ||
       board[r][c] == '*' || node.children[getIndex(board[r][c])] == nil {
        return
    }

    // 現在のセルを処理
    tmp := board[r][c]
    board[r][c] = '*'  // 訪問済みマーク
    prev := node
    node = node.children[getIndex(tmp)]

    // 単語を見つけた場合
    if node.idx != -1 {
        res = append(res, words[node.idx])
        node.idx = -1  // 重複防止
        node.refs--
        if node.refs == 0 {
            prev.children[getIndex(tmp)] = nil  // 枝刈り
            board[r][c] = tmp
            return
        }
    }

    // 4方向に探索
    dfs(r+1, c, node)
    dfs(r-1, c, node)
    dfs(r, c+1, node)
    dfs(r, c-1, node)

    board[r][c] = tmp  // バックトラック
}
```

### 視覚的な理解

```
board:                 words: ["cat", "back"]
a b c d
s a a t
a c k e
a c d n

Trie:
      root
     /    \
    b      c
    |      |
    a      a
    |      |
    c      t
    |    idx=0
    k
  idx=1

━━━ 探索: (0,2)='c' から開始 ━━━

Step 1: 'c' → Trie: root → c
        board[0][2] = '*'

Step 2: 4方向を探索
        (1,2)='a' → Trie: c → a
        board[1][2] = '*'

Step 3: 4方向を探索
        (1,3)='t' → Trie: a → t
        idx=0 → "cat" を発見！
        res = ["cat"]
        refs を更新、idx = -1 に
```

### 重複防止

```go
if node.idx != -1 {
    res = append(res, words[node.idx])
    node.idx = -1  // ← これが重要！
```

```
同じ単語を複数のパスで見つける可能性がある:

board:
c a t
a   a
t a c

"cat" は (0,0)→(0,1)→(0,2) でも
        (0,0)→(1,0)→(2,0) でも形成可能

idx = -1 にすることで、2回目以降は無視
```

### 枝刈りの詳細

```go
if node.refs == 0 {
    prev.children[getIndex(tmp)] = nil  // 親から切り離す
    board[r][c] = tmp
    return
}
```

```
refs が 0 になった = このノード以下に未発見の単語がない

例: "cat" のみを格納したTrie
    root → c → a → t

"cat" 発見後:
- t.refs = 0 → 削除
- a.refs = 0 → 削除
- c.refs = 0 → 削除

次の探索では root.children['c'] = nil
→ 'c' から始まるパスは即座にスキップ
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(m × n × 4^L + W × L) | m×n=グリッド, L=最長単語, W=単語数 |
| 空間計算量 | O(W × L) | Trieのサイズ |

### 最適化まとめ

```
1. Trie による共通プレフィックス共有
   → 複数単語を同時に検索

2. refs によるノード削除
   → 見つけた単語のパスを枝刈り

3. idx = -1 による重複防止
   → 同じ単語を2回追加しない

4. board 書き換えによる訪問管理
   → 別途 visited 配列が不要
```

### 別解：シンプル版（最適化なし）

```go
type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}

	// Trie構築
	for _, word := range words {
		cur := root
		for _, c := range word {
			idx := c - 'a'
			if cur.children[idx] == nil {
				cur.children[idx] = &TrieNode{}
			}
			cur = cur.children[idx]
		}
		cur.word = word
	}

	rows, cols := len(board), len(board[0])
	var res []string

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		c := board[r][c]
		if c == '*' || node.children[c-'a'] == nil {
			return
		}

		node = node.children[c-'a']
		if node.word != "" {
			res = append(res, node.word)
			node.word = ""  // 重複防止
		}

		board[r][c] = '*'
		dfs(r+1, c, node)
		dfs(r-1, c, node)
		dfs(r, c+1, node)
		dfs(r, c-1, node)
		board[r][c] = c
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return res
}
```

### Word Search I との違い

| 項目 | Word Search I | Word Search II |
|------|--------------|----------------|
| 単語数 | 1つ | 複数（最大30,000） |
| データ構造 | なし | Trie |
| 探索 | 単語ごとに探索 | 1回の探索で全単語 |
| 最適化 | 不要 | 枝刈りが重要 |

### 重要なポイント

```
1. Trie + DFS の組み合わせ
   → 複数単語を効率的に同時検索

2. refs による枝刈り
   → 見つけた単語のパスを削除
   → 無駄な探索を削減

3. idx / word による重複防止
   → 同じ単語を複数回追加しない

4. Word Search I の発展版
   → 1単語なら単純DFS
   → 複数単語ならTrie必須
```

### 関連問題

| 問題 | 共通点 |
|------|--------|
| Word Search | 1単語版（Trie不要） |
| Implement Trie | Trieの基本実装 |
| Design Add and Search Word | Trie + ワイルドカード |

**Hard** 問題。Trie + グリッドDFS + 枝刈りの総合問題！
