package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LongestPalindrome5 {
    fun longestPalindrome(os: String): String {     // O(n) Manacher
        val s = buildString {
            append("$#")
            for (ch in os)
                append(ch, '#')
            append("|")
        }
        val n = s.length
        var (mi, l, r) = listOf(0, 0, 0)
        val odd_longest = MutableList(n) { 0 }
        for (i in 1 until n - 1) {
            var k = if (r < i) 1 else minOf(r - i + 1, odd_longest[l + r - i])
            while (s[i - k] == s[i + k])
                k += 1
            odd_longest[i] = k
            if (odd_longest[i] > odd_longest[mi])    // len is k-1 both odd and even case
                mi = i
            if (i + k - 1 > r) {
                r = i + k - 1
                l = i - k + 1
            }
        }

        val k = odd_longest[mi]
        return buildString {
            for (i in (mi - k + 2)..(mi + k - 2) step 2)
                append(s[i])
        }
    }


    @Test
    fun test1() {
        assertEquals("bab", longestPalindrome("babad"))
        assertEquals("bb", longestPalindrome("cbbd"))
        assertEquals("a", longestPalindrome("a"))
        assertEquals("a", longestPalindrome("ac"))
    }
}
