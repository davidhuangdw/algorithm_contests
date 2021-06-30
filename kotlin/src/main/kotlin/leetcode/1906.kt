package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test


class MinDifference {
    fun minDifference(nums: IntArray, queries: Array<IntArray>): IntArray {
        val n = nums.size
        val pos_arr = List(101) { mutableListOf<Int>() }
        for ((i, v) in nums.withIndex())
            pos_arr[v].add(i)

        return queries.map { q ->
            val (a, b) = q[0] to q[1]

            val vs = (1..100).filter { v ->
                val p = pos_arr[v]
                var (l, r) = 0 to p.size - 1
                while (l <= r) {
                    val md = (l + r) ushr 1
                    if (p[md] < a) l = md + 1
                    else r = md - 1
                }
                l < p.size && p[l] <= b
            }


            var min = Int.MAX_VALUE
            for (i in 0 until vs.size - 1)
                min = minOf(min, vs[i + 1] - vs[i])
            if (min == Int.MAX_VALUE) -1 else min
        }.toIntArray()
    }

    @Test
    fun test1() {
        var nums = intArrayOf(1, 3, 4, 8)
        var queries = arrayOf(
            intArrayOf(0, 3),
            intArrayOf(0, 1),
            intArrayOf(1, 2),
            intArrayOf(2, 3),
        )
        assertEquals(listOf(1, 2, 1, 4).toList(), minDifference(nums, queries).toList())

    }
}