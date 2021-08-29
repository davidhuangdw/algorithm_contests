package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class KthLargestNumber5855 {
    fun kthLargestNumber(nums: Array<String>, k: Int): String {
        nums.sortWith(compareBy({ it.length }, { it }))
        return nums[nums.size - k]
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
