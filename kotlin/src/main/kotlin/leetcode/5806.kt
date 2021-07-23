package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SplitPainting5806 {
    fun splitPainting(segments: Array<IntArray>): List<List<Long>> {
        data class Node(val x: Long, val st: Boolean, val c: Long)

        val all = mutableListOf<Node>()
        for ((s, e, c) in segments) {
            all.add(Node(s.toLong(), true, c.toLong()))
            all.add(Node(e.toLong(), false, c.toLong()))
        }
        all.sortBy { it.x }

        var sum = 0L
        var prex = 0L
        var i = 0
        val res = mutableListOf<List<Long>>()
        val n = all.size
        while (i < n) {
            val x = all[i].x
            if (sum > 0L)
                res.add(listOf(prex, x, sum))
            while (i < n && all[i].x == x) {
                sum += all[i].c * if (all[i].st) 1 else -1
                i += 1
            }
            prex = x
        }

        return res
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
