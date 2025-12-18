# Subtree of Another Tree (Easy)

## 問題内容

2つの二分木 `root` と `subRoot` が与えられたとき、`root` のいずれかの部分木が `subRoot` と同じ構造・同じ値であれば `true` を、そうでなければ `false` を返す。

### 例

```
Input:
root:           1           subRoot:    2
               / \                     / \
              2   3                   4   5
             / \
            4   5

Output: true
Explanation: rootの左部分木がsubRootと一致
```

```
Input:
root:           1           subRoot:    2
               / \                     / \
              2   3                   4   5
             / \
            4   5
           /
          6

Output: false
Explanation: ノード4の下に6があるため、構造が異なる
```

### 制約

- `1 <= ノード数 <= 100`
- `-100 <= Node.val <= 100`

## ソースコード

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func serialize(root *TreeNode) string {
    if root == nil {
        return "$#"
    }
    return "$" + strconv.Itoa(root.Val) + serialize(root.Left) + serialize(root.Right)
}

func zFunction(s string) []int {
    n := len(s)
    z := make([]int, n)
    l, r := 0, 0

    for i := 1; i < n; i++ {
        if i <= r {
            z[i] = min(r-i+1, z[i-l])
        }
        for i+z[i] < n && s[z[i]] == s[i+z[i]] {
            z[i]++
        }
        if i+z[i]-1 > r {
            l = i
            r = i + z[i] - 1
        }
    }
    return z
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    serializedRoot := serialize(root)
    serializedSubRoot := serialize(subRoot)
    combined := serializedSubRoot + "|" + serializedRoot

    zValues := zFunction(combined)
    subLen := len(serializedSubRoot)

    for i := subLen + 1; i < len(combined); i++ {
        if zValues[i] == subLen {
            return true
        }
    }
    return false
}
```

## アルゴリズムなど解説

### 基本戦略

木をシリアライズ（文字列化）し、Z-Algorithmを使って部分文字列マッチングを行う。

### なぜこのアプローチ？

```
部分木の判定 = 木構造の比較
↓
木を文字列に変換
↓
部分文字列の検索問題に帰着
↓
Z-Algorithm で効率的に検索
```

### Step 1: シリアライズ（木→文字列）

```go
func serialize(root *TreeNode) string {
    if root == nil {
        return "$#"  // nilを特殊文字で表現
    }
    return "$" + strconv.Itoa(root.Val) + serialize(root.Left) + serialize(root.Right)
}
```

```
例:
        2
       / \
      4   5

serialize(2) = "$2" + serialize(4) + serialize(5)
             = "$2" + "$4$#$#" + "$5$#$#"
             = "$2$4$#$#$5$#$#"

各要素の意味:
$2   → ノード2
$4   → ノード4
$#   → 左の子がnil
$#   → 右の子がnil
$5   → ノード5
$#   → 左の子がnil
$#   → 右の子がnil
```

### なぜ区切り文字が必要か

```
区切り文字なしだと:
- ノード12 → "12"
- ノード1と2 → "12"
→ 区別できない！

$を使うと:
- ノード12 → "$12"
- ノード1と2 → "$1$2"
→ 明確に区別できる
```

### Step 2: Z-Algorithm（文字列マッチング）

Z-Algorithmは、文字列の各位置から始まる最長の接頭辞一致長を求める。

```
combined = subRoot + "|" + root
         = "pattern|text"

Z-Algorithmで:
- z[i] = patternと一致する長さ
- z[i] == len(pattern) → 完全一致！
```

### 具体例

```
root:       1           subRoot:    2
           / \                     / \
          2   3                   4   5
         / \
        4   5

serializedRoot    = "$1$2$4$#$#$5$#$#$3$#$#"
serializedSubRoot = "$2$4$#$#$5$#$#"

combined = "$2$4$#$#$5$#$#|$1$2$4$#$#$5$#$#$3$#$#"
            ↑ pattern    ↑  ↑ text
            (subLen=14)  区切り

Z-Algorithmで各位置のz値を計算:
位置16 (rootの$2から): z[16] = 14 = subLen
→ 完全一致！return true
```

### Z-Algorithmの動作

```go
func zFunction(s string) []int {
    // z[i] = s[i:]とs[0:]の最長共通接頭辞の長さ
}
```

```
例: s = "aabxaab"
        0123456

z[0] = 0 (定義上)
z[1] = 1 ("ab..." vs "aabx..." → "a"が一致)
z[2] = 0 ("bx..." vs "aabx..." → 不一致)
z[3] = 0 ("x..." vs "aabx..." → 不一致)
z[4] = 3 ("aab" vs "aabx..." → "aab"が一致)
z[5] = 1 ("ab" vs "aabx..." → "a"が一致)
z[6] = 0 ("b" vs "aabx..." → 不一致)
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n + m) | シリアライズO(n) + Z-Algorithm O(n+m) |
| 空間計算量 | O(n + m) | シリアライズされた文字列 |

n = rootのノード数, m = subRootのノード数

### 別解：再帰版（シンプル）

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    if isSameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

```
動作:
1. rootの各ノードを訪問
2. そのノードを起点にsubRootと同じか判定
3. 一致したらtrue、なければ次のノードへ
```

### 解法の比較

| 解法 | 時間計算量 | 空間計算量 | 特徴 |
|------|-----------|-----------|------|
| Z-Algorithm | O(n + m) | O(n + m) | 高速だが複雑 |
| 再帰（シンプル） | O(n × m) | O(h) | シンプルで実用的 |
| ハッシュ | O(n + m) | O(n + m) | 中間的 |

面接では**再帰版**がシンプルで説明しやすい。Z-Algorithmは最適化が求められた場合に有効。
