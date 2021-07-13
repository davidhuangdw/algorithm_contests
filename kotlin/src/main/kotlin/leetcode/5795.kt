package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MinCost {
    // key idea:
    // 1. (time, city) as index(vertex)
    // 2. no cycle(able to O(n) topo-order update): because only old time update later time
    fun minCost(maxTime: Int, edges: Array<IntArray>, passingFees: IntArray): Int {
        val MX = 1e8.toInt()
        val n = passingFees.size
        val adj = List(n) { hashSetOf<Int>() }
        val min_cost = List(maxTime + 1) { MutableList(n) { MX } }
        min_cost[0][0] = passingFees[0]

        for(t in 1..maxTime){
            for((a, b, et) in edges){
                val pre = t - et
                if(pre >= 0){
                    min_cost[t][a] = minOf(min_cost[t][a], min_cost[pre][b] + passingFees[a])
                    min_cost[t][b] = minOf(min_cost[t][b], min_cost[pre][a] + passingFees[b])
                }
            }
        }

        val cost = (0..maxTime).fold(MX) { x, t -> minOf(x, min_cost[t][n - 1]) }
        return if (cost == MX) -1 else cost
    }


    @Test
    fun test1() {
        assertEquals(
            11, minCost(
                30,
                parse2dIntArray("[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]"),
                parseIntArray("[5,1,2,20,20,3]")
            )
        )

        assertEquals(
            48, minCost(
                29,
                parse2dIntArray("[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]"),
                parseIntArray("[5,1,2,20,20,3]")
            )
        )

        assertEquals(
            -1, minCost(
                25,
                parse2dIntArray("[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]"),
                parseIntArray("[5,1,2,20,20,3]")
            )
        )
    }
}
