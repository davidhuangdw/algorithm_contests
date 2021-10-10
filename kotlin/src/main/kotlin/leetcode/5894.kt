package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RwoOutOfThree5894 {
    fun twoOutOfThree(nums1: IntArray, nums2: IntArray, nums3: IntArray): List<Int> {
        val a = nums1.toHashSet()
        val b = nums2.toHashSet()
        val c = nums3.toHashSet()
        return (a.intersect(b)).union(b.intersect(c)).union(c.intersect(a)).toList()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
