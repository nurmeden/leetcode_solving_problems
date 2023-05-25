from collections import Counter

class Solution:
    def lengthOfLongestSubstring(s: str) -> int:
        chars = Counter()
        left = right = 0
        
        res = 0
        while right < len(s):
            r = s[right]
            chars[r] += 1

            while chars[r] > 1:
                l = s[left]
                chars[l] -= 1
                left += 1

            res = max(res, right - left + 1)
            print(r)
            right += 1
        return res

print(Solution.lengthOfLongestSubstring("abcabcbb"))