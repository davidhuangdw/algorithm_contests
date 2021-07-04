package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CountGoodNumbers {
    val MOD = (1000_000_007).toLong()
    fun pow(a: Int, kk: Long): Long {
        val m = MOD
        var (x, mul, k) = Triple(1L, a.toLong(), kk)
        while (k > 0) {
            if (k and 1L != 0L) x = x * mul % m
            k = k ushr 1
            mul = mul * mul % m
        }
        return x
    }

    fun countGoodNumbers(n: Long): Int {
        val (odd, even) = n / 2 to (n + 1) / 2
        return (1L * pow(4, odd) * pow(5, even) % MOD).toInt()
    }


    @Test
    fun test1() {
        assertEquals(564908303, countGoodNumbers(50))
    }
}
