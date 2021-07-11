package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ColorTheGrid {
    fun colorTheGrid(m: Int, n: Int): Int {
        val MOD = (1e9 + 7).toLong()

        var bound = 1
        repeat(m) { bound *= 3 }

        fun parse(x: Int): List<Int> {
            var order = mutableListOf<Int>()
            var cur = x
            repeat(m) {
                order.add(cur % 3)
                cur /= 3
            }
            return order.toList()
        }

        fun noConflict(a: Int, b: Int) = parse(a).zip(parse(b)).all { it.first != it.second }
        fun valid(line: List<Int>) = line.zipWithNext().all { it.first != it.second }

        val lines = mutableListOf<Int>()
        for (v in 0 until bound) {
            if (valid(parse(v)))
                lines.add(v)
        }

        val K = lines.size
        val adj = List(K) { hashSetOf<Int>() }
        for (i in 0 until lines.size)
            for (j in i + 1 until lines.size) {
                val (a, b) = lines[i] to lines[j]
                if (noConflict(a, b)) {
                    adj[i].add(j)
                    adj[j].add(i)
                }
            }

        val dp = List(n) { MutableList(bound) { 1L } }
        for (i in 1 until n) {
            for (j in 0 until K) {
                dp[i][j] = 0L
                for (k in adj[j])
                    dp[i][j] = (dp[i][j] + dp[i - 1][k]) % MOD
            }
        }

        var res = 0L
        for (j in 0 until K)
            res = (res + dp[n - 1][j]) % MOD
        return res.toInt()
    }


    @Test
    fun test1() {
        assertEquals(580986, colorTheGrid(5, 5))
    }
}
