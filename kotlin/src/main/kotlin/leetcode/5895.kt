package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinOperations5895 {
    fun minOperations(grid: Array<IntArray>, x: Int): Int {
        val m = grid[0][0] % x
        val vs = mutableListOf<Int>()
        for (r in grid)
            for (v in r) {
                if (v % x != m) return -1
                vs.add(v / x)
            }
        vs.sort()
        val n = vs.size
        val md = vs[(n - 1) / 2]
        var sum = 0
        for (v in vs)
            sum += if (md >= v) md - v else v - md
        return sum
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
