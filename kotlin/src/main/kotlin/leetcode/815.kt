package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumBusesToDestination {
    fun numBusesToDestination(routes: Array<IntArray>, source: Int, target: Int): Int {
        if(source == target) return 0
        val adj_routes = hashMapOf<Int, MutableList<Int>>()

        for ((i, route) in routes.withIndex()) {
            for (u in route) {
                if (u !in adj_routes)
                    adj_routes[u] = mutableListOf()
                adj_routes[u]!!.add(i)
            }
        }

        val vis = BooleanArray(routes.size)
        val min_cnt = hashMapOf<Int, Int>()
        min_cnt[source] = 0
        val que = ArrayDeque<Int>()
        que.add(source)

        while (que.isNotEmpty()) {
            val u = que.removeFirst()
            val c = min_cnt[u]!! + 1
            for (i in adj_routes.getOrDefault(u, mutableListOf())) {
                if (!vis[i]) {
                    vis[i] = true
                    for (v in routes[i]) {
                        if (v !in min_cnt) {
                            min_cnt[v] = c
                            if (v == target) return c
                            que.add(v)
                        }
                    }
                }
            }
        }

        return -1
    }

    @Test
    fun test1() {
        assertEquals(
            2, numBusesToDestination(
                parse2dIntArray("[[1,2,7],[3,6,7]]"), 1, 6
            )
        )

        assertEquals(
            -1, numBusesToDestination(
                parse2dIntArray("[[7,12],[4,5,15],[6],[15,19],[9,12,13]]"), 15, 12
            )
        )
    }
}