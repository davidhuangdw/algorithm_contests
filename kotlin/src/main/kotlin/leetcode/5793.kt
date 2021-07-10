package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NearestExit {
    fun nearestExit(maze: Array<CharArray>, entrance: IntArray): Int {
        val n = maze.size
        val m = maze[0].size

        val dis = hashMapOf<Pair<Int, Int>, Int>()
        val que = ArrayDeque<Pair<Int, Int>>()
        val st = entrance[0] to entrance[1]
        que.add(st)
        dis[st] = 0

        fun edge(i: Int, j: Int) = i == 0 || i == n - 1 || j == 0 || j == m - 1
        fun valid(i: Int, j: Int) = i in 0 until n && j in 0 until m && maze[i][j] == '.' && i to j !in dis
        while (que.isNotEmpty()) {
            val node = que.removeFirst()
            val d = dis[node]!!
            val (i, j) = node
            for ((ni, nj) in listOf(i - 1 to j, i + 1 to j, i to j - 1, i to j + 1)) {
                if (!valid(ni, nj)) continue
                if (edge(ni, nj)) return d + 1
                dis[ni to nj] = d + 1
                que.add(ni to nj)
            }
        }
        return -1
    }


    @Test
    fun test1() {
        assertEquals(
            2, nearestExit(
                arrayOf(
                    "+++".toCharArray(),
                    "...".toCharArray(),
                    "+++".toCharArray(),
                ),
                intArrayOf(1, 0)
            )
        )
    }
}
