package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RepeatedSubstringPattern459 {
    fun repeatedSubstringPattern(s: String): Boolean {
        val n = s.length
        if (n <= 1) return false
        val max_match = MutableList(n) { 0 }
        for (i in 1 until n) {
            var k = max_match[i - 1]
            while (k > 0 && s[k] != s[i])
                k = max_match[k - 1]
            max_match[i] = if (s[k] == s[i]) k + 1 else 0
        }
        return max_match[n - 1] > 0 && n % (n - max_match[n - 1]) == 0
    }

    @Test
    fun test1() {
        assertEquals(true, repeatedSubstringPattern("abcabcabc"))
        assertEquals(false, repeatedSubstringPattern("abcabcab"))
    }
}
