package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxPoints {
    fun maxPoints(points: Array<IntArray>): Long {
        val (n, m) = points.size to points[0].size
        var f = MutableList(m) { points[0][it].toLong() }
        val lmax = MutableList(m) { 0L }
        val rmax = MutableList(m + 1) { -1000000L }
        for (i in 1 until n) {
            for (j in m - 1 downTo 0)
                rmax[j] = maxOf(rmax[j + 1], f[j] - j)
            lmax[0] = f[0]
            for (j in 1 until m)
                lmax[j] = maxOf(lmax[j - 1], f[j] + j)

            val nxt = MutableList(m) { 0L }
            for (j in 0 until m)
                nxt[j] = points[i][j] + maxOf(lmax[j] - j, rmax[j + 1] + j)
            f = nxt
        }
        return f.maxOrNull()!!
    }


    @Test
    fun test1() {
        assertEquals(15, maxPoints(parse2dIntArray("[[4,1,0,4,0],[1,0,4,0,5],[1,3,0,4,1],[4,4,0,4,0]]")))
        assertEquals(9, maxPoints(parse2dIntArray("[[1,2,3],[1,5,1],[3,1,1]]")))
        assertEquals(11, maxPoints(parse2dIntArray("[[1,5],[2,3],[4,2]]")))
    }
}
