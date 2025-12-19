# Implement Trie (Prefix Tree) (Medium)

## 問題内容

プレフィックス木（Trie）は、文字列の集合を効率的に格納・検索するための木構造データ。オートコンプリートやスペルチェッカーなどで使用される。

`PrefixTree` クラスを実装する:
- `PrefixTree()`: プレフィックス木オブジェクトを初期化
- `void insert(String word)`: 文字列 `word` を挿入
- `boolean search(String word)`: `word` が存在すれば `true`
- `boolean startsWith(String prefix)`: `prefix` で始まる単語が存在すれば `true`

### 例

```
Input:
["Trie", "insert", "dog", "search", "dog", "search", "do", "startsWith", "do", "insert", "do", "search", "do"]

Output:
[null, null, true, false, true, null, true]

Explanation:
PrefixTree prefixTree = new PrefixTree();
prefixTree.insert("dog");
prefixTree.search("dog");    // return true
prefixTree.search("do");     // return false（"do"は挿入されていない）
prefixTree.startsWith("do"); // return true（"dog"が"do"で始まる）
prefixTree.insert("do");
prefixTree.search("do");     // return true
```

### 制約

- `1 <= word.length, prefix.length <= 1000`
- `word` と `prefix` は小文字の英字のみ

## ソースコード

```go
type TrieNode struct {
	children map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			cur.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return cur.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return true
}
```

## アルゴリズムなど解説

### Trie（トライ）とは

```
単語 "dog", "do", "day" を格納した Trie:

        root
       /    \
      d      ...
     /
    [o]  ← "do" の終端（endOfWord = true）
   /  \
  g    ...
  ↓
[g]  ← "dog" の終端（endOfWord = true）

※ 同時に "day" も格納:
        root
       /
      d
     / \
    o   a
   /     \
  g       y
  ↓       ↓
 [g]     [y]
```

### なぜ Trie を使うか

```
配列/ハッシュマップとの比較:

検索:
- 配列: O(n × m)  n=単語数, m=単語長
- ハッシュ: O(m)  平均
- Trie: O(m)     確定

プレフィックス検索:
- 配列: O(n × m)  全単語をチェック
- ハッシュ: 困難（全キーをチェック）
- Trie: O(p)      p=プレフィックス長 ← 高速！

オートコンプリートに最適！
```

### TrieNode の構造

```go
type TrieNode struct {
	children map[rune]*TrieNode  // 子ノードへのマップ
	endOfWord bool               // 単語の終端か？
}
```

```
各ノードは:
- children: 次の文字へのポインタ群
- endOfWord: このノードで単語が終わるか

例: "do" と "dog" を格納
root → 'd' → 'o' → 'g'
              ↑      ↑
         endOfWord  endOfWord
           =true     =true
```

### Insert の動作

```go
func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			// 子が存在しなければ作成
			cur.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true  // 単語の終端をマーク
}
```

### 視覚的な理解（Insert）

```
Insert("dog"):

Step 1: cur = root, c = 'd'
        root.children['d'] = nil → 新規作成
        root
         ↓
        [d]
        cur = [d]

Step 2: cur = [d], c = 'o'
        [d].children['o'] = nil → 新規作成
        root
         ↓
        [d]
         ↓
        [o]
        cur = [o]

Step 3: cur = [o], c = 'g'
        [o].children['g'] = nil → 新規作成
        root
         ↓
        [d]
         ↓
        [o]
         ↓
        [g] ← endOfWord = true

Insert("do"):
        root
         ↓
        [d]
         ↓
        [o] ← endOfWord = true（ここに追加！）
         ↓
        [g] ← endOfWord = true
```

### Search vs StartsWith

```go
// Search: 完全一致を検索
func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return cur.endOfWord  // 終端かどうかをチェック！
}

// StartsWith: プレフィックスを検索
func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return true  // パスが存在すればOK
}
```

```
違い:

Trie: "dog" のみ格納

Search("do"):
root → 'd' → 'o'
             ↑
        endOfWord = false
→ return false

StartsWith("do"):
root → 'd' → 'o'
             ↑
        パスは存在する！
→ return true
```

### 計算量

| 操作 | 時間計算量 | 説明 |
|------|-----------|------|
| Insert | O(m) | m = 単語長 |
| Search | O(m) | m = 単語長 |
| StartsWith | O(p) | p = プレフィックス長 |

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 空間計算量 | O(n × m) | n=単語数, m=平均単語長 |

### 別解：配列ベース（26文字固定）

```go
type TrieNode struct {
	children [26]*TrieNode  // 'a'〜'z' の26文字
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{root: &TrieNode{}}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word {
		idx := c - 'a'  // 'a'=0, 'b'=1, ..., 'z'=25
		if cur.children[idx] == nil {
			cur.children[idx] = &TrieNode{}
		}
		cur = cur.children[idx]
	}
	cur.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		idx := c - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}
```

### map vs 配列の比較

| 方法 | メモリ | アクセス速度 | 用途 |
|------|--------|------------|------|
| map | 疎な場合に効率的 | O(1)平均 | 文字種が多い/不定 |
| 配列[26] | 固定26ポインタ | O(1)確定 | 小文字英字のみ |

### Trie の応用

```
1. オートコンプリート
   - 入力 "do" → "dog", "door", "done" を提案

2. スペルチェッカー
   - 入力 "dgo" → "dog" の候補を提示

3. IPルーティング
   - 最長プレフィックスマッチング

4. Word Search II
   - 複数単語を同時に検索
```

### 重要なポイント

```
1. TrieNode の構造
   - children: 次の文字へのマップ/配列
   - endOfWord: 単語の終端フラグ

2. Search vs StartsWith
   - Search: endOfWord をチェック
   - StartsWith: パスの存在のみチェック

3. 共通プレフィックスの共有
   - "dog" と "do" は 'd' → 'o' を共有
   - メモリ効率が良い

4. 操作は全て O(m)
   - 単語/プレフィックスの長さに比例
   - 格納された単語数に依存しない！
```

### 関連問題

| 問題 | Trie の活用 |
|------|------------|
| Word Search II | Trie + DFS |
| Design Add and Search Words | ワイルドカード対応 |
| Replace Words | プレフィックス置換 |
| Longest Word in Dictionary | 辞書順最長単語 |

Trie は文字列検索の基本データ構造。プレフィックス検索が必要な問題で活躍！
