package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import kotlin.math.absoluteValue

class NearestPalindromic {
    fun nearestPalindromic(n: String): String {
        val nv = n.toLong()
        var (min_diff, res) = Long.MAX_VALUE to 0L
        for (src in nv - 1..nv + 1) {
            val ss = src.toString()
            val l = ss.substring(0, (ss.length + 1) / 2)
            for (d in -9..9) {
                val left = l.toLong() + d
                if (left < 0) continue
                val ls = left.toString()
                for (md in ls.length - 1..ls.length) {
                    val cur = (ls + ls.substring(0, md).reversed()).toLong()
                    if (cur == nv) continue
                    val abs = if (cur > nv) cur - nv else nv - cur
                    if (abs < min_diff || (abs == min_diff && cur < res)) {
                        min_diff = abs
                        res = cur
                    }
                }
            }
        }
        return res.toString()
    }

    @Test
    fun test1() {
        assertEquals("1001", nearestPalindromic("999"))
        assertEquals("9", nearestPalindromic("11"))
        assertEquals("9", nearestPalindromic("10"))
        assertEquals("8", nearestPalindromic("9"))
        assertEquals("99", nearestPalindromic("100"))
        assertEquals("121", nearestPalindromic("123"))
        assertEquals("919", nearestPalindromic("920"))
        assertEquals("99", nearestPalindromic("100"))
    }
}