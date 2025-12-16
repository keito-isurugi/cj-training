# Valid Parentheses (Easy)

## 問題内容

文字列 `s` が `'('`, `')'`, `'{'`, `'}'`, `'['`, `']'` のみで構成されているとき、以下の条件をすべて満たす場合に `true` を返す。

- すべての開き括弧は同じ種類の閉じ括弧で閉じられている
- 開き括弧は正しい順序で閉じられている
- すべての閉じ括弧には対応する開き括弧がある

### 例

```
Input: s = "[]"
Output: true

Input: s = "([{}])"
Output: true

Input: s = "[(])"
Output: false
Explanation: 括弧が正しい順序で閉じられていない
```

### 制約

- `1 <= s.length <= 1000`

## ソースコード

```go
func isValid(s string) bool {
    stack := linkedliststack.New()
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if !stack.Empty() {
                top, ok := stack.Pop()
                if ok && top.(rune) != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack.Push(c)
        }
    }

    return stack.Empty()
}
```

## アルゴリズムなど解説

### 基本戦略

Stackを使って、開き括弧を積み上げ、閉じ括弧が来たら対応する開き括弧をpopする。

### 動作の仕組み

1. **マッピングを定義**
   ```go
   closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}
   ```
   - 閉じ括弧 → 対応する開き括弧のマッピング

2. **各文字を処理**
   ```go
   for _, c := range s {
       if open, exists := closeToOpen[c]; exists {  // 閉じ括弧の場合
           if !stack.Empty() {
               top, ok := stack.Pop()
               if ok && top.(rune) != open {
                   return false  // 不一致
               }
           } else {
               return false  // 空スタック
           }
       } else {  // 開き括弧の場合
           stack.Push(c)  // スタックに追加
       }
   }
   ```

3. **最終判定**
   ```go
   return stack.Empty()
   ```
   - スタックが空 → すべて正しくペアになった → `true`
   - スタックに残りあり → 閉じられていない括弧がある → `false`

### 具体例

#### 例1: `s = "([{}])"`

```
文字   アクション         スタック
'('    push              ['(']
'['    push              ['(', '[']
'{'    push              ['(', '[', '{']
'}'    pop ({}ペア)       ['(', '[']
']'    pop ([]ペア)       ['(']
')'    pop (()ペア)       []

スタックが空 → true
```

#### 例2: `s = "[(])"`

```
文字   アクション         スタック
'['    push              ['[']
'('    push              ['[', '(']
']'    ']'の対応は'['だが
       スタックトップは'(' → 不一致！

return false
```

### なぜStackが適切か

```
括弧の特性:
- 最後に開いた括弧を最初に閉じる必要がある（LIFO）
- これはStackの特性と完全に一致

例: "([{}])"

開く順: ( → [ → {
閉じる順: } → ] → )  ← 逆順！

Stack: push, push, push, pop, pop, pop
```

### 計算量

| 計算量 | 値 | 説明 |
|--------|-----|------|
| 時間計算量 | O(n) | 文字列を1回スキャン |
| 空間計算量 | O(n) | 最悪の場合、全て開き括弧でスタックに格納 |
