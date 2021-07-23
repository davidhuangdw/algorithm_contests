package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class SmallestChair5805 {
    fun smallestChair(times: Array<IntArray>, targetFriend: Int): Int {
        val n = times.size

        data class Node(val t: Int, val i: Int, val add: Int)

        val order = mutableListOf<Node>()
        for ((i, tp) in times.withIndex()) {
            val (s, e) = tp
            order.add(Node(s, i, 1))
            order.add(Node(e, i, -1))
        }

        var cur = PriorityQueue<Int>()
        for (i in 0 until n)
            cur.add(i)
        val pos = hashMapOf<Int, Int>()

        order.sortWith(compareBy({ it.t }, { it.add }))
        for ((t, i, add) in order) {
            if (add == 1 && i == targetFriend)
                return cur.peek()
            if (add == 1) {
                pos[i] = cur.poll()
            } else cur.add(pos[i])

        }
        return -1
    }


    @Test
    fun test1() {
        assertEquals(
            2,
            smallestChair(
                parse2dIntArray("[[33889,98676],[80071,89737],[44118,52565],[52992,84310],[78492,88209],[21695,67063],[84622,95452],[98048,98856],[98411,99433],[55333,56548],[65375,88566],[55011,62821],[48548,48656],[87396,94825],[55273,81868],[75629,91467]]"),
                6
            )
        )
        assertEquals(1, smallestChair(parse2dIntArray("[[1,4],[2,3],[4,6]]"), 1))
    }
}
