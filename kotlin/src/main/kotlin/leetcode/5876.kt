package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SumOfBeauties5876 {
    fun sumOfBeauties(nums: IntArray): Int {
        val n = nums.size
        val lmax = nums.toMutableList()
        val rmin = nums.toMutableList()
        for (i in 1 until n)
            lmax[i] = maxOf(lmax[i - 1], nums[i])
        for (i in n - 2 downTo 0)
            rmin[i] = minOf(rmin[i + 1], nums[i])

        var sum = 0
        for (i in 1..n - 2) {
            val v = nums[i]
            sum += if (lmax[i - 1] < v && v < rmin[i + 1]) 2
            else if (nums[i - 1] < v && v < nums[i + 1]) 1
            else 0
        }
        return sum
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
