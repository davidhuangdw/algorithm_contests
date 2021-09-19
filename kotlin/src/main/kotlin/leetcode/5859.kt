package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountKDifference5859 {
    fun countKDifference(nums: IntArray, k: Int): Int {
        val count = hashMapOf<Int, Int>()
        var cnt = 0
        for (v in nums) {
            cnt += count.getOrDefault(v - k, 0) + count.getOrDefault(v + k, 0)
            count[v] = count.getOrDefault(v, 0) + 1
        }
        return cnt
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
