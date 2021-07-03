package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LargestPerimeter {
    fun largestPerimeter(nums: IntArray): Int {
        nums.sort()
        val n = nums.size
        for (i in n - 1 downTo 2) {
            val j = i - 2
            if (nums[i] < nums[j] + nums[j + 1])
                return nums[i] + nums[j] + nums[j + 1]
        }
        return 0
    }

    @Test
    fun test1() {
        assertEquals(0, largestPerimeter(parseIntArray("[1,2,1]")))
        assertEquals(5, largestPerimeter(parseIntArray("[2,2,1]")))
        assertEquals(8, largestPerimeter(parseIntArray("[3,6,2,3]")))
    }
}
