package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GetMaximumGold1219 {
    fun getMaximumGold(grid: Array<IntArray>): Int {
        val (n, m) = grid.size to grid[0].size

        fun dfs(i: Int, j: Int): Int {
            if (grid[i][j] == 0) return 0
            var x = grid[i][j]
            grid[i][j] = 0
            var max = 0
            for ((ni, nj) in listOf(i to j - 1, i to j + 1, i - 1 to j, i + 1 to j))
                if (ni in (0 until n) && nj in (0 until m))
                    max = maxOf(max, dfs(ni, nj))
            grid[i][j] = x
            return x + max
        }

        var max = 0
        for (i in 0 until n)
            for (j in 0 until m)
                max = maxOf(max, dfs(i, j))
        return max
    }


    @Test
    fun test1() {
        assertEquals(24, getMaximumGold(parse2dIntArray("[[0,6,0],[5,8,7],[0,9,0]]")))
        assertEquals(28, getMaximumGold(parse2dIntArray("[[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]")))
    }
}
