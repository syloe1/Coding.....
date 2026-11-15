/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
0 <= s.length <= 1e5
*/
function lengthOfLongestSubstring(s: string): number {
    const cnt = new Array(128).fill(0);
    let left = 0;
    let ans = 0;
    for (let right = 0; right < s.length; right++) {
        const c = s.charCodeAt(right);
        cnt[c]++;
        while (cnt[c] > 1) {
            const leftC = s.charCodeAt(left);
            cnt[leftC]--;
            left++;
        }
        ans = Math.max(ans, right - left + 1);
    }
    return ans;
}
