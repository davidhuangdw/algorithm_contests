package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinSpaceWastedKResizing5828 {
    fun minSpaceWastedKResizing(nums: IntArray, k: Int): Int {
        val n = nums.size
        val max_range = List(n) { MutableList(n) { 0 } }
        val pre_sum = MutableList(n + 1) { 0 }
        for (i in 1..n)
            pre_sum[i] = pre_sum[i - 1] + nums[i - 1]
        for (i in 0 until n) {
            max_range[i][i] = nums[i]
            for (j in i + 1 until n)
                max_range[i][j] = maxOf(nums[j], max_range[i][j - 1])
        }

        val cost = IntArray(n + 1) { l -> if (l == 0) 0 else max_range[0][l - 1] * l - pre_sum[l] }
        cost[0] = 0
        for (t in 1..k) {
            for (l in n downTo 1) {
                var x = cost[l]
                fun calc(j: Int) = cost[j] + (max_range[j][l - 1] * (l - j) - (pre_sum[l] - pre_sum[j]))
                for (j in 0 until l)
                    x = minOf(x, calc(j))
                cost[l] = x
            }
        }
        return cost[n]
    }


    @Test
    fun test1() {
        assertEquals(10, minSpaceWastedKResizing(parseIntArray("[10,20]"), 0))
        assertEquals(10, minSpaceWastedKResizing(parseIntArray("[10,20,30]"), 1))
        assertEquals(15, minSpaceWastedKResizing(parseIntArray("[10,20,15,30,20]"), 2))
    }
}
