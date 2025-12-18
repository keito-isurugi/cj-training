# Combination Sum (Medium)

## 問題内容

異なる整数の配列 `nums` とターゲット整数 `target` が与えられる。`nums` から選んだ数の和が `target` になる、すべてのユニークな組み合わせのリストを返す。

同じ数を無制限に使用可能。選んだ数の出現頻度が同じなら同じ組み合わせ、そうでなければ異なる組み合わせ。

### 例

```
Input: nums = [2,5,6,9], target = 9

Output: [[2,2,5],[9]]

Explanation:
2 + 2 + 5 = 9 （2を2回、5を1回）
9 = 9 （9を1回）
```

```
Input: nums = [3,4,5], target = 16

Output: [[3,3,3,3,4],[3,3,5,5],[4,4,4,4],[3,4,4,5]]
```

```
Input: nums = [3], target = 5

Output: []
```

### 制約

- `nums` の全要素は異なる
- `1 <= nums.length <= 20`
- `2 <= nums[i] <= 30`
- `2 <= target <= 30`

## ソースコード

```go
func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    var dfs func(int, []int, int)
    dfs = func(i int, cur []int, total int) {
        if total == target {
            temp := make([]int, len(cur))
            copy(temp, cur)
            res = append(res, temp)
            return
        }

        for j := i; j < len(nums); j++ {
            if total + nums[j] > target {
                return
            }
            cur = append(cur, nums[j])
            dfs(j, cur, total + nums[j])
            cur = cur[:len(cur)-1]
        }
    }

    dfs(0, []int{}, 0)
    return res
}
```

## アルゴリズムなど解説

### 基本戦略

バックトラッキング（深さ優先探索 + 戻り）で全ての組み合わせを探索。重複を避けるため、現在のインデックス以降の要素のみを使用。

### バックトラッキングとは

```
探索木を深さ優先で進み、条件を満たさなくなったら「戻る」

            []
         /   |   \
       [2]  [5]  [6]  [9]
      / | \   |
   [2,2] [2,5] ...
    /
[2,2,2] → targetを超えたら戻る
   ↑
 backtrack!
```

### なぜソートするか

```go
sort.Ints(nums)
```

```
ソートすることで:

1. 早期終了が可能
   nums = [2, 5, 6, 9], target = 9

   total=8 のとき:
   - nums[j]=5 → 8+5=13 > 9 → これ以降は全部超える！
   - すぐにreturnできる（枝刈り）

2. 重複回避が簡単
   インデックス i 以降のみ使うことで重複を防ぐ
```

### 核心的なポイント

```go
for j := i; j < len(nums); j++ {
    // j を i から開始 → 自分と同じ or 後ろの要素のみ使用
```

```
なぜ j = i から？

nums = [2, 5, 6]

j = 0 からだと:
[2, 5] と [5, 2] が両方生成されてしまう

j = i からだと:
[2] を選んだ後は [2, 5, 6] から選ぶ
[5] を選んだ後は [5, 6] から選ぶ
→ [5, 2] は生成されない！
```

### バックトラッキングの3ステップ

```go
// 1. 選択する
cur = append(cur, nums[j])

// 2. 再帰で探索
dfs(j, cur, total + nums[j])

// 3. 選択を取り消す（バックトラック）
cur = cur[:len(cur)-1]
```

```
例: nums = [2, 5], target = 4

dfs(0, [], 0)
├── 選択: cur = [2]
│   └── dfs(0, [2], 2)
│       ├── 選択: cur = [2, 2]
│       │   └── dfs(0, [2, 2], 4)
│       │       └── total == target → 結果に追加！
│       │   取消: cur = [2]
│       ├── 選択: cur = [2, 5]
│       │   └── 2 + 5 = 7 > 4 → return
│       取消: cur = []
├── 選択: cur = [5]
│   └── dfs(1, [5], 5)
│       └── 5 > 4 → return
│   取消: cur = []
```

### なぜコピーが必要か

```go
if total == target {
    temp := make([]int, len(cur))
    copy(temp, cur)
    res = append(res, temp)
    return
}
```

```
Goのスライスは参照型

❌ コピーしないと:
res = append(res, cur)
// curは後で変更される → resの中身も変わってしまう！

✅ コピーすると:
temp := make([]int, len(cur))
copy(temp, cur)
res = append(res, temp)
// tempは独立したコピー → 安全
```

### 視覚的な理解

```
nums = [2, 5, 6, 9], target = 9

探索木:

                     []
           /       |      \      \
         [2]      [5]    [6]    [9] ← 9==9 ✓
        / | \      |      |
     [2,2] [2,5] [2,6] [5,5]  [6,6]
      |      ↑           ↑      ↑
   [2,2,2] 2+5=7   5+5>9  6+6>9
    |    [2,2,5] ← 2+2+5=9 ✓
 [2,2,2,2]   ↑
    ↑     9==9 ✓
  8+2>9

結果: [[2,2,5], [9]]
```

### 早期終了（枝刈り）

```go
if total + nums[j] > target {
    return
}
```

```
ソート済みなので、nums[j]を超えたら
nums[j+1], nums[j+2], ... も全部超える

例: total=7, nums=[2,5,6,9], target=9
j=1: 7+5=12 > 9 → return
     (6, 9 もチェック不要！)
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n^(t/m)) | n=nums長, t=target, m=最小値 |
| 空間計算量 | O(t/m) | 再帰の深さ |

### 別解：インデックスなしで重複許可

```go
// 同じ要素を何度も使える別の書き方
func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}

    var dfs func(start int, cur []int, remain int)
    dfs = func(start int, cur []int, remain int) {
        if remain == 0 {
            temp := make([]int, len(cur))
            copy(temp, cur)
            res = append(res, temp)
            return
        }
        if remain < 0 {
            return
        }

        for i := start; i < len(nums); i++ {
            cur = append(cur, nums[i])
            dfs(i, cur, remain-nums[i])  // i から（同じ要素を再利用可能）
            cur = cur[:len(cur)-1]
        }
    }

    dfs(0, []int{}, target)
    return res
}
```

### バックトラッキングのテンプレート

```go
func backtrack(candidates, path, result) {
    // 1. 終了条件
    if 条件を満たす {
        result に path を追加
        return
    }

    // 2. 選択肢をループ
    for 選択肢 in candidates {
        // 3. 選択する
        path に追加

        // 4. 再帰
        backtrack(残りの候補, path, result)

        // 5. 選択を取り消す
        path から削除
    }
}
```

### 関連問題

| 問題 | 違い |
|------|------|
| Combination Sum II | 各要素は1回のみ使用 |
| Combination Sum III | k個の数で合計n |
| Subsets | 和ではなく全部分集合 |
| Permutations | 順列（順序が重要） |

バックトラッキングは組み合わせ・順列問題の基本パターン！
