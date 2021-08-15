package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinNonZeroProduct5844 {
    fun minNonZeroProduct(p: Int): Int {
        val MOD = (1e9 + 7).toLong()
        val x = (1L shl p) - 1

        var res = x % MOD
        fun pow(x: Long, kk: Long): Long {
            var mul = x % MOD
            var r = 1L
            var k = kk
            while (k > 0) {
                if (k and 1 > 0)
                    r = r * mul % MOD
                mul = mul * mul % MOD
                k /= 2
            }
            return r
        }
        res = res * pow(x - 1, x / 2) % MOD
        return res.toInt()
    }


    @Test
    fun test1() {
        assertEquals(1, minNonZeroProduct(1))
        assertEquals(6, minNonZeroProduct(2))
        assertEquals(1512, minNonZeroProduct(3))
        assertEquals(505517599, minNonZeroProduct(32))
        assertEquals(112322054, minNonZeroProduct(35))
    }
}
