package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxValue1881 {
    fun maxValue(n: String, x: Int): String {
        val m = n.length
        var xc = '0' + x
        val k = if (n[0] == '-') {
            (1 until m).firstOrNull { n[it] > xc } ?: m
        } else {
            (0 until m).firstOrNull { n[it] < xc } ?: m
        }

        return n.substring(0, k) + xc + n.substring(k, m)
    }


    @Test
    fun test1() {
        assertEquals("999", maxValue("99", 9))
        assertEquals("88778877766", maxValue("8877887766", 7))
        assertEquals("-123", maxValue("-13", 2))
        assertEquals("-112211293", maxValue("-11221193", 2))
    }
}
