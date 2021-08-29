package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimumDifference5854 {
    fun minimumDifference(nums: IntArray, k: Int): Int {
        val s = nums.sorted()
        var min = Int.MAX_VALUE
        for (i in 0..s.size - k) {
            min = minOf(min, s[i + k - 1] - s[i])
        }
        return min
    }


    @Test
    fun test1() {
        assertEquals(0, minimumDifference(parseIntArray("[90]"), 1))
    }
}
