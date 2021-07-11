package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GetConcatenation {
    fun getConcatenation(nums: IntArray): IntArray {
        val n = nums.size
        val res = IntArray(n * 2)
        for ((i, v) in nums.withIndex()) {
            res[i] = v
            res[i + n] = v
        }
        return res
    }
}
