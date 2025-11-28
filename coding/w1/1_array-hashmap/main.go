package arrayhashmap

func main() {
	
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
	return nil
}