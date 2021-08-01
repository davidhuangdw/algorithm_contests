package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountSpecialSubsequences5833 {
    fun countSpecialSubsequences(nums: IntArray): Int {
        val MOD = 1e9.toLong() + 7
        var sum = 0L
        val pre = mutableListOf(0L, 0L, 0L)
        for (v in nums) {
            if (v == 0) {
                pre[0] = (pre[0] + pre[0] + 1) % MOD
            } else if (v == 1) {
                pre[1] = (pre[1] + pre[0] + pre[1]) % MOD
            } else {
                sum = (sum + pre[1] + pre[2]) % MOD
                pre[2] = (pre[2] + pre[1] + pre[2]) % MOD
            }
        }
        return sum.toInt()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
