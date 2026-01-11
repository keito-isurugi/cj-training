# Linked List Cycle (Easy)

## 問題内容

リンクリストの先頭 `head` が与えられたとき、リストにサイクル（循環）が存在すれば `true` を、存在しなければ `false` を返す。

### 例

```
Input: head = [1,2,3,4], index = 1
Output: true
Explanation: 末尾がindex=1のノード（値2）に接続してサイクルを形成

Input: head = [1,2], index = -1
Output: false
Explanation: サイクルなし（末尾がnilを指す）
```

### 制約

- `1 <= リストの長さ <= 1000`
- `-1000 <= Node.val <= 1000`

## ソースコード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

## アルゴリズムなど解説

### 基本戦略

Floyd's Cycle Detection Algorithm（通称: 亀とウサギ）を使用。2つのポインタを異なる速度で進め、サイクルがあれば必ず出会う。

### 動作の仕組み

1. **2つのポインタを初期化**
   ```go
   slow := head  // 亀（1歩ずつ）
   fast := head  // ウサギ（2歩ずつ）
   ```

2. **異なる速度で進む**
   ```go
   for fast != nil && fast.Next != nil {
       slow = slow.Next       // 1歩
       fast = fast.Next.Next  // 2歩
   ```

3. **出会ったらサイクルあり**
   ```go
       if slow == fast {
           return true
       }
   }
   ```

4. **fastがnilに到達したらサイクルなし**
   ```go
   return false
   ```

### 視覚的な理解（サイクルあり）

```
index = 1 のサイクル:

[1] → [2] → [3] → [4]
       ↑___________↓

初期: slow=1, fast=1
```

#### ステップ1

```
slow: 1 → 2（1歩）
fast: 1 → 2 → 3（2歩）

[1] → [2] → [3] → [4]
       s     f    ↓
       ↑__________↓
```

#### ステップ2

```
slow: 2 → 3（1歩）
fast: 3 → 4 → 2（2歩、サイクルを通過）

[1] → [2] → [3] → [4]
       f     s    ↓
       ↑__________↓
```

#### ステップ3

```
slow: 3 → 4（1歩）
fast: 2 → 3 → 4（2歩）

[1] → [2] → [3] → [4]
                  s,f ← 出会った！
       ↑__________↓

slow == fast → return true
```

### 視覚的な理解（サイクルなし）

```
[1] → [2] → nil

初期: slow=1, fast=1

ステップ1:
slow: 1 → 2
fast: 1 → 2 → nil

fast.Next == nil → ループ終了
return false
```

### なぜ必ず出会うのか？

```
サイクル内での相対速度:
- slowは1歩/回
- fastは2歩/回
- 相対速度 = 2 - 1 = 1歩/回

サイクルに入った後:
- fastはslowに毎回1歩ずつ近づく
- サイクルの長さに関係なく、必ず追いつく

例: サイクル長が5の場合
    距離4 → 3 → 2 → 1 → 0（出会う）
```

### なぜ `fast != nil && fast.Next != nil` か？

```go
for fast != nil && fast.Next != nil {
    fast = fast.Next.Next  // これが安全に実行できる
}
```

```
fast.Next.Next を実行するには:
1. fast != nil（fastが存在）
2. fast.Next != nil（fast.Nextが存在）

両方チェックしないとnilポインタエラー
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 最悪でもリスト長+サイクル長の回数で出会う |
| 空間計算量 | O(1) | ポインタ2つのみ使用 |

### 別解：HashSetを使う方法

```go
func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]bool)
    curr := head

    for curr != nil {
        if seen[curr] {
            return true  // 既に訪問済み = サイクル
        }
        seen[curr] = true
        curr = curr.Next
    }
    return false
}
```

| 比較 | Floyd's | HashSet |
|------|---------|---------|
| 時間計算量 | O(n) | O(n) |
| 空間計算量 | O(1) | O(n) |

Floyd'sの方がメモリ効率が良い。
