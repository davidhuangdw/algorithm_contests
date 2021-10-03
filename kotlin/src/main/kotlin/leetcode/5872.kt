package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MumOfPairs5872 {
    fun numOfPairs(nums: Array<String>, target: String): Int {
        val n = nums.size
        var cnt = 0
        for (i in 0 until n)
            for (j in 0 until n)
                if (i != j) {
                    if (nums[i] + nums[j] == target)
                        cnt += 1
                }

        return cnt
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
