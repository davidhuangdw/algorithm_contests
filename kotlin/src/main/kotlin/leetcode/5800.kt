package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class BuildArray {
    fun buildArray(nums: IntArray): IntArray {
        val n = nums.size
        val res = IntArray(n)
        for (i in 0 until n)
            res[i] = nums[nums[i]]
        return res
    }
}
