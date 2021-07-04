package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LongestCommonSubpath {
    fun longestCommonSubpath(n: Int, paths: Array<IntArray>): Int {
        val m = paths.size
        paths.sortBy { it.size }
        var (l, r) = 1 to paths[0].size

        val (P, MOD) = (57).toLong() to (1e9 + 7).toLong()
        val SZ = paths[paths.size - 1].size
        val POW = mutableListOf(1L)
        for (i in 1..SZ)
            POW.add(POW[POW.size - 1] * P % MOD)
        fun exist(k: Int): Boolean {
            val count = hashMapOf<Long, Int>()
            var h = 0L
            val hs = LongArray(SZ)
            var path = paths[0]


            for ((i, v) in path.withIndex()) {
                h = (h * P + v) % MOD
                hs[i] = h
                if (i - k >= -1) {
                    val x = if (i - k == -1) h else (h - hs[i - k] * POW[k] % MOD + MOD) % MOD
                    if (x !in count)
                        count[x] = 1
                }
            }

            for (j in 1 until m) {
                var found = false
                path = paths[j]
                if (count.size == 0) return false
                h = 0
                for ((i, v) in path.withIndex()) {
                    h = (h * P + v) % MOD
                    hs[i] = h
                    if (i - k >= -1) {
                        val x = if (i - k == -1) h else (h - hs[i - k] * POW[k] % MOD + MOD) % MOD
                        if (x !in count || count[x] == j + 1) continue
                        if (count[x]!! < j)
                            count.remove(x)
                        else if (count[x] == j) {
                            count[x] = count[x]!! + 1
                            found = true
                        }
                    }
                }
                if (!found) return false
            }
            return true
        }

        while (l <= r) {
            val md = (l + r) ushr 1
            if (exist(md)) l = md + 1
            else r = md - 1
        }
        return r
    }


    @Test
    fun test1() {
        assertEquals(1, longestCommonSubpath(5, parse2dIntArray("[[1,2,3,4],[4,1,2,3],[4]]")))
        assertEquals(2, longestCommonSubpath(5, parse2dIntArray("[[0,1,2,3,4], [2,3,4], [4,0,1,2,3]]")))
    }
}
