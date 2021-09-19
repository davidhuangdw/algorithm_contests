package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinOperations5862 {
    fun minOperations(nums: IntArray): Int {
        val v = nums.toHashSet().sorted()       // bug: forget to remove repeat ones
        val n = nums.size
        var r = 0
        var min = n
        for ((i, x) in v.withIndex()) {
            while (r < v.size && v[r] < x + n)
                r += 1
            min = minOf(min, n - (r - i))
        }
        return min
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
