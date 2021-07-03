package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaximumGap {
    fun maximumGap(nums: IntArray): Int {
        val distinct = nums.toHashSet()
        val n = distinct.size
        if (n < 2) return 0
        var (max, min) = Int.MIN_VALUE to Int.MAX_VALUE
        for (v in distinct) {
            max = maxOf(max, v)
            min = minOf(min, v)
        }

        val d = ((max + 1 - min) + (n - 1)) / n         // d>=1 because distinct
        val buckets = hashMapOf<Int, Pair<Int, Int>>()
        for (v in distinct) {
            val key = (v - min) / d
            val (a, b) = buckets.getOrDefault(key, Int.MAX_VALUE to Int.MIN_VALUE)
            buckets[key] = minOf(a, v) to maxOf(b, v)
        }
        var max_diff = 0
        var lmax = min
        for (i in 0 until n)             // n buckets: impossible inside the same bucket, because otherwise one bucket will be empty
            if (i in buckets) {
                val (a, b) = buckets[i]!!
                max_diff = maxOf(max_diff, a - lmax)
                lmax = b
            }
        return max_diff
    }


    @Test
    fun test1() {
        assertEquals(3, maximumGap(parseIntArray("[3,6,9,1]")))
        assertEquals(0, maximumGap(parseIntArray("[10]")))
        assertEquals(0, maximumGap(parseIntArray("[3, 3, 3]")))
        assertEquals(4, maximumGap(parseIntArray("[1, 11, 2, 3, 9, 7]")))
    }
}
