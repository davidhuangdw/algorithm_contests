package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class WiggleMaxLength376 {
    fun wiggleMaxLength(nums: IntArray): Int {
        var (up, down) = 1 to 1
        for (i in 1 until nums.size) {
            val d = nums[i] - nums[i - 1]
            if (d == 0) continue
            if (d > 0) up = down + 1
            else down = up + 1
        }
        return maxOf(up, down)
    }


    @Test
    fun test1() {
        assertEquals(6, wiggleMaxLength(parseIntArray("[1,1,7,4,9,2,5]")))
        assertEquals(6, wiggleMaxLength(parseIntArray("[1,7,4,9,2,5]")))
        assertEquals(7, wiggleMaxLength(parseIntArray("[1,17,5,10,13,15,10,5,16,8]")))
        assertEquals(2, wiggleMaxLength(parseIntArray("[1,2,3,4,5,6,7,8,9]")))
    }
}
