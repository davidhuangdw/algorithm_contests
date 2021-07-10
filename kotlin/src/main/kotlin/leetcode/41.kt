package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FirstMissingPositive41 {
    fun firstMissingPositive(nums: IntArray): Int {
        val n = nums.size
        for (v in nums) {
            var j = v
            while (j in 1..n && nums[j - 1] != j) {
                nums[j - 1] = j.also { j = nums[j - 1] }
            }
        }
        return (1..n).firstOrNull { nums[it - 1] != it } ?: n + 1
    }


    @Test
    fun test1() {
        assertEquals(3, firstMissingPositive(parseIntArray("[1,2,0]")))
        assertEquals(2, firstMissingPositive(parseIntArray("[3,4,-1,1]")))
        assertEquals(1, firstMissingPositive(parseIntArray("[7,8,9,11,12]")))
        assertEquals(5, firstMissingPositive(parseIntArray("[4,3,1,2]")))
    }
}
