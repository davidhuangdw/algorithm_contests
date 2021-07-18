package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class AddRungs5814 {
    fun addRungs(rungs: IntArray, dist: Int): Int {
        val h = listOf(0) + rungs.toList()
        var res = 0
        for (i in 1 until h.size) {
            val x = h[i] - h[i - 1]
            res += (x - 1) / dist
        }
        return res
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
