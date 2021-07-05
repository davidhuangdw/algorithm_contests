package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class TrapRainWater {
    fun trapRainWater(heightMap: Array<IntArray>): Int {
        val (n, m) = heightMap.size to heightMap[0].size
        val done = hashMapOf<Pair<Int, Int>, Int>()

        fun getAdjUnits(i: Int, j: Int): List<Pair<Int, Int>> {
            return listOf(i - 1 to j, i + 1 to j, i to j - 1, i to j + 1).filter {
                it.first in 0 until n && it.second in 0 until m && it !in done
            }
        }

        val que = PriorityQueue<Pair<Int, Pair<Int, Int>>>(compareBy { it.first })
        for (i in 0 until n)
            for (j in 0 until m)
                if (i == 0 || i == n - 1 || j == 0 || j == m - 1) {
                    done[i to j] = heightMap[i][j]
                    for (pos in getAdjUnits(i, j))
                        que.add(heightMap[i][j] to pos)
                }

        var sum = 0
        while (que.isNotEmpty()) {
            val (h, pos) = que.poll()
            if (pos in done) continue
            val (i, j) = pos
            val now = maxOf(h, heightMap[i][j])
            done[pos] = now
            sum += now - heightMap[i][j]

            for (nxt in getAdjUnits(i, j))
                que.add(now to nxt)
        }

        return sum
    }


    @Test
    fun test1() {
        assertEquals(4, trapRainWater(parse2dIntArray("[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]")))
    }
}
