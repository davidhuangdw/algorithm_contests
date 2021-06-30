package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountSubIslands {
    fun countSubIslands(grid1: Array<IntArray>, grid2: Array<IntArray>): Int {
        val (M, N) = grid1.size to grid1[0].size
        fun build(grid: Array<IntArray>): Pair<HashMap<Pair<Int, Int>, Int>, MutableList<List<Pair<Int, Int>>>> {
            var pos_to_id = hashMapOf<Pair<Int, Int>, Int>()
            var islands = mutableListOf<List<Pair<Int, Int>>>()

            var cur = mutableListOf<Pair<Int, Int>>()
            var cnt = 0
            fun dfs(i: Int, j: Int) {
                if (grid[i][j] == 0 || i to j in pos_to_id) return
                cur.add(i to j)
                pos_to_id[i to j] = cnt
                for ((ni, nj) in listOf(i - 1 to j, i + 1 to j, i to j - 1, i to j + 1))
                    if (ni in 0 until M && nj in 0 until N) {
                        dfs(ni, nj)
                    }
            }

            for (i in 0 until M)
                for (j in 0 until N)
                    if (grid[i][j] == 1 && i to j !in pos_to_id) {
                        cur = mutableListOf()
                        cnt += 1
                        dfs(i, j)
                        islands.add(cur)
                    }
            return pos_to_id to islands
        }

        val pos_to_id1 = build(grid1).first
        val islands2 = build(grid2).second

        var res = 0
        for (island in islands2) {
            val ids = island.map { pos_to_id1.getOrDefault(it, -1) }.toSet()
            if (ids.size == 1 && -1 !in ids)
                res += 1
        }
        return res
    }

    @Test
    fun test1() {
        var g1 = arrayOf(
            intArrayOf(1, 1, 1, 0, 0),
            intArrayOf(0, 1, 1, 1, 1),
            intArrayOf(0, 0, 0, 0, 0),
            intArrayOf(1, 0, 0, 0, 0),
            intArrayOf(1, 1, 0, 1, 1),
        )

        var g2 = arrayOf(
            intArrayOf(1, 1, 1, 0, 0),
            intArrayOf(0, 0, 1, 1, 1),
            intArrayOf(0, 1, 0, 0, 0),
            intArrayOf(1, 0, 1, 1, 0),
            intArrayOf(0, 1, 0, 1, 0),
        )

        assertEquals(3, countSubIslands(g1, g2))
    }

}