package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountQuadruplets5863 {
    fun countQuadruplets(nums: IntArray): Int {
        val count = hashMapOf<Int, Int>()
        for (v in nums)
            count[v] = count.getOrDefault(v, 0) + 1

        var sum = 0
        for ((i, v) in nums.withIndex()) {
            count[v] = count[v]!! - 1
            for (j in 0 until i)
                for (k in 0 until j) {
                    val s = v + nums[j] + nums[k]
                    sum += count.getOrDefault(s, 0)
                }
        }
        return sum
    }


    @Test
    fun test1() {
        assertEquals(2, countQuadruplets(parseIntArray("[9,6,8,23,39,23]")))
    }
}
