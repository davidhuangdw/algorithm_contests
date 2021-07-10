package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxProfit188 {
    fun maxProfit(k: Int, prices: IntArray): Int {
        val n = prices.size
        val b = MutableList(k + 1) { if (it == 0) 0 else (-1e8).toInt() }
        val s = MutableList(k + 1) { 0 }
        for (j in 1..n) {
            for (i in k downTo 1) {
                val p = prices[j - 1]
                b[i] = maxOf(b[i], s[i - 1] - p)
                s[i] = maxOf(s[i], b[i] + p)
            }
        }
        return s[k]
    }


    @Test
    fun test1() {
        assertEquals(2, maxProfit(2, parseIntArray("[2,4,1]")))
        assertEquals(7, maxProfit(2, parseIntArray("[3,2,6,5,0,3]")))
    }
}
