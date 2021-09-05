package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindFarmland5847 {
    fun findFarmland(land: Array<IntArray>): Array<IntArray> {
        val (n, m) = land.size to land[0].size
        fun bfs(si: Int, sj: Int): IntArray {
            var (i, j) = si + 1 to sj + 1
            while (i < n && land[i][sj] == 1)
                i += 1
            while (j < m && land[si][j] == 1)
                j += 1
            for (x in si until i)
                for (y in sj until j)
                    land[x][y] = 0
            return intArrayOf(si, sj, i - 1, j - 1)
        }

        val res = mutableListOf<IntArray>()
        for (i in 0 until n)
            for (j in 0 until m)
                if (land[i][j] == 1)
                    res.add(bfs(i, j))
        return res.toTypedArray()
    }

    @Test
    fun test1() {
        assertEquals(1, findFarmland(parse2dIntArray("[[1,0,0],[0,1,1],[0,1,1]]")).map { it.toList() })
    }
}
