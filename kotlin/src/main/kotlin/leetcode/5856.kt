package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinSessions5856 {
    fun minSessions(tasks: IntArray, sessionTime: Int): Int {
        val n = tasks.size
        val M = 1 shl n
        val INF = 100
        val f = List(M) { i -> MutableList(sessionTime + 1) { if (i == 0) 1 else INF } }
        for (b in 1 until M)
            for (r in 0..sessionTime) {
                for ((i, t) in tasks.withIndex()) {
                    val d = b and (1 shl i)
                    if (b and d > 0)
                        f[b][r] = minOf(f[b][r], 1 + f[b - d][sessionTime - t], if (r >= t) f[b - d][r - t] else INF)
                }

            }
        return f[M - 1][sessionTime]
    }


    @Test
    fun test1() {
        assertEquals(1, minSessions(parseIntArray("[1,2,3,4,5]"), 15))
    }
}
