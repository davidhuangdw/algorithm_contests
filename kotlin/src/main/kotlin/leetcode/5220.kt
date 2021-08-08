package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxProduct5220 {
    fun maxProduct(s: String): Long {
        val n = s.length
        var (l, r) = 0 to -1
        val md_longest = MutableList(n) { 0 }
        for (i in 0 until n) {
            var k = if (i > r) 1 else minOf(r + 1 - i, md_longest[l + r - i])
            while (0 <= i - k && i + k < n && s[i - k] == s[i + k])
                k += 1
            md_longest[i] = k
            if (i + k - 1 > r) {
                r = i + k - 1
                l = i - k + 1
            }
        }

        val left = MutableList(n) { 1 }
        val right = MutableList(n) { 1 }
        r = 1
        for (i in 0 until n) {
            val k = md_longest[i]
            while (r < i + k - 1) {
                r += 1
                left[r] = maxOf(left[r - 1], (r - i) * 2 + 1)
            }
        }
        l = n - 1
        for (i in n - 1 downTo 0) {
            val k = md_longest[i]
            while (i - k + 1 < l) {
                l -= 1
                right[l] = maxOf(right[l + 1], (i - l) * 2 + 1)
            }
        }

        var mul = 0L
        for (i in 0 until n - 1)
            mul = maxOf(mul, left[i] * 1L * right[i + 1])
        return mul
    }


    @Test
    fun test1() {
        assertEquals(9, maxProduct("ababbba"))
        assertEquals(15, maxProduct("zaaaxbbbbby"))
        assertEquals(41, maxProduct("rofcjxfkbzcvvlbkgcwtcjctwcgkblvvczbkfxjcfor"))
    }
}
