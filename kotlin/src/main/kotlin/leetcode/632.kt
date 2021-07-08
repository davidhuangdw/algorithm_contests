package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*
import kotlin.collections.HashSet

class SmallestRange {
    fun smallestRange(nums: List<List<Int>>): IntArray {
        val n = nums.size
        val res = intArrayOf((-1e5).toInt(), 1e5.toInt())

        data class Node(val v: Int, val who: Int, val i: Int) : Comparable<Node> {
            override fun compareTo(other: Node): Int = v - other.v
        }

        val vp = hashMapOf<Int, HashSet<Int>>()
        val que = PriorityQueue<Node>()
        for (i in 0 until n)
            que.add(Node(nums[i][0], i, 0))
        val vs = mutableListOf<Int>()
        while (que.isNotEmpty()) {
            val node = que.poll()
            if (node.v !in vp) {
                vs.add(node.v)
                vp[node.v] = hashSetOf()
            }
            vp[node.v]!!.add(node.who)
            if (node.i + 1 < nums[node.who].size)
                que.add(Node(nums[node.who][node.i + 1], node.who, node.i + 1))
        }

        var r = -1
        var total = 0
        var count = MutableList(n) { 0 }
        for (l in 0 until vs.size) {
            while (r + 1 < vs.size && total < n) {
                r += 1
                for (w in vp[vs[r]]!!) {
                    count[w] += 1
                    if (count[w] == 1)
                        total += 1
                }
            }
            if (total < n) break
            if (vs[r] - vs[l] < res[1] - res[0]) {
                res[0] = vs[l]
                res[1] = vs[r]
            }

            for (w in vp[vs[l]]!!) {
                count[w] -= 1
                if (count[w] == 0)
                    total -= 1
            }
        }
        return res
    }


    @Test
    fun test1() {
        assertEquals(listOf(20, 24), smallestRange(
            parse2dIntArray("[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]").map { it.toList() }
        ).toList())
        assertEquals(listOf(1, 1), smallestRange(
            parse2dIntArray("[[1,2,3],[1,2,3],[1,2,3]]").map { it.toList() }
        ).toList())
        assertEquals(listOf(10, 11), smallestRange(
            parse2dIntArray("[[10,10],[11,11]]").map { it.toList() }
        ).toList())
        assertEquals(listOf(1, 7), smallestRange(
            parse2dIntArray("[[1],[2],[3],[4],[5],[6],[7]]").map { it.toList() }
        ).toList())
    }
}
