package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class CountPaths5836 {
    fun countPaths(n: Int, roads: Array<IntArray>): Int {
        val MOD = 1000_000_007
        val adj = List(n) { hashMapOf<Int, Long>() }
        for ((u, v, t) in roads) {
            adj[u][v] = t + 0L
            adj[v][u] = t + 0L
        }
        val dis = MutableList(n) { Long.MAX_VALUE }
        val count = MutableList(n) { -1L }
        dis[0] = 0L
        count[0] = 1L
        val que = PriorityQueue<Pair<Long, Int>>(compareBy { it.first })
        que.add(0L to 0)
        while (que.isNotEmpty()) {
            val (d, u) = que.poll()
            if (dis[u] != d) continue
            for ((v, t) in adj[u]) {
                if (dis[u] + t < dis[v]) {
                    dis[v] = dis[u] + t
                    count[v] = count[u]
                    que.add(dis[v] to v)
                } else if (dis[u] + t == dis[v]) {
                    count[v] = (count[v] + count[u]) % MOD
                }
            }
        }
        return count[n - 1].toInt()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
