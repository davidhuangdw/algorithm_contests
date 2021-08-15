package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RearrangeArray5832 {
    fun rearrangeArray(nums: IntArray): IntArray {
        nums.sort()
        val n = nums.size
        val res = IntArray(n)
        for (i in 0 until (n + 1) / 2)
            res[i * 2] = nums[i]
        val hf = (n + 1) / 2
        var j = 1
        for (i in 0 until n / 2)
            res[j + i * 2] = nums[hf + i]
        return res
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
