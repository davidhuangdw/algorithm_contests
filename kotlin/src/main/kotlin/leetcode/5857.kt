package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfUniqueGoodSubsequences5857 {
    fun numberOfUniqueGoodSubsequences(binary: String): Int {
        val n = binary.length
        val MOD = 1000_000_007
        val f = MutableList(n) { 0L }       // count of extra distinct substr[i~]
        val suf_all = MutableList(n + 1) { 1L } // sum{f[i~n]}, include empty
        val pre = MutableList(2) { n }
        var cnt = 0L
        for (i in n - 1 downTo 0) {
            val x = binary[i] - '0'
            f[i] = (MOD + suf_all[i + 1] - if (pre[x] < n) suf_all[pre[x] + 1] else 0) % MOD
            pre[x] = i
            if (x == 1)
                cnt = (cnt + f[i]) % MOD
            suf_all[i] = (suf_all[i + 1] + f[i]) % MOD
        }
        if (pre[0] < n)
            cnt = (cnt + 1) % MOD
        return cnt.toInt()
    }


    @Test
    fun test1() {
        assertEquals(
            846803618,
            numberOfUniqueGoodSubsequences("1100100010101010100100000111110001111001000010000010010111011")
        )
        assertEquals(2, numberOfUniqueGoodSubsequences("11"))
        assertEquals(2, numberOfUniqueGoodSubsequences("001"))
        assertEquals(5, numberOfUniqueGoodSubsequences("101"))
    }
}
