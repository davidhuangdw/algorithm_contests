package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfCombinations5837 {
    fun numberOfCombinations(num: String): Int { // todo: fix it
        if ('0' in num) return 0
        val MOD = 1000_000_007

        val n = num.length
        var k = 0
        while (1 shl (k + 1) <= n) {
            k += 1
        }
        val s = buildString {
            append(num)
            append("9".repeat(n))
        }
        val rank = List(k + 1) {
            MutableList(n * 2) {
                if (it < n) num[it] - '0' else {
                    n * 2
                }
            }
        }
        for (i in 1..k) {
            val cmp = compareBy<Int>({ rank[i - 1][it] }, { rank[i - 1][it + (1 shl (k - 1))] })
            var ord = (0 until n).sortedWith(cmp)
            rank[i][ord[0]] = 0
            for ((a, b) in ord.zipWithNext()) {
                rank[i][b] = rank[i][a] + if (cmp.compare(a, b) == 0) 0 else 1
            }
        }

        val log = MutableList(n + 1) { 0 }
        for (x in 2..n) {
            log[x] = log[x / 2] + 1
        }

        fun ge(i: Int, j: Int): Boolean {     //i <= j
            val l = j - i
            val k = log[l]
            val d = 1 shl k
            if (d == l)
                return rank[k][i] <= rank[k][j]
            else
                return rank[k][i] < rank[k][j] || (rank[k][i] == rank[k][j] && rank[k][i + d] <= rank[k][j + d])
        }

        val cnt = List(n + 1) { i -> MutableList(n / 2 + 1) { if (i == n) 1 else 0 } }
        val suf = List(n + 1) { i -> MutableList(n / 2 + 2) { 0 } }
        for (i in n downTo 0) {
            for (l in 1..n - i) {
                if (i == n) {
                    cnt[i][l] = if (i == n) 1 else 0
                    suf[i][l] = cnt[i][l]
                } else cnt[i][l] = (suf[i + l][l + 1] + if (ge(i, i + l)) cnt[i + l][l] else 0) % MOD
                suf[i][l] = suf[i][l + 1] + cnt[i][l] % MOD
            }
        }

        var sum = 0L
        for (l in 1..n / 2)
            sum = (sum + cnt[0][l]) % MOD
        return sum.toInt()
    }


    @Test
    fun test1() {
        assertEquals(2, numberOfCombinations("327"))
        assertEquals(0, numberOfCombinations("094"))
        assertEquals(101, numberOfCombinations("9999999999999"))
    }
}
