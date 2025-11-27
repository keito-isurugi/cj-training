# Valid Anagram

## Question
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

## Solution
```.go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }

    // 初期値0の要素を26個持つ固定配列
    // 小文字のアルファベットa~zを表す
    count := [26]int{}
    for i := 0; i < len(s); i++ {
        // 自分の文字列の箇所の要素を+1
        count[s[i] - 'a']++
        // 自分の文字列の箇所の要素を-1
        count[t[i] - 'a']--
    }

    // 配列の要素が全て0ならアナグラムとみなす
    for _, v := range count {
        if v != 0 {
            return false
        }
    }

    return true
}
```
