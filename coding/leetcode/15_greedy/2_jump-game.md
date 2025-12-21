# Jump Game (Medium)

## 問題内容

整数配列 `nums` が与えられる。各要素 `nums[i]` は、その位置からの**最大ジャンプ長**を示す。

インデックス `0` から開始して、**最後のインデックスに到達できるかどうか**を返す。

### 例

```
Input: nums = [1,2,0,1,0]
Output: true
```
説明: インデックス 0→1→3→4 とジャンプして到達可能。

```
Input: nums = [1,2,1,0,1]
Output: false
```

### 制約

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`

## ソースコード

```go
func canJump(nums []int) bool {
    goal := len(nums) - 1

    for i := len(nums) - 2; i >= 0; i-- {
        if i+nums[i] >= goal {
            goal = i
        }
    }
    return goal == 0
}
```

## アルゴリズムなど解説

### 基本戦略

**後ろから前へ**考える：
- 「最後のインデックスに到達できる最も左の位置」を追跡
- その位置に到達できれば、ゴールに到達可能

### 核心となる洞察

```
位置 i から goal に到達できるなら、goal を i に更新する
```

`i + nums[i] >= goal` であれば、位置 `i` から現在の `goal` に到達可能。

### 動作の仕組み

1. **ゴールの初期化**
   ```go
   goal := len(nums) - 1
   ```
   - 最初のゴールは配列の最後のインデックス

2. **後ろから前へ走査**
   ```go
   for i := len(nums) - 2; i >= 0; i-- {
       if i+nums[i] >= goal {
           goal = i
       }
   }
   ```
   - 各位置で、現在の `goal` に到達できるかチェック
   - 到達可能なら、その位置が新しい `goal` になる

3. **結果の判定**
   ```go
   return goal == 0
   ```
   - `goal` が `0` に到達していれば、スタートから最後まで到達可能

### 具体例

```
nums = [1, 2, 0, 1, 0]
       [0] [1] [2] [3] [4]

初期: goal = 4

i=3: 3 + nums[3] = 3 + 1 = 4 >= 4 ✓ → goal = 3
i=2: 2 + nums[2] = 2 + 0 = 2 < 3  ✗
i=1: 1 + nums[1] = 1 + 2 = 3 >= 3 ✓ → goal = 1
i=0: 0 + nums[0] = 0 + 1 = 1 >= 1 ✓ → goal = 0

goal == 0 → true
```

```
nums = [1, 2, 1, 0, 1]
       [0] [1] [2] [3] [4]

初期: goal = 4

i=3: 3 + nums[3] = 3 + 0 = 3 < 4  ✗
i=2: 2 + nums[2] = 2 + 1 = 3 < 4  ✗
i=1: 1 + nums[1] = 1 + 2 = 3 < 4  ✗
i=0: 0 + nums[0] = 0 + 1 = 1 < 4  ✗

goal == 4 != 0 → false
```

### なぜ後ろからなのか

前からだと：
- すべての到達可能な位置を追跡する必要がある
- 複雑になりやすい

後ろからだと：
- 「ゴールに到達できる最も近い位置」だけを追跡
- シンプルで効率的

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 配列を1回走査 |
| 空間計算量 | O(1) | 変数1つのみ使用 |

### 別解：前から走査（最大到達距離）

```go
func canJump(nums []int) bool {
    maxReach := 0

    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            return false
        }
        if i+nums[i] > maxReach {
            maxReach = i + nums[i]
        }
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    return true
}
```

**アイデア**: 現在到達可能な最大インデックスを追跡

### 別解：DP（Bottom-Up）

```go
func canJump(nums []int) bool {
    n := len(nums)
    dp := make([]bool, n)
    dp[n-1] = true

    for i := n - 2; i >= 0; i-- {
        end := min(n, i+nums[i]+1)
        for j := i + 1; j < end; j++ {
            if dp[j] {
                dp[i] = true
                break
            }
        }
    }
    return dp[0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

**計算量**: 時間 O(n²)、空間 O(n)

### Greedy パターン

この問題は **ゴール追跡型 Greedy** パターン：
- 目標に到達できる最も近い位置を追跡
- 局所最適解（各位置でゴールに到達可能か）を積み重ねる
- 最終的に全体最適解を得る

### 視覚化

```
[1, 2, 0, 1, 0]
 ↓  ↓     ↓  ↓
 0→1  →  3→4

パス: 0 → 1 → 3 → 4 (到達可能)
```

### 関連問題との比較

| 問題 | パターン |
|------|---------|
| **Jump Game** | ゴール追跡 Greedy |
| Jump Game II | 最小ジャンプ数（BFS/Greedy） |
| Minimum Jumps to Reach Home | BFS |
| Frog Jump | DP |
