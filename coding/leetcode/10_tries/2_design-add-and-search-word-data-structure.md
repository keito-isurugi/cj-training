# Design Add and Search Words Data Structure (Medium)

## 問題内容

新しい単語の追加と既存の単語の検索をサポートするデータ構造を設計する。

`WordDictionary` クラスを実装する:
- `void addWord(word)`: 単語を追加
- `bool search(word)`: 単語が存在すれば `true`。`word` には `.`（ドット）を含む場合があり、任意の1文字にマッチする

### 例

```
Input:
["WordDictionary", "addWord", "day", "addWord", "bay", "addWord", "may",
 "search", "say", "search", "day", "search", ".ay", "search", "b.."]

Output:
[null, null, null, null, false, true, true, true]

Explanation:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("day");
wordDictionary.addWord("bay");
wordDictionary.addWord("may");
wordDictionary.search("say"); // return false（"say"は存在しない）
wordDictionary.search("day"); // return true
wordDictionary.search(".ay"); // return true（"day", "bay", "may" がマッチ）
wordDictionary.search("b.."); // return true（"bay" がマッチ）
```

### 制約

- `1 <= word.length <= 20`
- `addWord` の `word` は小文字英字のみ
- `search` の `word` は `.` または小文字英字
- 検索クエリには最大2つの `.` が含まれる
- `addWord` と `search` は最大10,000回呼ばれる

## ソースコード

```go
type TrieNode struct {
	children [26]*TrieNode
	word     bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewTrieNode()}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		index := c - 'a'
		if cur.children[index] == nil {
			cur.children[index] = NewTrieNode()
		}
		cur = cur.children[index]
	}
	cur.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {
	cur := root
	for i := j; i < len(word); i++ {
		c := word[i]
		if c == '.' {
			for _, child := range cur.children {
				if child != nil && this.dfs(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			index := c - 'a'
			if cur.children[index] == nil {
				return false
			}
			cur = cur.children[index]
		}
	}
	return cur.word
}
```

## アルゴリズムなど解説

### 基本戦略

Trie（プレフィックス木）+ DFS。通常のTrieに「ワイルドカード検索」を追加。`.` が来たら全ての子ノードを探索する。

### 通常の Trie との違い

```
通常のTrie:
Search("day")
→ 'd' → 'a' → 'y' と1本道で進む

この問題:
Search(".ay")
→ '.' は任意の文字
→ 全ての子を試す必要がある！
```

### AddWord の動作

```go
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		index := c - 'a'
		if cur.children[index] == nil {
			cur.children[index] = NewTrieNode()
		}
		cur = cur.children[index]
	}
	cur.word = true
}
```

これは標準的なTrieの挿入と同じ。

### Search の核心：DFS

```go
func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {
	cur := root
	for i := j; i < len(word); i++ {
		c := word[i]
		if c == '.' {
			// ワイルドカード: 全ての子を試す
			for _, child := range cur.children {
				if child != nil && this.dfs(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			// 通常の文字: 1本道で進む
			index := c - 'a'
			if cur.children[index] == nil {
				return false
			}
			cur = cur.children[index]
		}
	}
	return cur.word
}
```

### ワイルドカードの処理

```
格納: "day", "bay", "may"

Trie構造:
      root
     / | \
    b  d  m
    |  |  |
    a  a  a
    |  |  |
   [y][y][y]

Search(".ay"):

Step 1: c = '.'
        全ての子を試す: b, d, m

Step 2: 'b' から始まる探索
        dfs("ay", 1, node_b)
        → 'a' → 'y' → word=true ✓

1つでも true なら return true!
```

### 視覚的な理解

```
Search("b.."):

      root
       ↓
       b  ← c='b' マッチ
       ↓
       a  ← c='.' 全ての子を試す（'a'しかない）
       ↓
      [y] ← c='.' 全ての子を試す（'y'しかない）
       ↓
    word=true ✓

結果: true（"bay" がマッチ）
```

```
Search("say"):

      root
       ↓
       s  ← 存在しない！

結果: false
```

### なぜ DFS が必要か

```
ループだけでは実現できない:

Search("..y"):
root の全ての子を試す
├── 'b' の全ての子を試す
│   └── 'a' → 'y' チェック
├── 'd' の全ての子を試す
│   └── 'a' → 'y' チェック
└── 'm' の全ての子を試す
    └── 'a' → 'y' チェック

→ 分岐が発生するので再帰（DFS）が必要
```

### DFS のフロー

```go
// j: 現在の検索位置
// root: 現在のノード

dfs(word, 0, root)
│
├── c='.' の場合
│   └── for child in children:
│       └── dfs(word, i+1, child)  // 再帰
│
└── c='a'-'z' の場合
    └── cur = cur.children[index]  // 1本道で進む
```

### 計算量

| 操作 | 時間計算量 | 説明 |
|------|-----------|------|
| AddWord | O(m) | m = 単語長 |
| Search | O(m) ～ O(26^d × m) | d = `.`の数 |

```
最悪の場合（全て '.'）:
Search("...") で単語長3
→ 26 × 26 × 26 = 17,576 通りを探索

制約「最大2つの '.'」があるので:
→ 最大 26² = 676 通り × 残りの文字
→ 実用的な速度
```

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 空間計算量 | O(n × m) | n=単語数, m=平均単語長 |

### 最適化のアイデア

```go
// 長さごとに単語をグルーピング
type WordDictionary struct {
	root *TrieNode
	wordsByLen map[int][]string  // 長さ → 単語リスト
}

// '.' が多い場合、短い単語リストを線形探索する方が速いことも
```

### 別解：純粋な再帰版

```go
func (this *WordDictionary) Search(word string) bool {
	return this.searchHelper(word, 0, this.root)
}

func (this *WordDictionary) searchHelper(word string, idx int, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if idx == len(word) {
		return node.word
	}

	c := word[idx]
	if c == '.' {
		// 全ての子を試す
		for i := 0; i < 26; i++ {
			if this.searchHelper(word, idx+1, node.children[i]) {
				return true
			}
		}
		return false
	}
	// 特定の子に進む
	return this.searchHelper(word, idx+1, node.children[c-'a'])
}
```

### 元の解法との違い

```
元の解法:
- ループ + 部分的な再帰
- 通常文字はループで進む
- '.' の時だけ再帰
- 効率的

純粋な再帰版:
- 全て再帰
- コードがシンプル
- 関数呼び出しのオーバーヘッド

→ 元の解法の方が効率的
```

### 重要なポイント

```
1. 通常の文字
   → 1本道で進む（ループ）

2. ワイルドカード '.'
   → 全ての子を試す（DFS）
   → 1つでも true なら成功

3. 終了条件
   → 単語の最後まで到達したら word フラグをチェック

4. 計算量の考慮
   → '.' の数が多いと探索空間が爆発
   → 制約で最大2つに制限されている
```

### 関連問題

| 問題 | 共通点 |
|------|--------|
| Implement Trie | 基本的なTrie実装 |
| Word Search II | Trie + グリッドDFS |
| Wildcard Matching | パターンマッチング |

Trie + DFS の組み合わせパターン！ワイルドカード対応の典型問題。
