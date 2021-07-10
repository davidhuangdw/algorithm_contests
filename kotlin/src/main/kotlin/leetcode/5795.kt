package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MinCost {
    fun minCost(maxTime: Int, edges: Array<IntArray>, passingFees: IntArray): Int {
        val n = passingFees.size
        val adj = List(n) { hashMapOf<Int, Int>() }
        for (e in edges) {
            val (a, b, t) = e
            adj[a][b] = minOf(adj[a].getOrDefault(b, Int.MAX_VALUE), t)
            adj[b][a] = minOf(adj[b].getOrDefault(a, Int.MAX_VALUE), t)
        }
        val min_time = MutableList(n) { Int.MAX_VALUE }

        data class Node(val c: Long, val i: Int, val t: Int) : Comparable<Node> {
            override fun compareTo(other: Node): Int {
                if (c != other.c) return compareValues(c, other.c)
                return t - other.t
            }
        }

        val que = PriorityQueue<Node>()
        que.add(Node(0L + passingFees[0], 0, 0))
        min_time[0] = 0
        while (que.isNotEmpty()) {
            val (c, t, i) = que.poll()
            for ((j, et) in adj[i]) {
                val tj = t + et
                val jc = c + passingFees[j]
                if (tj > maxTime) continue
                if (tj >= min_time[j]) continue
                min_time[j] = tj
                if (j == n - 1) return jc.toInt()
                que.add(Node(jc, tj, j))
            }
        }
        return -1
    }
    // todo: to fix


    @Test
    fun test1() {
        assertEquals(
            11,
            minCost(
                30,
                parse2dIntArray("[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]"),
                parseIntArray("[5,1,2,20,20,3]")
            )
        )
    }
}
