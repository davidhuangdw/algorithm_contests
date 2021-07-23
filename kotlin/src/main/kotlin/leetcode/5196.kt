package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CanSeePersonsCount5196 {
    fun canSeePersonsCount(heights: IntArray): IntArray {
        val n = heights.size
        val dec = mutableListOf<Int>()
        val res = MutableList(n) { 0 }
        for ((i, h) in heights.withIndex()) {
            while (dec.isNotEmpty() && heights[dec.last()] < h) {
                res[dec.last()] += 1
                dec.removeAt(dec.size - 1)
            }
            if (dec.isNotEmpty()) res[dec.last()] += 1
            dec.add(i)
        }
        return res.toIntArray()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
