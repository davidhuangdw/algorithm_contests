package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LatestDayToCross5845 {
    fun latestDayToCross(n: Int, m: Int, cells: Array<IntArray>): Int {
        val cs = cells.map { it[0] - 1 to it[1] - 1 }
        val cur = List(n) { MutableList(m) { 0 } }
        for ((i, j) in cs)
            cur[i][j] = 1

        fun enc(i: Int, j: Int) = i * m + j
        val par = MutableList(n * m + 2) { it }
        fun find(x: Int): Int {
            if (par[x] != x)
                par[x] = find(par[x])
            return par[x]
        }

        fun union(x: Int, y: Int) {
            val (rx, ry) = find(x) to find(y)
            if (rx != ry)
                par[rx] = ry
        }
        val (s, t) = n * m to n * m + 1

        for (j in 0 until m) {
            union(enc(0, j), s)
            union(enc(n - 1, j), t)
        }

        fun valid(i: Int, j: Int) = i in 0 until n && j in 0 until m && cur[i][j] == 0
        for (i in 0 until n)
            for (j in 0 until m) {
                if (valid(i, j)) {
                    if (valid(i + 1, j))
                        union(enc(i, j), enc(i + 1, j))
                    if (valid(i, j + 1))
                        union(enc(i, j), enc(i, j + 1))
                }
            }

        if (find(s) == find(t)) return cs.size
        val dirs = listOf(0 to 1, 0 to -1, 1 to 0, -1 to 0)
        for (k in cs.size - 1 downTo 0) {
            val (i, j) = cs[k]
            cur[i][j] = 0
            for ((di, dj) in dirs) {
                val (ni, nj) = i + di to j + dj
                if (valid(ni, nj))
                    union(enc(i, j), enc(ni, nj))
            }
            if (find(s) == find(t)) return k
        }
        return -1
    }


    @Test
    fun test1() {
        assertEquals(2, latestDayToCross(2, 2, parse2dIntArray("[[1,1],[2,1],[1,2],[2,2]]")))
        assertEquals(1, latestDayToCross(2, 2, parse2dIntArray("[[1,1],[1,2],[2,1],[2,2]]")))
        assertEquals(
            3,
            latestDayToCross(3, 3, parse2dIntArray("[[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]"))
        )
    }
}
