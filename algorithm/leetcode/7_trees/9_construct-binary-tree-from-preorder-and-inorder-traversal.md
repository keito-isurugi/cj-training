# Construct Binary Tree from Preorder and Inorder Traversal (Medium)

## 問題内容

前順走査（preorder）と中順走査（inorder）の結果から、元の二分木を再構築する。

### 例

```
Input: preorder = [1,2,3,4], inorder = [2,1,3,4]

Output:
    1
   / \
  2   3
       \
        4
```

```
Input: preorder = [1], inorder = [1]
Output: [1]
```

### 制約

- `1 <= inorder.length <= 1000`
- `inorder.length == preorder.length`
- `-1000 <= preorder[i], inorder[i] <= 1000`
- すべての値はユニーク

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    head := &TreeNode{}
    curr := head
    i, j, n := 0, 0, len(preorder)

    for i < n && j < n {
        curr.Right = &TreeNode{Val: preorder[i], Right: curr.Right}
        curr = curr.Right
        i++
        for i < n && curr.Val != inorder[j] {
            curr.Left = &TreeNode{Val: preorder[i], Right: curr}
            curr = curr.Left
            i++
        }
        j++
        for curr.Right != nil && j < n && curr.Right.Val == inorder[j] {
            prev := curr.Right
            curr.Right = nil
            curr = prev
            j++
        }
    }
    return head.Right
}
```

## アルゴリズムなど解説

### 走査の基本知識

```
       1
      / \
     2   3
        / \
       4   5

前順（Pre-order）: Root → Left → Right
= [1, 2, 3, 4, 5]

中順（In-order）: Left → Root → Right
= [2, 1, 4, 3, 5]
```

### 核心的なアイデア

```
Preorder の性質:
- 最初の要素 = ルート
- 次に左部分木、その後右部分木

Inorder の性質:
- ルートの左側 = 左部分木のノード
- ルートの右側 = 右部分木のノード

この2つを組み合わせて木を再構築！
```

### 例での理解

```
preorder = [1, 2, 3, 4]
inorder  = [2, 1, 3, 4]

Step 1: preorder[0] = 1 がルート

Step 2: inorder で 1 の位置を探す
inorder = [2, | 1, | 3, 4]
          左部分木  右部分木

Step 3: 左部分木 = [2]、右部分木 = [3, 4]

Step 4: 再帰的に構築
      1
     / \
    2   3
         \
          4
```

### このコードの特徴：反復的アプローチ

このコードは**スタックを使わない反復的な方法**で、`Right` ポインタを一時的に親への参照として使用している。

```go
curr.Right = &TreeNode{Val: preorder[i], Right: curr.Right}
// Right を一時的に「親へのポインタ」として使用
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 各ノードを1回処理 |
| 空間計算量 | O(n) | 構築するノード分 |

### 別解：再帰版（より理解しやすい）

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    // inorderの値→インデックスのマップを作成
    inorderMap := make(map[int]int)
    for i, val := range inorder {
        inorderMap[val] = i
    }

    var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
    build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
        if preStart > preEnd {
            return nil
        }

        // preorderの最初の要素がルート
        rootVal := preorder[preStart]
        root := &TreeNode{Val: rootVal}

        // inorderでルートの位置を探す
        rootIndex := inorderMap[rootVal]
        leftSize := rootIndex - inStart  // 左部分木のサイズ

        // 再帰的に構築
        root.Left = build(preStart+1, preStart+leftSize, inStart, rootIndex-1)
        root.Right = build(preStart+leftSize+1, preEnd, rootIndex+1, inEnd)

        return root
    }

    return build(0, len(preorder)-1, 0, len(inorder)-1)
}
```

### 再帰版の動作

```
preorder = [1, 2, 3, 4]
inorder  = [2, 1, 3, 4]

build(0, 3, 0, 3):
  root = 1
  rootIndex = 1 (inorderでの位置)
  leftSize = 1

  root.Left = build(1, 1, 0, 0)
    → root = 2, 子なし

  root.Right = build(2, 3, 2, 3)
    → root = 3
       root.Right = build(3, 3, 3, 3)
         → root = 4, 子なし

結果:
    1
   / \
  2   3
       \
        4
```

### 視覚的な理解（再帰版）

```
preorder = [1, 2, 3, 4]
           ↑
          ルート

inorder  = [2, 1, 3, 4]
           左 ↑ 右部分木
             ルート

━━━ Step 1: ルート = 1 ━━━
preorder: [1 | 2 | 3, 4]
              左   右

inorder:  [2 | 1 | 3, 4]
           左     右

━━━ Step 2: 左部分木 ━━━
preorder: [2]
inorder:  [2]
→ ノード 2

━━━ Step 3: 右部分木 ━━━
preorder: [3, 4]
inorder:  [3, 4]
→ ルート 3、右の子 4
```

### なぜ2つの配列から木を再構築できるか

```
Preorder だけだと:
[1, 2, 3] → どの形？

    1        1        1
   /        / \        \
  2        2   3        2
 /                       \
3                         3

Inorder を加えると:
preorder = [1, 2, 3]
inorder  = [2, 1, 3]

inorder で 1 の左に 2、右に 3
→ 一意に決定！

    1
   / \
  2   3
```

### 解法の比較

| 解法 | 特徴 | おすすめ度 |
|------|------|----------|
| 反復版（提供コード） | 空間効率が良いが複雑 | 上級者向け |
| 再帰版 + HashMap | 直感的で理解しやすい | **面接推奨** |
| 再帰版（スライス分割） | シンプルだがO(n²) | 小さい入力向け |

面接では**再帰版 + HashMap**が説明しやすく、効率も良いのでおすすめ！
