package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinDistance {
    fun minDistance(word1: String, word2: String): Int {
        val m = word2.length
        val d = MutableList(m + 1) { it }
        for (ch in word1) {
            var pre = d[0]
            d[0] += 1
            for ((i, c) in word2.withIndex()) {
                pre = d[i + 1].also{
                    d[i + 1] = if (c == ch) pre else 1 + minOf(d[i], pre, d[i + 1])
                }
            }
        }
        return d[m]
    }

    @Test
    fun test1() {
        assertEquals(3, minDistance("horse", "ros"))
        assertEquals(5, minDistance("intention", "execution"))
    }
}
