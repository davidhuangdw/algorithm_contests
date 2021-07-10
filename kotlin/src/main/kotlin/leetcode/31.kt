package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NextPermutation31 {
    fun nextPermutation(nums: IntArray): Unit {
        val n = nums.size
        if (n <= 1) return

        fun swap(i: Int, j: Int) {
            nums[i] = nums[j].also { nums[j] = nums[i] }
        }

        fun rev(i: Int, j: Int) {
            for (k in 0 until (j - i + 1) / 2)
                swap(i + k, j - k)
        }

        val low = (n - 2 downTo 0).firstOrNull { nums[it] < nums[it + 1] }
        if (low == null) {
            rev(0, n - 1)
        } else {
            val high = (n - 1 downTo low + 1).first { nums[it] > nums[low] }
            swap(low, high)
            rev(low + 1, n - 1)
        }
    }

    @Test
    fun test1() {
        var nums = parseIntArray("[1,2,3]")
        nextPermutation(nums)
        assertEquals(parseIntArray("[1,3,2]").toList(), nums.toList())

        nums = parseIntArray("[3,2,1]")
        nextPermutation(nums)
        assertEquals(parseIntArray("[1,2,3]").toList(), nums.toList())

        nums = parseIntArray("[1,1,5]")
        nextPermutation(nums)
        assertEquals(parseIntArray("[1,5,1]").toList(), nums.toList())
    }
}
