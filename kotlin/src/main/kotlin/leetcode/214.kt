package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ShortestPalindrome {
    fun shortestPalindrome(s: String): String { //KMP
        val n = s.length
        if (n == 0) return ""
        val pre = IntArray(n)
        for (i in 1 until n) {
            var k = pre[i - 1]
            while (k > 0 && s[k] != s[i]) {
                k = pre[k - 1]
            }
            pre[i] = if (s[k] == s[i]) k + 1 else 0
        }

        val rev = s.reversed()
        val match = IntArray(n)
        for ((i, ch) in rev.withIndex()) {
            var k = if (i == 0) 0 else match[i - 1]
            while (k > 0 && s[k] != ch) {
                k = pre[k - 1]
            }
            match[i] = if (s[k] == ch) k + 1 else 0
        }
        return s.substring(match[n - 1], n).reversed() + s

    }

    fun shortestPalindrome2(s: String): String { //hash
        val n = s.length
        if (n == 0) return ""

        val (P, MOD) = 97 to (1e9 + 7).toLong()
        val pres = hashSetOf<Long>()
        var last = 0L
        for (ch in s) {
            last = (last * P + ch.toInt()) % MOD
            pres.add(last)
        }

        var max = 0
        var (hs, mul) = 0L to 1L
        for (k in 0 until n) {
            hs = ((s[k].toInt() * mul) % MOD + hs) % MOD
            if (hs in pres) max = k
            mul = (mul * P) % MOD
        }

        return s.substring(max + 1, n).reversed() + s
    }

    fun shortestPalindrome1(s: String): String {
        val n = s.length
        if (n == 0) return ""
        for (i in n - 1 downTo 1) {
            var (l, r) = 0 to i
            while (l < r && s[l] == s[r]) {
                l += 1
                r -= 1
            }
            if (l >= r)
                return s.substring(i + 1, n).reversed() + s
        }
        return s.substring(1, n).reversed() + s
    }


    @Test
    fun test1() {
        assertEquals("aaacecaaa", shortestPalindrome("aacecaaa"))
        assertEquals("dcbabcd", shortestPalindrome("abcd"))
        assertEquals("", shortestPalindrome(""))
        assertEquals("aaa", shortestPalindrome("aaa"))
    }
}
