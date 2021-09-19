package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxTaxiEarnings5861 {
    fun maxTaxiEarnings(n: Int, rides: Array<IntArray>): Long {
        val dp = MutableList(n + 1) { 0L }
        val ends = (0..n).map { mutableListOf<IntArray>() }
        for (r in rides)
            ends[r[1]].add(r)
        for (e in 1..n) {
            dp[e] = dp[e - 1]
            for ((s, _, t) in ends[e]) {
                dp[e] = maxOf(dp[e], dp[s] + e - s + t)
            }
        }
        return dp[n]
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
