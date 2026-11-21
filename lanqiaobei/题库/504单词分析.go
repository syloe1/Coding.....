//https://www.lanqiao.cn/problems/504/learning/?page=1&first_category_id=1&name=%E5%8D%95%E8%AF%8D%E5%88%86%E6%9E%90
/*
【问题描述】
小蓝正在学习一门神奇的语言，这门语言中的单词都是由小写英文字母组成，有些单词很长，远远超过正常英文单词的长度。小蓝学了很长时间也记不住一些单词，他准备不再完全记忆这些单词，而是根据单词中哪个字母出现得最多来分辨单词。

现在，请你帮助小蓝，给了一个单词后，帮助他找到出现最多的字母和这个字母出现的次数。

【输入格式】
输入一行包含一个单词，单词只由小写英文字母组成。
对于所有的评测用例，输入的单词长度不超过 1000。

【输出格式】
输出两行，第一行包含一个英文字母，表示单词中出现得最多的字母是哪个。
如果有多个字母出现的次数相等，输出字典序最小的那个。
第二行包含一个整数，表示出现得最多的那个字母在单词中出现的次数。

*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	mx := 0
	anschar := byte('a')
	for i := 0; i < 26; i++ {
		if cnt[i] > mx {
			mx = cnt[i]
			anschar = byte('a' + i)
		}
	}
	fmt.Println(string(anschar))
	fmt.Println(mx)
}
