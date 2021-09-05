package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindMiddleIndex5846 {
    fun findMiddleIndex(nums: IntArray): Int {
        var (l, r) = 0 to nums.sum()
        for ((i, v) in nums.withIndex()) {
            r -= v
            if (l == r) return i
            l += v
        }
        return -1
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
