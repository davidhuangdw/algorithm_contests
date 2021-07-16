package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CheckSubarraySum523 {
    fun checkSubarraySum(nums: IntArray, k: Int): Boolean {
        val lpos = hashMapOf<Int, Int>()
        lpos[0] = -1
        var sum = 0
        for ((i, v) in nums.withIndex()) {
            sum += v
            val d = sum % k
            if (d !in lpos)
                lpos[d] = i
            else if (lpos[d]!! <= i - 2)
                return true
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(true, checkSubarraySum(parseIntArray("[23,2,4,6,6]"), 7))
        assertEquals(true, checkSubarraySum(parseIntArray("[23,2,4,6,8]"), 6))
        assertEquals(true, checkSubarraySum(parseIntArray("[23,2,6,4,7]"), 6))
        assertEquals(false, checkSubarraySum(parseIntArray("[23,2,6,4,7]"), 13))
    }
}
