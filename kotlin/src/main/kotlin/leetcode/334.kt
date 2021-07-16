package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IncreasingTriplet334 {
    fun increasingTriplet(nums: IntArray): Boolean {
        var (a, b) = nums[0] to Int.MAX_VALUE
        for (v in nums) {
            if (v > b) return true
            a = minOf(a, v)
            if (v > a) {
                b = minOf(b, v)
            }
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(true, increasingTriplet(parseIntArray("[2,1,5,0,4,6]")))
        assertEquals(false, increasingTriplet(parseIntArray("[10, 12, 8, 10, 6, 8, 4]")))
    }
}
