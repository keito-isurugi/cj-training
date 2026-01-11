package arrayhashmap

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	hasDuplicate([]int{1,2,3,4,5})
}

// Contains Duplicate
func hasDuplicate(nums []int) bool {
	// numsの要素数が0以下なら return false
	if len(nums) <= 0 {
		return false
	}
	
    // 空のmapを作るindex = int, value = bool
    m := make(map[int]bool)
    
    // numsをループで回す
    for _, n := range nums {
    	// mapの存在確認
    	if m[n] {
    		return true
    	}
    	// 存在しなければtrueを入れる
    	m[n] = true
    }
    
    return false
}

// Valid Anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// 初期値0の要素を26個持つ固定配列
	// アルファベットa~zを表す
	count := [26]int{}
	
	// 文字列分ループ
	for i := 0; i < len(s); i++ {
		// sのi番目の文字の箇所を+1(-aすることで26に収まる)
		count[s[i] - 'a']++
		// tのi番目の文字の箇所を-1
		count[t[i] - 'a']--
	}
	
	// 固定配列分繰り返してすべて0かをチェック
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	
	return true
}


// Two Sum
func twoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return nil
	}

	// 空のmapを作成 key = int(targetとnums[i]の差分), value = int(nums[i])
	nm := make(map[int]int)
	for i, n := range nums {
		// targetとnums[i]の差分と等しい数値が存在するかチェック、jにはindexが入る
		if j, ok := nm[target - n]; ok {
			return []int{j, i} // iより先に追加されているjを先頭にして返す
		}
		
		// 存在しなければkeyに数値、valueにiを入れる
		nm[n] = i
	}
	
	return nil
}

// Group Anagrams
func groupAnagrams(strs []string) [][]string {
	// mapのキーは a〜z の文字出現回数を格納した配列（アナグラム判定用キー）
	// 値はそのキーに属する単語のスライス
	// "eat"  -> [1,0,0,0,1,0, ... ,1] 
	// "tea"  -> 同じ配列
	res := make(map[[26]int][]string)
	
	for _, str := range strs {
		// 単語のマッピング用配列
		var count [26]int
		for _, c := range str {
			// 単語のi番目の文字列の箇所に+1(-aすることで26におさめる)
			count[c - 'a']++
		}
		
		// 単語の文字カウント（[26]int）をキーにしてアナグラムをまとめる
		// 例: [1,0,0,0,1,0,...,1] -> ["eat", "ate", "tea"]
		res[count] = append(res[count], str)
	}
	

	// 戻り値用の空の多次元スライスを作成
    result := [][]string{}
    for _, r := range res {
    	// 単語だけ取り出して戻り値用のスライスに追加する
    	result = append(result, r)
    }
    
    return result
}

// Top K Frequent Elements
func topKFrequent(nums []int, k int) []int {
	// 空のmapを作成 value => count
	counts := make(map[int]int)
	// それぞれの数値が何個あるかをカウント
	for _, n := range nums {
		counts[n]++
	}
	
	// int型の空のスライス作成、capはnumsの種類数
	keys := make([]int, 0, len(counts))
	// numsの種類数分ループ、numsを空のスライスに追加(numsの重複が排除されたスライスが作れる)
	for key := range counts {
		keys = append(keys, key)
	}
	
	// countsを降順でソート
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})
	
	// 降順ソートされたスライスからtarget分までの数値を返す
	return keys[:k]
}

// Encode and Decode Strings
// 複数の文字列のリストを 1つの文字列にまとめて（encode）、またそこから元の文字列リストに戻す（decode）

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// 文字列結合のための専用バッファ構造体(文字列はイミュータブルだからいちいちコピーしているとメモリを食う)
	var builder strings.Builder // 文字列バッファを作成
	
	// 文字列分ループ
	for _, str := range strs {
		// 連結したときに「長さ#文字列」という形にする => 1#i4#love3#you
		// 最初に「長さ」を書き込み
		builder.WriteString(strconv.Itoa(len(str)))
		// 次に区切り文字の「#」を書き込み
		builder.WriteString("#")
		// 最後に「文字列」を書き込み
		builder.WriteString(str)
	}
	
	// 文字列を結合して返す。
	return builder.String()
}

// 「1#i4#love3#you」みたいな「長さ#文字列」形式の文字列が与えられる
func (s *Solution) Decode(st string) []string {
	// 戻り値に使用する空のスライス
	var result []string
	
	// ループ用のindex
	i := 0
	
	// 文字列の長さ分ループ
	for i < len(st) {
		// '#'までを長さとして読む
		j := i
		
		// 4#, 11# の '4', '11'の文字数分ループ
		for st[j] != '#' {
			j++
		}
		
		// 文字列の中から長さを取得し数値に変換
		length, _ := strconv.Atoi(st[i:j])
		
		// '#'をスキップするためにインクリメント
		j++ 
		
		// 文字列部分を取得して変数に代入
		word := st[j : j+length]
		
		// 戻り値に使用するスライスに文字列を追加
		result = append(result, word)
		
		// 次の長さを取得できるようにindexを調整
		i = j + length
	}
	
	return result
}

// Products of Array Except Self
// Input: nums = [1,2,4,6]
// Output: [48,24,12,8]
func productExceptSelf(nums []int) []int {
	// numsの要素数
	length := len(nums)
	// 戻り値用のスライス
	result := make([]int, length)
	
	// 左からの積を格納(自分の左側の積を格納していく)
	// 例: index 2（値4）のときは、その左側 1*2 = 2 を result[2] に入れる
	// 後から右側だけの積と掛けて、自分以外の数字との積を求める
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i - 1] * nums[i - 1]
		// ループ前 result [1, 0, 0, 0]
		// i=1, 1*1 = result [1, 1, 0, 0]
		// i=2, 1*2 = result [1, 1, 2, 0]
		// i=3, 2*4 = result [1, 1, 2, 8]
	}
	
	// 右からの積をかける
	right := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
		// ループ前 result [1, 1, 2, 8], right = 1
		// i=3, 8*1 = result [1, 1, 2, 8], right = 6 (1*6)
		// i=2, 2*6 = result [1, 1, 12, 8], right = 24 (6*4)
		// i=1, 1*24 = result [1, 24, 12, 8], right = 48 (24*2)
		// i=0, 1*48 = result [48, 24, 12, 8], right = 48 (48*1)
	}
	
	return result
}

// Longest Consecutive Sequence
// Input: nums = [2,20,4,10,3,4,5]
// Output: 4
func longestConsecutive(nums []int) int {
	return 0
}
