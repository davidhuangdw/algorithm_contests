package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaximalNetworkRank {
    fun maximalNetworkRank(n: Int, roads: Array<IntArray>): Int {
        val deg = MutableList(n) { 0 }
        val conn = hashSetOf<Pair<Int, Int>>()
        for ((a, b) in roads) {
            deg[a] += 1
            deg[b] += 1
            conn.add(a to b)
            conn.add(b to a)
        }

        var a = Int.MIN_VALUE to hashSetOf<Int>()       // top
        var b = Int.MIN_VALUE to hashSetOf<Int>()       // second

        for ((i, d) in deg.withIndex()) {
            if (d > a.first) {
                b = a
                a = d to hashSetOf(i)
            } else if (d == a.first)
                a.second.add(i)
            else if (d > b.first)
                b = d to hashSetOf(i)
            else if (d == b.first)
                b.second.add(i)
        }

        fun exist_no_conn(s1: HashSet<Int>, s2: HashSet<Int>): Boolean {
            for (x in s1)
                for (y in s2)
                    if (x != y && x to y !in conn)
                        return true
            return false
        }
        return if (a.second.size > 1)
            a.first * 2 - if (exist_no_conn(a.second, a.second)) 0 else 1
        else
            a.first + b.first - if (exist_no_conn(a.second, b.second)) 0 else 1
    }


    @Test
    fun test1() {
        assertEquals(4, maximalNetworkRank(4, parse2dIntArray("[[0,1],[0,3],[1,2],[1,3]]")))
        assertEquals(5, maximalNetworkRank(5, parse2dIntArray("[[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]")))
        assertEquals(5, maximalNetworkRank(8, parse2dIntArray("[[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]")))
    }
}
